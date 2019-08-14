// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extensions/extra/extra.proto

package extra

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type ExtraMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (x *ExtraMessage) Reset() {
	*x = ExtraMessage{}
}

func (x *ExtraMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraMessage) ProtoMessage() {}

func (x *ExtraMessage) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_extra_extra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraMessage.ProtoReflect.Descriptor instead.
func (*ExtraMessage) Descriptor() ([]byte, []int) {
	return file_extensions_extra_extra_proto_rawDescGZIP(), []int{0}
}

func (x *ExtraMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_extensions_extra_extra_proto protoreflect.FileDescriptor

var file_extensions_extra_extra_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x22,
	0x0a, 0x0c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61,
}

var (
	file_extensions_extra_extra_proto_rawDescOnce sync.Once
	file_extensions_extra_extra_proto_rawDescData = file_extensions_extra_extra_proto_rawDesc
)

func file_extensions_extra_extra_proto_rawDescGZIP() []byte {
	file_extensions_extra_extra_proto_rawDescOnce.Do(func() {
		file_extensions_extra_extra_proto_rawDescData = protoimpl.X.CompressGZIP(file_extensions_extra_extra_proto_rawDescData)
	})
	return file_extensions_extra_extra_proto_rawDescData
}

var file_extensions_extra_extra_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_extensions_extra_extra_proto_goTypes = []interface{}{
	(*ExtraMessage)(nil), // 0: goproto.protoc.extension.extra.ExtraMessage
}
var file_extensions_extra_extra_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_extensions_extra_extra_proto_init() }
func file_extensions_extra_extra_proto_init() {
	if File_extensions_extra_extra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extensions_extra_extra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraMessage); i {
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
			RawDescriptor: file_extensions_extra_extra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extensions_extra_extra_proto_goTypes,
		DependencyIndexes: file_extensions_extra_extra_proto_depIdxs,
		MessageInfos:      file_extensions_extra_extra_proto_msgTypes,
	}.Build()
	File_extensions_extra_extra_proto = out.File
	file_extensions_extra_extra_proto_rawDesc = nil
	file_extensions_extra_extra_proto_goTypes = nil
	file_extensions_extra_extra_proto_depIdxs = nil
}
