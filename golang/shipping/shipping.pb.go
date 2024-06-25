// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: shipping/shipping.proto

package shipping

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

type CreateShippingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateShippingRequest) Reset() {
	*x = CreateShippingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_shipping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShippingRequest) ProtoMessage() {}

func (x *CreateShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_shipping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShippingRequest.ProtoReflect.Descriptor instead.
func (*CreateShippingRequest) Descriptor() ([]byte, []int) {
	return file_shipping_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShippingRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateShippingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateShippingResponse) Reset() {
	*x = CreateShippingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_shipping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShippingResponse) ProtoMessage() {}

func (x *CreateShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_shipping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShippingResponse.ProtoReflect.Descriptor instead.
func (*CreateShippingResponse) Descriptor() ([]byte, []int) {
	return file_shipping_shipping_proto_rawDescGZIP(), []int{1}
}

var File_shipping_shipping_proto protoreflect.FileDescriptor

var file_shipping_shipping_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x18, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x47, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75,
	0x73, 0x65, 0x79, 0x69, 0x6e, 0x62, 0x61, 0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipping_shipping_proto_rawDescOnce sync.Once
	file_shipping_shipping_proto_rawDescData = file_shipping_shipping_proto_rawDesc
)

func file_shipping_shipping_proto_rawDescGZIP() []byte {
	file_shipping_shipping_proto_rawDescOnce.Do(func() {
		file_shipping_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipping_shipping_proto_rawDescData)
	})
	return file_shipping_shipping_proto_rawDescData
}

var file_shipping_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shipping_shipping_proto_goTypes = []any{
	(*CreateShippingRequest)(nil),  // 0: CreateShippingRequest
	(*CreateShippingResponse)(nil), // 1: CreateShippingResponse
}
var file_shipping_shipping_proto_depIdxs = []int32{
	0, // 0: Shipping.Create:input_type -> CreateShippingRequest
	1, // 1: Shipping.Create:output_type -> CreateShippingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shipping_shipping_proto_init() }
func file_shipping_shipping_proto_init() {
	if File_shipping_shipping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipping_shipping_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateShippingRequest); i {
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
		file_shipping_shipping_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateShippingResponse); i {
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
			RawDescriptor: file_shipping_shipping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipping_shipping_proto_goTypes,
		DependencyIndexes: file_shipping_shipping_proto_depIdxs,
		MessageInfos:      file_shipping_shipping_proto_msgTypes,
	}.Build()
	File_shipping_shipping_proto = out.File
	file_shipping_shipping_proto_rawDesc = nil
	file_shipping_shipping_proto_goTypes = nil
	file_shipping_shipping_proto_depIdxs = nil
}
