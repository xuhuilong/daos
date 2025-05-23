/*
 * (C) Copyright 2016-2024 Intel Corporation.
 * (C) Copyright 2025 Hewlett Packard Enterprise Development LP
 *
 * SPDX-License-Identifier: BSD-2-Clause-Patent
 */
/**
 * This file is part of CaRT. It gives out the main CaRT internal function
 * declarations which are not included by other specific header files.
 */

#ifndef __CRT_INTERNAL_FNS_H__
#define __CRT_INTERNAL_FNS_H__

/** crt_init.c */
bool crt_initialized(void);

/** crt_register.c */
int crt_opc_map_create(void);
void crt_opc_map_destroy(struct crt_opc_map *map);
struct crt_opc_info *crt_opc_lookup(struct crt_opc_map *map, crt_opcode_t opc,
				    int locked);

/** crt_context.c */
/* return values of crt_context_req_track, in addition to standard
 * gurt error values.
 */
enum {
	CRT_REQ_TRACK_IN_INFLIGHQ = 0,
	CRT_REQ_TRACK_IN_WAITQ,
};

int crt_context_req_track(struct crt_rpc_priv *rpc_priv);
bool crt_context_empty(crt_provider_t provider, int locked);
void crt_context_req_untrack(struct crt_rpc_priv *rpc_priv);
crt_context_t crt_context_lookup(int ctx_idx);
crt_context_t crt_context_lookup_locked(int ctx_idx);
void
record_quota_resource(crt_context_t crt_ctx, crt_quota_type_t quota);
int
get_quota_resource(crt_context_t crt_ctx, crt_quota_type_t quota);
void
     put_quota_resource(crt_context_t crt_ctx, crt_quota_type_t quota);
void crt_rpc_complete_and_unlock(struct crt_rpc_priv *rpc_priv, int rc);
int crt_req_timeout_track(struct crt_rpc_priv *rpc_priv);
void crt_req_timeout_untrack(struct crt_rpc_priv *rpc_priv);
void crt_req_force_completion(struct crt_rpc_priv *rpc_priv);

/** some simple helper functions */

static inline bool
crt_is_service()
{
	return crt_gdata.cg_server;
}

static inline void
crt_bulk_desc_dup(struct crt_bulk_desc *bulk_desc_new,
		  struct crt_bulk_desc *bulk_desc)
{
	D_ASSERT(bulk_desc_new != NULL && bulk_desc != NULL);
	*bulk_desc_new = *bulk_desc;
}

void
crt_hdlr_proto_query(crt_rpc_t *rpc_req);

int
crt_register_proto_fi(crt_endpoint_t *ep);

int
crt_register_proto_ctl(crt_endpoint_t *ep);

void
crt_trigger_hlc_error_cb(void);

void
crt_trigger_event_cbs(d_rank_t rank, uint64_t incarnation, enum crt_event_source src,
		      enum crt_event_type type);

#endif /* __CRT_INTERNAL_FNS_H__ */
