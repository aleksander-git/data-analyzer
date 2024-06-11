// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: version/v1/common.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_version_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_version_v1_common_proto protoreflect.FileDescriptor

var file_version_v1_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x01, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12, 0x32, 0x10, 0xd0, 0x92, 0xd0, 0xb5,
	0xd1, 0x80, 0xd1, 0x81, 0xd0, 0xb8, 0xd1, 0x8f, 0x20, 0x41, 0x50, 0x49, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0xad, 0x01, 0x92, 0x41, 0xa9, 0x01, 0x0a, 0x57, 0x2a, 0x2e,
	0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8,
	0xd0, 0xb5, 0x20, 0xd0, 0xb2, 0xd0, 0xb5, 0xd1, 0x80, 0xd1, 0x81, 0xd0, 0xb8, 0xd0, 0xb8, 0x20,
	0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb5, 0xd0, 0xba, 0xd1, 0x82, 0xd0, 0xb0, 0x32, 0x25,
	0xd0, 0x92, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x89, 0xd0, 0xb0,
	0xd0, 0xb5, 0xd1, 0x82, 0x20, 0xd0, 0xb2, 0xd0, 0xb5, 0xd1, 0x80, 0xd1, 0x81, 0xd0, 0xb8, 0xd1,
	0x8e, 0x20, 0x41, 0x50, 0x49, 0x2a, 0x4e, 0x0a, 0x1c, 0x41, 0x50, 0x49, 0x20, 0xd0, 0xb4, 0xd0,
	0xbe, 0xd0, 0xba, 0xd1, 0x83, 0xd0, 0xbc, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb0, 0xd1,
	0x86, 0xd0, 0xb8, 0xd1, 0x8f, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65,
	0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0xad, 0x01, 0x0a, 0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x92, 0x41, 0x46, 0x0a,
	0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x25, 0xd0, 0x92, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1,
	0x80, 0xd0, 0xb0, 0xd1, 0x89, 0xd0, 0xb0, 0xd0, 0xb5, 0xd1, 0x82, 0x20, 0xd0, 0xb2, 0xd0, 0xb5,
	0xd1, 0x80, 0xd1, 0x81, 0xd0, 0xb8, 0xd1, 0x8e, 0x20, 0x41, 0x50, 0x49, 0x2a, 0x0a, 0x67, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12,
	0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xfa, 0x02, 0x92, 0x41, 0xb9, 0x02, 0x12, 0x8f, 0x02, 0x0a,
	0x08, 0x45, 0x63, 0x68, 0x6f, 0x20, 0x41, 0x50, 0x49, 0x12, 0x3a, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
	0x81, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xbf, 0xd1, 0x80,
	0xd0, 0xb8, 0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x80, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0,
	0xbe, 0xd1, 0x82, 0xd1, 0x8b, 0x20, 0xd1, 0x81, 0x20, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0xa7, 0x01, 0x0a, 0x73, 0xd0, 0x9f, 0xd1, 0x80, 0xd0, 0xbe,
	0xd0, 0xb5, 0xd0, 0xba, 0xd1, 0x82, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0x20, 0xd0, 0xb7, 0xd0, 0xb0,
	0xd0, 0xb3, 0xd1, 0x80, 0xd1, 0x83, 0xd0, 0xb7, 0xd0, 0xba, 0xd0, 0xb5, 0x20, 0xd0, 0xb4, 0xd0,
	0xb0, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8b, 0xd1, 0x85, 0x20, 0xd1, 0x81, 0x20, 0x63, 0x73, 0x76,
	0x20, 0xd1, 0x84, 0xd0, 0xb0, 0xd0, 0xb9, 0xd0, 0xbb, 0xd0, 0xb0, 0x20, 0xd0, 0xb8, 0x20, 0xd0,
	0xb4, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xb9, 0xd1, 0x88, 0xd0,
	0xb5, 0xd0, 0xb9, 0x20, 0xd0, 0xb5, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xbe, 0xd0, 0xb1, 0xd1,
	0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x82, 0xd0, 0xba, 0xd0, 0xb8, 0x12, 0x30, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x6b, 0x73, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x2d, 0x67, 0x69, 0x74,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x2a,
	0x16, 0x0a, 0x14, 0x42, 0x53, 0x44, 0x20, 0x33, 0x2d, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x20,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x2a, 0x01,
	0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x6b, 0x73, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x2d, 0x67, 0x69, 0x74,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_version_v1_common_proto_rawDescOnce sync.Once
	file_version_v1_common_proto_rawDescData = file_version_v1_common_proto_rawDesc
)

func file_version_v1_common_proto_rawDescGZIP() []byte {
	file_version_v1_common_proto_rawDescOnce.Do(func() {
		file_version_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_version_v1_common_proto_rawDescData)
	})
	return file_version_v1_common_proto_rawDescData
}

var file_version_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_version_v1_common_proto_goTypes = []interface{}{
	(*VersionResponse)(nil), // 0: version.v1.VersionResponse
	(*emptypb.Empty)(nil),   // 1: google.protobuf.Empty
}
var file_version_v1_common_proto_depIdxs = []int32{
	1, // 0: version.v1.Microservice.GetVersion:input_type -> google.protobuf.Empty
	0, // 1: version.v1.Microservice.GetVersion:output_type -> version.v1.VersionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_version_v1_common_proto_init() }
func file_version_v1_common_proto_init() {
	if File_version_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_version_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_version_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_version_v1_common_proto_goTypes,
		DependencyIndexes: file_version_v1_common_proto_depIdxs,
		MessageInfos:      file_version_v1_common_proto_msgTypes,
	}.Build()
	File_version_v1_common_proto = out.File
	file_version_v1_common_proto_rawDesc = nil
	file_version_v1_common_proto_goTypes = nil
	file_version_v1_common_proto_depIdxs = nil
}