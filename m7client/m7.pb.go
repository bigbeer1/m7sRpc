// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: m7.proto

package m7client

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

// 通用空返回
type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m7_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_m7_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_m7_proto_rawDescGZIP(), []int{0}
}

type TestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestReq) Reset() {
	*x = TestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m7_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReq) ProtoMessage() {}

func (x *TestReq) ProtoReflect() protoreflect.Message {
	mi := &file_m7_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReq.ProtoReflect.Descriptor instead.
func (*TestReq) Descriptor() ([]byte, []int) {
	return file_m7_proto_rawDescGZIP(), []int{1}
}

func (x *TestReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestResp) Reset() {
	*x = TestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m7_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResp) ProtoMessage() {}

func (x *TestResp) ProtoReflect() protoreflect.Message {
	mi := &file_m7_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResp.ProtoReflect.Descriptor instead.
func (*TestResp) Descriptor() ([]byte, []int) {
	return file_m7_proto_rawDescGZIP(), []int{2}
}

func (x *TestResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_m7_proto protoreflect.FileDescriptor

var file_m7_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x37, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x33, 0x0a,
	0x02, 0x4d, 0x37, 0x12, 0x2d, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x6d, 0x37,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x6d, 0x37, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6d, 0x37, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_m7_proto_rawDescOnce sync.Once
	file_m7_proto_rawDescData = file_m7_proto_rawDesc
)

func file_m7_proto_rawDescGZIP() []byte {
	file_m7_proto_rawDescOnce.Do(func() {
		file_m7_proto_rawDescData = protoimpl.X.CompressGZIP(file_m7_proto_rawDescData)
	})
	return file_m7_proto_rawDescData
}

var file_m7_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_m7_proto_goTypes = []interface{}{
	(*CommonResp)(nil), // 0: m7client.CommonResp
	(*TestReq)(nil),    // 1: m7client.TestReq
	(*TestResp)(nil),   // 2: m7client.TestResp
}
var file_m7_proto_depIdxs = []int32{
	1, // 0: m7client.M7.Test:input_type -> m7client.TestReq
	2, // 1: m7client.M7.Test:output_type -> m7client.TestResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_m7_proto_init() }
func file_m7_proto_init() {
	if File_m7_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_m7_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_m7_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReq); i {
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
		file_m7_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResp); i {
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
			RawDescriptor: file_m7_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_m7_proto_goTypes,
		DependencyIndexes: file_m7_proto_depIdxs,
		MessageInfos:      file_m7_proto_msgTypes,
	}.Build()
	File_m7_proto = out.File
	file_m7_proto_rawDesc = nil
	file_m7_proto_goTypes = nil
	file_m7_proto_depIdxs = nil
}