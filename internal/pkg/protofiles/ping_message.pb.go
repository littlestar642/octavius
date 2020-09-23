// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/ping_message.proto

package protofiles

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_ping_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_ping_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_ping_message_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recieved bool `protobuf:"varint,1,opt,name=recieved,proto3" json:"recieved,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_ping_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_ping_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_ping_message_proto_rawDescGZIP(), []int{1}
}

func (x *HealthResponse) GetRecieved() bool {
	if x != nil {
		return x.Recieved
	}
	return false
}

var File_internal_pkg_protofiles_ping_message_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_ping_message_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x22, 0x2c, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76,
	0x65, 0x64, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_protofiles_ping_message_proto_rawDescOnce sync.Once
	file_internal_pkg_protofiles_ping_message_proto_rawDescData = file_internal_pkg_protofiles_ping_message_proto_rawDesc
)

func file_internal_pkg_protofiles_ping_message_proto_rawDescGZIP() []byte {
	file_internal_pkg_protofiles_ping_message_proto_rawDescOnce.Do(func() {
		file_internal_pkg_protofiles_ping_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_protofiles_ping_message_proto_rawDescData)
	})
	return file_internal_pkg_protofiles_ping_message_proto_rawDescData
}

var file_internal_pkg_protofiles_ping_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_pkg_protofiles_ping_message_proto_goTypes = []interface{}{
	(*Ping)(nil),           // 0: Ping
	(*HealthResponse)(nil), // 1: HealthResponse
}
var file_internal_pkg_protofiles_ping_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_ping_message_proto_init() }
func file_internal_pkg_protofiles_ping_message_proto_init() {
	if File_internal_pkg_protofiles_ping_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_protofiles_ping_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_internal_pkg_protofiles_ping_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
			RawDescriptor: file_internal_pkg_protofiles_ping_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_ping_message_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_ping_message_proto_depIdxs,
		MessageInfos:      file_internal_pkg_protofiles_ping_message_proto_msgTypes,
	}.Build()
	File_internal_pkg_protofiles_ping_message_proto = out.File
	file_internal_pkg_protofiles_ping_message_proto_rawDesc = nil
	file_internal_pkg_protofiles_ping_message_proto_goTypes = nil
	file_internal_pkg_protofiles_ping_message_proto_depIdxs = nil
}