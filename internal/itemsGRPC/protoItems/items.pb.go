// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: items.proto

// protoc items.proto --go_out=.
// protoc items.proto --go-grpc_out=.

package protoItems

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

type ItemGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []uint64 `protobuf:"varint,1,rep,packed,name=Signature,proto3" json:"Signature,omitempty"`
	Text      string   `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *ItemGRPC) Reset() {
	*x = ItemGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemGRPC) ProtoMessage() {}

func (x *ItemGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemGRPC.ProtoReflect.Descriptor instead.
func (*ItemGRPC) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{0}
}

func (x *ItemGRPC) GetSignature() []uint64 {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ItemGRPC) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_items_proto protoreflect.FileDescriptor

var file_items_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3c, 0x0a, 0x08, 0x49, 0x74, 0x65,
	0x6d, 0x47, 0x52, 0x50, 0x43, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x32, 0x49, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x52,
	0x50, 0x43, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_items_proto_rawDescOnce sync.Once
	file_items_proto_rawDescData = file_items_proto_rawDesc
)

func file_items_proto_rawDescGZIP() []byte {
	file_items_proto_rawDescOnce.Do(func() {
		file_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_items_proto_rawDescData)
	})
	return file_items_proto_rawDescData
}

var file_items_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_items_proto_goTypes = []interface{}{
	(*ItemGRPC)(nil), // 0: protoItems.ItemGRPC
}
var file_items_proto_depIdxs = []int32{
	0, // 0: protoItems.ItemService.SetSignature:input_type -> protoItems.ItemGRPC
	0, // 1: protoItems.ItemService.SetSignature:output_type -> protoItems.ItemGRPC
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_items_proto_init() }
func file_items_proto_init() {
	if File_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemGRPC); i {
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
			RawDescriptor: file_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_items_proto_goTypes,
		DependencyIndexes: file_items_proto_depIdxs,
		MessageInfos:      file_items_proto_msgTypes,
	}.Build()
	File_items_proto = out.File
	file_items_proto_rawDesc = nil
	file_items_proto_goTypes = nil
	file_items_proto_depIdxs = nil
}
