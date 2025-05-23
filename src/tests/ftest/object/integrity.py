"""
  (C) Copyright 2020-2024 Intel Corporation.
  (C) Copyright 2025 Hewlett Packard Enterprise Development LP

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""
import ctypes
import time

import avocado
from apricot import TestWithServers
from general_utils import create_string_buffer
from pydaos.raw import DaosApiError, DaosObj, IORequest
from test_utils_container import add_container
from test_utils_pool import add_pool


class ObjectDataValidation(TestWithServers):
    """
    Test Class Description:
        Tests that create Different length records,
        Disconnect the pool/container and reconnect,
        validate the data after reconnect.

    :avocado: recursive
    """

    def get_object(self, container):
        """Get a new DaosObj object.

        Args:
            container (TestContainer): _description_

        Returns:
            DaosObj: a new DaosObj object
        """
        self.log_step('Creating a DaosObj')
        container.open()
        obj = DaosObj(self.context, container.container)
        obj.create(objcls=1)
        return obj

    def get_io_request(self, obj):
        """Get a new IORequest object.

        Args:
            obj (DaosObj): the DaosObj used to create the IORequest

        Returns:
            IORequest: a new IORequest object
        """
        obj.open()
        return IORequest(self.context, obj.container, obj, objtype=4)

    def reconnect(self, obj, container):
        """Reconnect the pool/container and reopen the Object for read verification.

        Args:
            obj (DaosObj): _description_
            container (TestContainer): _description_

        Returns:
            IORequest: a new IORequest object
        """
        # Close the Obj/Container, Disconnect the Pool.
        obj.close()
        container.close()
        container.pool.disconnect()
        time.sleep(5)
        # Connect Pool, Open Container and Object
        container.pool.connect(2)
        container.open()
        obj.open()
        return self.get_io_request(obj)

    @avocado.fail_on(DaosApiError)
    def test_invalid_tx_commit_close(self):
        """
        Test ID:
            (1)DAOS-1346: Verify commit tx bad parameter behavior.
            (2)DAOS-1343: Verify tx_close bad parameter behavior.
            (3)DAOS-1342: Verify tx_close through daos_api.
            (4)DAOS-1338: Add and verify tx_abort through daos_api.
            (5)DAOS-1339: Verify tx_abort bad parameter behavior.
        Test Description:
            Write Avocado Test to verify commit tx and close tx
                          bad parameter behavior.
        :avocado: tags=all,full_regression
        :avocado: tags=vm
        :avocado: tags=object,objectvalidation
        :avocado: tags=ObjectDataValidation,test_invalid_tx_commit_close

        """
        record_length = self.params.get("length", '/run/record/*')

        pool = add_pool(self)
        container = add_container(self, pool)
        obj = self.get_object(container)
        ioreq = self.get_io_request(obj)

        self.log.info("==Writing the Single Dataset for negative test...")
        record_index = 0
        expected_error = "RC: -1002"
        dkey = 0
        akey = 0
        indata = "{0}".format(str(akey)[0]) * record_length[record_index]
        c_dkey = create_string_buffer("dkey {0}".format(dkey))
        c_akey = create_string_buffer("akey {0}".format(akey))
        c_value = create_string_buffer(indata)
        c_size = ctypes.c_size_t(ctypes.sizeof(c_value))
        try:
            new_transaction = container.container.get_new_tx()
        except DaosApiError as excep:
            # initial container get_new_tx failed, skip rest of the test
            self.fail("##container get_new_tx failed: {}".format(excep))
        invalid_transaction = new_transaction + self.random.randint(1000, 383838)
        self.log.info("==new_transaction=     %s", new_transaction)
        self.log.info("==invalid_transaction= %s", invalid_transaction)
        ioreq.single_insert(c_dkey, c_akey, c_value, c_size, new_transaction)
        try:
            container.container.commit_tx(invalid_transaction)
            self.fail("##(1.1)Container.commit_tx passing with invalid handle")
        except DaosApiError as excep:
            self.log.info(str(excep))
            self.log.info("==(1)Expecting failure: invalid Container.commit_tx.")
            if expected_error not in str(excep):
                self.fail("##(1.2)Expecting error RC: -1002, but got {}.".format(str(excep)))
        try:
            container.container.close_tx(invalid_transaction)
            self.fail("##(2.1)Container.close_tx passing with invalid handle")
        except DaosApiError as excep:
            self.log.info(str(excep))
            self.log.info("==(2)Expecting failure: invalid Container.commit_tx.")
            if expected_error not in str(excep):
                self.fail("##(2.2)Expecting error RC: -1002, but got {}.".format(str(excep)))
        try:
            container.container.close_tx(new_transaction)
            self.log.info("==(3)container.close_tx test passed.")
        except DaosApiError as excep:
            self.log.info(str(excep))
            self.fail("##(3)Failed on close_tx.")

        try:
            container.container.abort_tx(invalid_transaction)
            self.fail("##(4.1)Container.abort_tx passing with invalid handle")
        except DaosApiError as excep:
            self.log.info(str(excep))
            self.log.info("==(4)Expecting failure: invalid Container.abort_tx.")
            if expected_error not in str(excep):
                self.fail("##(4.2)Expecting error RC: -1002, but got {}.".format(str(excep)))

        # Try to abort the transaction which already closed.
        try:
            container.container.abort_tx(new_transaction)
            self.fail("##(5.1)Container.abort_tx passing with a closed handle")
        except DaosApiError as excep:
            self.log.info(str(excep))
            self.log.info("==(5)Expecting failure: Container.abort_tx closed handle.")
            if expected_error not in str(excep):
                self.fail("##(5.2)Expecting error RC: -1002, but got {}.".format(str(excep)))

        # open another transaction for abort test
        try:
            new_transaction2 = container.container.get_new_tx()
        except DaosApiError as excep:
            self.fail("##(6.1)container get_new_tx failed: {}".format(excep))
        self.log.info("==new_transaction2=     %s", new_transaction2)
        ioreq.single_insert(c_dkey, c_akey, c_value, c_size, new_transaction2)
        try:
            container.container.abort_tx(new_transaction2)
            self.log.info("==(6)container.abort_tx test passed.")
        except DaosApiError as excep:
            self.log.info(str(excep))
            self.fail("##(6.2)Failed on abort_tx.")

        container.container.close_tx(new_transaction2)

        self.log.info('Test passed')

    @avocado.fail_on(DaosApiError)
    def test_single_object_validation(self):
        """
        Test ID: DAOS-707
        Test Description: Write Avocado Test to verify single data after
                          pool/container disconnect/reconnect.
        :avocado: tags=all,full_regression
        :avocado: tags=vm
        :avocado: tags=object,objectvalidation
        :avocado: tags=ObjectDataValidation,test_single_object_validation
        """
        no_of_dkeys = self.params.get("no_of_dkeys", '/run/dkeys/*')[0]
        no_of_akeys = self.params.get("no_of_akeys", '/run/akeys/*')[0]
        record_length = self.params.get("length", '/run/record/*')

        pool = add_pool(self)
        container = add_container(self, pool)
        obj = self.get_object(container)
        ioreq = self.get_io_request(obj)

        self.log_step("Writing the Single Dataset")
        record_index = 0
        for dkey in range(no_of_dkeys):
            for akey in range(no_of_akeys):
                indata = "{0}".format(str(akey)[0]) * record_length[record_index]
                c_dkey = create_string_buffer("dkey {0}".format(dkey))
                c_akey = create_string_buffer("akey {0}".format(akey))
                c_value = create_string_buffer(indata)
                c_size = ctypes.c_size_t(ctypes.sizeof(c_value))

                ioreq.single_insert(c_dkey, c_akey, c_value, c_size)
                record_index = record_index + 1
                if record_index == len(record_length):
                    record_index = 0

        ioreq = self.reconnect(obj, container)

        self.log_step("Single Dataset Verification -- Started")
        record_index = 0
        transaction_index = 0
        for dkey in range(no_of_dkeys):
            for akey in range(no_of_akeys):
                indata = str(akey)[0] * record_length[record_index]
                c_dkey = create_string_buffer("dkey {0}".format(dkey))
                c_akey = create_string_buffer("akey {0}".format(akey))
                val = ioreq.single_fetch(c_dkey, c_akey, len(indata) + 1)
                if indata != str(val.value, 'utf-8'):
                    msg = f"ERROR: Data mismatch for dkey = dkey {dkey}, akey = akey {akey}"
                    self.log.error(msg)
                    self.fail(msg)

                transaction_index = transaction_index + 1
                record_index = record_index + 1
                if record_index == len(record_length):
                    record_index = 0

        self.log_step('Test passed')

    @avocado.fail_on(DaosApiError)
    def test_array_object_validation(self):
        """
        Test ID: DAOS-707
        Test Description: Write Avocado Test to verify Array data after
                          pool/container disconnect/reconnect.
        :avocado: tags=all,full_regression
        :avocado: tags=vm
        :avocado: tags=object,objectvalidation
        :avocado: tags=ObjectDataValidation,test_array_object_validation
        """
        no_of_dkeys = self.params.get("no_of_dkeys", '/run/dkeys/*')[0]
        no_of_akeys = self.params.get("no_of_akeys", '/run/akeys/*')[0]
        array_size = self.params.get("size", '/array_size/')
        record_length = self.params.get("length", '/run/record/*')

        pool = add_pool(self)
        container = add_container(self, pool)
        obj = self.get_object(container)
        ioreq = self.get_io_request(obj)

        self.log_step("Writing the Array Dataset")
        record_index = 0
        for dkey in range(no_of_dkeys):
            for akey in range(no_of_akeys):
                c_values = []
                value = "{0}".format(str(akey)[0]) * record_length[record_index]
                for item in range(array_size):
                    c_values.append((create_string_buffer(value), len(value) + 1))
                c_dkey = create_string_buffer("dkey {0}".format(dkey))
                c_akey = create_string_buffer("akey {0}".format(akey))

                ioreq.insert_array(c_dkey, c_akey, c_values)

                record_index = record_index + 1
                if record_index == len(record_length):
                    record_index = 0

        ioreq = self.reconnect(obj, container)

        self.log_step("Array Dataset Verification -- Started")
        record_index = 0
        transaction_index = 0
        for dkey in range(no_of_dkeys):
            for akey in range(no_of_akeys):
                indata = []
                value = "{0}".format(str(akey)[0]) * record_length[record_index]
                for item in range(array_size):
                    indata.append(value)
                c_dkey = create_string_buffer("dkey {0}".format(dkey))
                c_akey = create_string_buffer("akey {0}".format(akey))
                c_rec_count = ctypes.c_uint(len(indata))
                c_rec_size = ctypes.c_size_t(len(indata[0]) + 1)

                outdata = ioreq.fetch_array(c_dkey, c_akey, c_rec_count, c_rec_size)

                for item in enumerate(indata):
                    if indata[item[0]] != str(outdata[item[0]], 'utf-8')[:-1]:
                        msg = f"ERROR: Data mismatch for dkey = dkey {dkey}, akey = akey {akey}"
                        self.log.error(msg)
                        self.fail(msg)

                transaction_index = transaction_index + 1
                record_index = record_index + 1
                if record_index == len(record_length):
                    record_index = 0

        self.log_step('Test passed')
