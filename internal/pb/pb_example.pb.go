// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: internal/pb/pb_example.proto

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

type PbExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Value int64 `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *PbExample) Reset() {
	*x = PbExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_pb_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbExample) ProtoMessage() {}

func (x *PbExample) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_pb_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbExample.ProtoReflect.Descriptor instead.
func (*PbExample) Descriptor() ([]byte, []int) {
	return file_internal_pb_pb_example_proto_rawDescGZIP(), []int{0}
}

func (x *PbExample) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PbExample) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_internal_pb_pb_example_proto protoreflect.FileDescriptor

var file_internal_pb_pb_example_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62,
	0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x31, 0x0a, 0x09, 0x50, 0x62, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_pb_example_proto_rawDescOnce sync.Once
	file_internal_pb_pb_example_proto_rawDescData = file_internal_pb_pb_example_proto_rawDesc
)

func file_internal_pb_pb_example_proto_rawDescGZIP() []byte {
	file_internal_pb_pb_example_proto_rawDescOnce.Do(func() {
		file_internal_pb_pb_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_pb_example_proto_rawDescData)
	})
	return file_internal_pb_pb_example_proto_rawDescData
}

var file_internal_pb_pb_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_pb_pb_example_proto_goTypes = []interface{}{
	(*PbExample)(nil), // 0: binary.PbExample
}
var file_internal_pb_pb_example_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pb_pb_example_proto_init() }
func file_internal_pb_pb_example_proto_init() {
	if File_internal_pb_pb_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_pb_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbExample); i {
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
			RawDescriptor: file_internal_pb_pb_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_pb_example_proto_goTypes,
		DependencyIndexes: file_internal_pb_pb_example_proto_depIdxs,
		MessageInfos:      file_internal_pb_pb_example_proto_msgTypes,
	}.Build()
	File_internal_pb_pb_example_proto = out.File
	file_internal_pb_pb_example_proto_rawDesc = nil
	file_internal_pb_pb_example_proto_goTypes = nil
	file_internal_pb_pb_example_proto_depIdxs = nil
}
