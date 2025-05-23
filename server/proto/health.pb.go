// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: health.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HeathRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeathRequest) Reset() {
	*x = HeathRequest{}
	mi := &file_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeathRequest) ProtoMessage() {}

func (x *HeathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeathRequest.ProtoReflect.Descriptor instead.
func (*HeathRequest) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

type HeathResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeathResponse) Reset() {
	*x = HeathResponse{}
	mi := &file_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeathResponse) ProtoMessage() {}

func (x *HeathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeathResponse.ProtoReflect.Descriptor instead.
func (*HeathResponse) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

func (x *HeathResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_health_proto protoreflect.FileDescriptor

const file_health_proto_rawDesc = "" +
	"\n" +
	"\fhealth.proto\x12\x06health\"\x0e\n" +
	"\fHeathRequest\")\n" +
	"\rHeathResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion2G\n" +
	"\fGreetService\x127\n" +
	"\bSayHello\x12\x14.health.HeathRequest\x1a\x15.health.HeathResponseb\x06proto3"

var (
	file_health_proto_rawDescOnce sync.Once
	file_health_proto_rawDescData []byte
)

func file_health_proto_rawDescGZIP() []byte {
	file_health_proto_rawDescOnce.Do(func() {
		file_health_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_health_proto_rawDesc), len(file_health_proto_rawDesc)))
	})
	return file_health_proto_rawDescData
}

var file_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_health_proto_goTypes = []any{
	(*HeathRequest)(nil),  // 0: health.HeathRequest
	(*HeathResponse)(nil), // 1: health.HeathResponse
}
var file_health_proto_depIdxs = []int32{
	0, // 0: health.GreetService.SayHello:input_type -> health.HeathRequest
	1, // 1: health.GreetService.SayHello:output_type -> health.HeathResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_health_proto_init() }
func file_health_proto_init() {
	if File_health_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_health_proto_rawDesc), len(file_health_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_proto_goTypes,
		DependencyIndexes: file_health_proto_depIdxs,
		MessageInfos:      file_health_proto_msgTypes,
	}.Build()
	File_health_proto = out.File
	file_health_proto_goTypes = nil
	file_health_proto_depIdxs = nil
}
