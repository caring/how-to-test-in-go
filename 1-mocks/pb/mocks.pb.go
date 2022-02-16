// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mocks.proto

package pb

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

type GetFizzBuzzRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetFizzBuzzRequest) Reset() {
	*x = GetFizzBuzzRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mocks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFizzBuzzRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFizzBuzzRequest) ProtoMessage() {}

func (x *GetFizzBuzzRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mocks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFizzBuzzRequest.ProtoReflect.Descriptor instead.
func (*GetFizzBuzzRequest) Descriptor() ([]byte, []int) {
	return file_mocks_proto_rawDescGZIP(), []int{0}
}

func (x *GetFizzBuzzRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type GetFizzBuzzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FizzBuzz string `protobuf:"bytes,1,opt,name=fizz_buzz,json=fizzBuzz,proto3" json:"fizz_buzz,omitempty"`
}

func (x *GetFizzBuzzResponse) Reset() {
	*x = GetFizzBuzzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mocks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFizzBuzzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFizzBuzzResponse) ProtoMessage() {}

func (x *GetFizzBuzzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mocks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFizzBuzzResponse.ProtoReflect.Descriptor instead.
func (*GetFizzBuzzResponse) Descriptor() ([]byte, []int) {
	return file_mocks_proto_rawDescGZIP(), []int{1}
}

func (x *GetFizzBuzzResponse) GetFizzBuzz() string {
	if x != nil {
		return x.FizzBuzz
	}
	return ""
}

var File_mocks_proto protoreflect.FileDescriptor

var file_mocks_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x63, 0x6b, 0x73, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x7a, 0x7a, 0x42,
	0x75, 0x7a, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a,
	0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x7a,
	0x7a, 0x5f, 0x62, 0x75, 0x7a, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x7a, 0x7a, 0x42, 0x75, 0x7a, 0x7a, 0x32, 0x53, 0x0a, 0x0b, 0x4d, 0x6f, 0x63, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x7a, 0x7a,
	0x42, 0x75, 0x7a, 0x7a, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x7a, 0x7a, 0x42,
	0x75, 0x7a, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x68, 0x6f, 0x77, 0x2d, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x69, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x31, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mocks_proto_rawDescOnce sync.Once
	file_mocks_proto_rawDescData = file_mocks_proto_rawDesc
)

func file_mocks_proto_rawDescGZIP() []byte {
	file_mocks_proto_rawDescOnce.Do(func() {
		file_mocks_proto_rawDescData = protoimpl.X.CompressGZIP(file_mocks_proto_rawDescData)
	})
	return file_mocks_proto_rawDescData
}

var file_mocks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mocks_proto_goTypes = []interface{}{
	(*GetFizzBuzzRequest)(nil),  // 0: mocks.GetFizzBuzzRequest
	(*GetFizzBuzzResponse)(nil), // 1: mocks.GetFizzBuzzResponse
}
var file_mocks_proto_depIdxs = []int32{
	0, // 0: mocks.MockService.GetFizzBuzz:input_type -> mocks.GetFizzBuzzRequest
	1, // 1: mocks.MockService.GetFizzBuzz:output_type -> mocks.GetFizzBuzzResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mocks_proto_init() }
func file_mocks_proto_init() {
	if File_mocks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mocks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFizzBuzzRequest); i {
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
		file_mocks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFizzBuzzResponse); i {
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
			RawDescriptor: file_mocks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mocks_proto_goTypes,
		DependencyIndexes: file_mocks_proto_depIdxs,
		MessageInfos:      file_mocks_proto_msgTypes,
	}.Build()
	File_mocks_proto = out.File
	file_mocks_proto_rawDesc = nil
	file_mocks_proto_goTypes = nil
	file_mocks_proto_depIdxs = nil
}
