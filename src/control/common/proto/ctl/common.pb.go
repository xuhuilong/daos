// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.5.0
// source: ctl/common.proto

package ctl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseStatus int32

const (
	ResponseStatus_CTL_SUCCESS     ResponseStatus = 0
	ResponseStatus_CTL_IN_PROGRESS ResponseStatus = 1  // Not yet completed
	ResponseStatus_CTL_WAITING     ResponseStatus = 2  // Blocked
	ResponseStatus_CTL_ERR_CONF    ResponseStatus = -1 // Config file parsing error
	ResponseStatus_CTL_ERR_NVME    ResponseStatus = -2 // NVMe subsystem error
	ResponseStatus_CTL_ERR_SCM     ResponseStatus = -3 // SCM subsystem error
	ResponseStatus_CTL_ERR_APP     ResponseStatus = -4 // Other application error
	ResponseStatus_CTL_ERR_UNKNOWN ResponseStatus = -5 // Unknown error
	ResponseStatus_CTL_NO_IMPL     ResponseStatus = -6 // No implementation
)

// Enum value maps for ResponseStatus.
var (
	ResponseStatus_name = map[int32]string{
		0:  "CTL_SUCCESS",
		1:  "CTL_IN_PROGRESS",
		2:  "CTL_WAITING",
		-1: "CTL_ERR_CONF",
		-2: "CTL_ERR_NVME",
		-3: "CTL_ERR_SCM",
		-4: "CTL_ERR_APP",
		-5: "CTL_ERR_UNKNOWN",
		-6: "CTL_NO_IMPL",
	}
	ResponseStatus_value = map[string]int32{
		"CTL_SUCCESS":     0,
		"CTL_IN_PROGRESS": 1,
		"CTL_WAITING":     2,
		"CTL_ERR_CONF":    -1,
		"CTL_ERR_NVME":    -2,
		"CTL_ERR_SCM":     -3,
		"CTL_ERR_APP":     -4,
		"CTL_ERR_UNKNOWN": -5,
		"CTL_NO_IMPL":     -6,
	}
)

func (x ResponseStatus) Enum() *ResponseStatus {
	p := new(ResponseStatus)
	*p = x
	return p
}

func (x ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ctl_common_proto_enumTypes[0].Descriptor()
}

func (ResponseStatus) Type() protoreflect.EnumType {
	return &file_ctl_common_proto_enumTypes[0]
}

func (x ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseStatus.Descriptor instead.
func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_ctl_common_proto_rawDescGZIP(), []int{0}
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_ctl_common_proto_rawDescGZIP(), []int{0}
}

type FilePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FilePath) Reset() {
	*x = FilePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePath) ProtoMessage() {}

func (x *FilePath) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePath.ProtoReflect.Descriptor instead.
func (*FilePath) Descriptor() ([]byte, []int) {
	return file_ctl_common_proto_rawDescGZIP(), []int{1}
}

func (x *FilePath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ResponseState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ctl.ResponseStatus" json:"status,omitempty"`
	Error  string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Info   string         `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ResponseState) Reset() {
	*x = ResponseState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseState) ProtoMessage() {}

func (x *ResponseState) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseState.ProtoReflect.Descriptor instead.
func (*ResponseState) Descriptor() ([]byte, []int) {
	return file_ctl_common_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseState) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_CTL_SUCCESS
}

func (x *ResponseState) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ResponseState) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_ctl_common_proto protoreflect.FileDescriptor

var file_ctl_common_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x74, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x63, 0x74, 0x6c, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x22, 0x1e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x66, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x74, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x2a, 0xe9, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x43, 0x54, 0x4c, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x43, 0x54, 0x4c, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x54, 0x4c, 0x5f, 0x57, 0x41, 0x49, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x0c, 0x43, 0x54, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x12, 0x19, 0x0a, 0x0c, 0x43, 0x54, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x56, 0x4d, 0x45,
	0x10, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x18, 0x0a, 0x0b, 0x43,
	0x54, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x43, 0x4d, 0x10, 0xfd, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x18, 0x0a, 0x0b, 0x43, 0x54, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x1c, 0x0a, 0x0f, 0x43, 0x54, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0xfb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x18, 0x0a,
	0x0b, 0x43, 0x54, 0x4c, 0x5f, 0x4e, 0x4f, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x10, 0xfa, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ctl_common_proto_rawDescOnce sync.Once
	file_ctl_common_proto_rawDescData = file_ctl_common_proto_rawDesc
)

func file_ctl_common_proto_rawDescGZIP() []byte {
	file_ctl_common_proto_rawDescOnce.Do(func() {
		file_ctl_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_ctl_common_proto_rawDescData)
	})
	return file_ctl_common_proto_rawDescData
}

var file_ctl_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ctl_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ctl_common_proto_goTypes = []interface{}{
	(ResponseStatus)(0),   // 0: ctl.ResponseStatus
	(*EmptyReq)(nil),      // 1: ctl.EmptyReq
	(*FilePath)(nil),      // 2: ctl.FilePath
	(*ResponseState)(nil), // 3: ctl.ResponseState
}
var file_ctl_common_proto_depIdxs = []int32{
	0, // 0: ctl.ResponseState.status:type_name -> ctl.ResponseStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ctl_common_proto_init() }
func file_ctl_common_proto_init() {
	if File_ctl_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ctl_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ctl_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilePath); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ctl_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ctl_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ctl_common_proto_goTypes,
		DependencyIndexes: file_ctl_common_proto_depIdxs,
		EnumInfos:         file_ctl_common_proto_enumTypes,
		MessageInfos:      file_ctl_common_proto_msgTypes,
	}.Build()
	File_ctl_common_proto = out.File
	file_ctl_common_proto_rawDesc = nil
	file_ctl_common_proto_goTypes = nil
	file_ctl_common_proto_depIdxs = nil
}
