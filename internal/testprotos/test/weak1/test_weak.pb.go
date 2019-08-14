// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/weak1/test_weak.proto

package weak1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type WeakImportMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
}

func (x *WeakImportMessage1) Reset() {
	*x = WeakImportMessage1{}
}

func (x *WeakImportMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeakImportMessage1) ProtoMessage() {}

func (x *WeakImportMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_test_weak1_test_weak_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeakImportMessage1.ProtoReflect.Descriptor instead.
func (*WeakImportMessage1) Descriptor() ([]byte, []int) {
	return file_test_weak1_test_weak_proto_rawDescGZIP(), []int{0}
}

func (x *WeakImportMessage1) GetA() int32 {
	if x != nil && x.A != nil {
		return *x.A
	}
	return 0
}

var File_test_weak1_test_weak_proto protoreflect.FileDescriptor

var file_test_weak1_test_weak_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x77, 0x65, 0x61, 0x6b, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x77, 0x65, 0x61, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x77, 0x65, 0x61, 0x6b, 0x22, 0x22, 0x0a, 0x12, 0x57, 0x65, 0x61, 0x6b, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x77, 0x65, 0x61, 0x6b, 0x31,
}

var (
	file_test_weak1_test_weak_proto_rawDescOnce sync.Once
	file_test_weak1_test_weak_proto_rawDescData = file_test_weak1_test_weak_proto_rawDesc
)

func file_test_weak1_test_weak_proto_rawDescGZIP() []byte {
	file_test_weak1_test_weak_proto_rawDescOnce.Do(func() {
		file_test_weak1_test_weak_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_weak1_test_weak_proto_rawDescData)
	})
	return file_test_weak1_test_weak_proto_rawDescData
}

var file_test_weak1_test_weak_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_weak1_test_weak_proto_goTypes = []interface{}{
	(*WeakImportMessage1)(nil), // 0: goproto.proto.test.weak.WeakImportMessage1
}
var file_test_weak1_test_weak_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_weak1_test_weak_proto_init() }
func file_test_weak1_test_weak_proto_init() {
	if File_test_weak1_test_weak_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_weak1_test_weak_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeakImportMessage1); i {
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
			RawDescriptor: file_test_weak1_test_weak_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_weak1_test_weak_proto_goTypes,
		DependencyIndexes: file_test_weak1_test_weak_proto_depIdxs,
		MessageInfos:      file_test_weak1_test_weak_proto_msgTypes,
	}.Build()
	File_test_weak1_test_weak_proto = out.File
	file_test_weak1_test_weak_proto_rawDesc = nil
	file_test_weak1_test_weak_proto_goTypes = nil
	file_test_weak1_test_weak_proto_depIdxs = nil
}
