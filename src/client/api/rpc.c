/**
 * (C) Copyright 2016-2024 Intel Corporation.
 * (C) Copyright 2025 Hewlett Packard Enterprise Development LP
 *
 * SPDX-License-Identifier: BSD-2-Clause-Patent
 */
#define D_LOGFAC	DD_FAC(client)

#include <daos/rpc.h>
#include <daos/event.h>
#include <daos/mgmt.h>

static void
daos_rpc_cb(const struct crt_cb_info *cb_info)
{
	tse_task_t	*task = cb_info->cci_arg;
	int		rc = cb_info->cci_rc;

	if (cb_info->cci_rc == -DER_TIMEDOUT) {
		/** TODO */
		;
	}

	tse_task_complete(task, rc);
}

int
daos_rpc_complete(crt_rpc_t *rpc, tse_task_t *task)
{
	struct crt_cb_info cbinfo;

	cbinfo.cci_arg = task;
	cbinfo.cci_rc  = 0;
	daos_rpc_cb(&cbinfo);
	crt_req_decref(rpc);
	return 0;
}

int
daos_rpc_send(crt_rpc_t *rpc, tse_task_t *task)
{
	int rc;

	rc = crt_req_send(rpc, daos_rpc_cb, task);
	if (rc != 0) {
		/** task will be completed in CB above */
		rc = 0;
	}

	return rc;
}

struct daos_rpc_status {
	int completed;
	int status;
};

static void
daos_rpc_wait(crt_context_t *ctx, struct daos_rpc_status *status)
{
	/* Wait on the event to complete */
	while (!status->completed) {
		int rc = 0;

		rc = crt_progress(ctx, 0);
		if (rc && rc != -DER_TIMEDOUT) {
			D_ERROR("failed to progress CART context: %d\n", rc);
			break;
		}
	}
}

static void
daos_rpc_wait_cb(const struct crt_cb_info *cb_info)
{
	struct daos_rpc_status *status = cb_info->cci_arg;

	status->completed = 1;
	if (status->status == 0)
		status->status = cb_info->cci_rc;
}

int
daos_rpc_send_wait(crt_rpc_t *rpc)
{
	struct daos_rpc_status status = { 0 };
	int rc;

	rc = crt_req_send(rpc, daos_rpc_wait_cb, &status);
	if (rc != 0)
		return rc;

	daos_rpc_wait(rpc->cr_ctx, &status);
	rc = status.status;

	return rc;
}

struct rpc_proto {
	int            rank_idx;
	crt_endpoint_t ep;
	int            version;
	int            rc;
	bool           completed;
	crt_opcode_t   base_opc;
	uint32_t      *ver_array;
	uint32_t       array_size;
	uint32_t       timeout;
};

static void
query_cb(struct crt_proto_query_cb_info *cb_info)
{
	struct rpc_proto	*rproto = (struct rpc_proto *)cb_info->pq_arg;
	int			 rc;

	if (daos_rpc_retryable_rc(cb_info->pq_rc)) {
		int      nr_ranks;
		d_rank_t rank;

		/** select next rank to issue the retry proto query rpc to */
		nr_ranks = dc_mgmt_net_get_num_srv_ranks();
		D_ASSERT(nr_ranks > 0);
		rproto->rank_idx = (rproto->rank_idx + 1) % nr_ranks;
		rank             = dc_mgmt_net_get_srv_rank(rproto->rank_idx);
		D_ASSERT(rank != CRT_NO_RANK);
		rproto->ep.ep_rank = rank;

		rproto->timeout += 3;
		rc = crt_proto_query_with_ctx(&rproto->ep, rproto->base_opc, rproto->ver_array,
					      rproto->array_size, rproto->timeout, query_cb, rproto,
					      daos_get_crt_ctx());
		if (rc) {
			D_ERROR("crt_proto_query_with_ctx() failed: "DF_RC"\n", DP_RC(rc));
			rproto->rc = rc;
			rproto->completed = true;
		}
	} else {
		rproto->rc = cb_info->pq_rc;
		rproto->version = cb_info->pq_ver;
		rproto->completed = true;
	}
}

int
daos_rpc_proto_query(crt_opcode_t base_opc, uint32_t *ver_array, int count, int *ret_ver)
{
	struct dc_mgmt_sys	*sys;
	struct rpc_proto	*rproto = NULL;
	crt_context_t		 ctx = daos_get_crt_ctx();
	int                      rc;
	int			 i;
	int                      nr_ranks;
	d_rank_t                 rank;

	rc = dc_mgmt_sys_attach(NULL, &sys);
	if (rc != 0) {
		D_ERROR("failed to attach to grp rc "DF_RC"\n", DP_RC(rc));
		return rc;
	}

	D_ALLOC_PTR(rproto);
	if (rproto == NULL)
		D_GOTO(out_detach, rc = -DER_NOMEM);

	/** select a random rank to issue the proto query rpc to */
	nr_ranks = dc_mgmt_net_get_num_srv_ranks();
	if (nr_ranks == 0) {
		D_ERROR("failed to select an attached ranks: no attached ranks");
		D_GOTO(out_free, -DER_NONEXIST);
	}
	rproto->rank_idx = d_rand() % nr_ranks;
	rank             = dc_mgmt_net_get_srv_rank(rproto->rank_idx);
	D_ASSERT(rank != CRT_NO_RANK);
	rproto->ep.ep_rank = rank;

	rproto->ep.ep_tag  = 0;
	rproto->ver_array = ver_array;
	rproto->array_size = count;
	rproto->ep.ep_grp  = sys->sy_group;
	rproto->base_opc = base_opc;
	rproto->timeout    = 3;

	rc = crt_proto_query_with_ctx(&rproto->ep, base_opc, ver_array, count, rproto->timeout,
				      query_cb, rproto, ctx);
	if (rc) {
		D_ERROR("crt_proto_query_with_ctx() failed: "DF_RC"\n", DP_RC(rc));
		D_GOTO(out_free, rc);
	}

	while (!rproto->completed) {
		rc = crt_progress(ctx, 0);
		if (rc && rc != -DER_TIMEDOUT) {
			D_ERROR("failed to progress CART context: %d\n", rc);
			D_GOTO(out_free, rc);
		}
	}

	if (rproto->rc != -DER_SUCCESS) {
		rc = rproto->rc;
		D_ERROR("crt_proto_query()failed: "DF_RC"\n", DP_RC(rc));
		D_GOTO(out_free, rc);
	}
	rc = 0;

	for (i = 0; i < count; i++)
		if (ver_array[i] == rproto->version)
			break;

	if (i == count) {
		D_ERROR("Invalid RPC protocol version %d\n", rproto->version);
		rc = -DER_PROTO;
	} else {
		*ret_ver = rproto->version;
	}
out_free:
	D_FREE(rproto);
out_detach:
	dc_mgmt_sys_detach(sys);

	return rc;

}
