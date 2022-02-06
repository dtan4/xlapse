// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: types/v1/gif_request.proto

package v1

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

type GifRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket    string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	KeyPrefix string `protobuf:"bytes,2,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix,omitempty"`
	Year      int32  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Month     int32  `protobuf:"varint,4,opt,name=month,proto3" json:"month,omitempty"`
	Day       int32  `protobuf:"varint,5,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *GifRequest) Reset() {
	*x = GifRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_gif_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GifRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GifRequest) ProtoMessage() {}

func (x *GifRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_gif_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GifRequest.ProtoReflect.Descriptor instead.
func (*GifRequest) Descriptor() ([]byte, []int) {
	return file_types_v1_gif_request_proto_rawDescGZIP(), []int{0}
}

func (x *GifRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GifRequest) GetKeyPrefix() string {
	if x != nil {
		return x.KeyPrefix
	}
	return ""
}

func (x *GifRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GifRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *GifRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

var File_types_v1_gif_request_proto protoreflect.FileDescriptor

var file_types_v1_gif_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x66, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x78, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x7f, 0x0a,
	0x0a, 0x47, 0x69, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x74, 0x61,
	0x6e, 0x34, 0x2f, 0x78, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_gif_request_proto_rawDescOnce sync.Once
	file_types_v1_gif_request_proto_rawDescData = file_types_v1_gif_request_proto_rawDesc
)

func file_types_v1_gif_request_proto_rawDescGZIP() []byte {
	file_types_v1_gif_request_proto_rawDescOnce.Do(func() {
		file_types_v1_gif_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_gif_request_proto_rawDescData)
	})
	return file_types_v1_gif_request_proto_rawDescData
}

var file_types_v1_gif_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_v1_gif_request_proto_goTypes = []interface{}{
	(*GifRequest)(nil), // 0: xlapse.types.v1.GifRequest
}
var file_types_v1_gif_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_v1_gif_request_proto_init() }
func file_types_v1_gif_request_proto_init() {
	if File_types_v1_gif_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_v1_gif_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GifRequest); i {
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
			RawDescriptor: file_types_v1_gif_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_gif_request_proto_goTypes,
		DependencyIndexes: file_types_v1_gif_request_proto_depIdxs,
		MessageInfos:      file_types_v1_gif_request_proto_msgTypes,
	}.Build()
	File_types_v1_gif_request_proto = out.File
	file_types_v1_gif_request_proto_rawDesc = nil
	file_types_v1_gif_request_proto_goTypes = nil
	file_types_v1_gif_request_proto_depIdxs = nil
}
