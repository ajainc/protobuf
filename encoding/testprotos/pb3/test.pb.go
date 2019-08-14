// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test Protobuf definitions with proto3 syntax.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb3/test.proto

package pb3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
	Enum_TEN  Enum = 10
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0:  "ZERO",
		1:  "ONE",
		2:  "TWO",
		10: "TEN",
	}
	Enum_value = map[string]int32{
		"ZERO": 0,
		"ONE":  1,
		"TWO":  2,
		"TEN":  10,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_pb3_test_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_pb3_test_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{0}
}

type Enums_NestedEnum int32

const (
	Enums_CERO Enums_NestedEnum = 0
	Enums_UNO  Enums_NestedEnum = 1
	Enums_DOS  Enums_NestedEnum = 2
	Enums_DIEZ Enums_NestedEnum = 10
)

// Enum value maps for Enums_NestedEnum.
var (
	Enums_NestedEnum_name = map[int32]string{
		0:  "CERO",
		1:  "UNO",
		2:  "DOS",
		10: "DIEZ",
	}
	Enums_NestedEnum_value = map[string]int32{
		"CERO": 0,
		"UNO":  1,
		"DOS":  2,
		"DIEZ": 10,
	}
)

func (x Enums_NestedEnum) Enum() *Enums_NestedEnum {
	p := new(Enums_NestedEnum)
	*p = x
	return p
}

func (x Enums_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enums_NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_pb3_test_proto_enumTypes[1].Descriptor()
}

func (Enums_NestedEnum) Type() protoreflect.EnumType {
	return &file_pb3_test_proto_enumTypes[1]
}

func (x Enums_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enums_NestedEnum.Descriptor instead.
func (Enums_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{1, 0}
}

// Scalars contains scalar field types.
type Scalars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SBool     bool    `protobuf:"varint,1,opt,name=s_bool,json=sBool,proto3" json:"s_bool,omitempty"`
	SInt32    int32   `protobuf:"varint,2,opt,name=s_int32,json=sInt32,proto3" json:"s_int32,omitempty"`
	SInt64    int64   `protobuf:"varint,3,opt,name=s_int64,json=sInt64,proto3" json:"s_int64,omitempty"`
	SUint32   uint32  `protobuf:"varint,4,opt,name=s_uint32,json=sUint32,proto3" json:"s_uint32,omitempty"`
	SUint64   uint64  `protobuf:"varint,5,opt,name=s_uint64,json=sUint64,proto3" json:"s_uint64,omitempty"`
	SSint32   int32   `protobuf:"zigzag32,6,opt,name=s_sint32,json=sSint32,proto3" json:"s_sint32,omitempty"`
	SSint64   int64   `protobuf:"zigzag64,7,opt,name=s_sint64,json=sSint64,proto3" json:"s_sint64,omitempty"`
	SFixed32  uint32  `protobuf:"fixed32,8,opt,name=s_fixed32,json=sFixed32,proto3" json:"s_fixed32,omitempty"`
	SFixed64  uint64  `protobuf:"fixed64,9,opt,name=s_fixed64,json=sFixed64,proto3" json:"s_fixed64,omitempty"`
	SSfixed32 int32   `protobuf:"fixed32,10,opt,name=s_sfixed32,json=sSfixed32,proto3" json:"s_sfixed32,omitempty"`
	SSfixed64 int64   `protobuf:"fixed64,11,opt,name=s_sfixed64,json=sSfixed64,proto3" json:"s_sfixed64,omitempty"`
	SFloat    float32 `protobuf:"fixed32,20,opt,name=s_float,json=sFloat,proto3" json:"s_float,omitempty"`
	SDouble   float64 `protobuf:"fixed64,21,opt,name=s_double,json=sDouble,proto3" json:"s_double,omitempty"`
	SBytes    []byte  `protobuf:"bytes,14,opt,name=s_bytes,json=sBytes,proto3" json:"s_bytes,omitempty"`
	SString   string  `protobuf:"bytes,13,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
}

func (x *Scalars) Reset() {
	*x = Scalars{}
}

func (x *Scalars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalars) ProtoMessage() {}

func (x *Scalars) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalars.ProtoReflect.Descriptor instead.
func (*Scalars) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{0}
}

func (x *Scalars) GetSBool() bool {
	if x != nil {
		return x.SBool
	}
	return false
}

func (x *Scalars) GetSInt32() int32 {
	if x != nil {
		return x.SInt32
	}
	return 0
}

func (x *Scalars) GetSInt64() int64 {
	if x != nil {
		return x.SInt64
	}
	return 0
}

func (x *Scalars) GetSUint32() uint32 {
	if x != nil {
		return x.SUint32
	}
	return 0
}

func (x *Scalars) GetSUint64() uint64 {
	if x != nil {
		return x.SUint64
	}
	return 0
}

func (x *Scalars) GetSSint32() int32 {
	if x != nil {
		return x.SSint32
	}
	return 0
}

func (x *Scalars) GetSSint64() int64 {
	if x != nil {
		return x.SSint64
	}
	return 0
}

func (x *Scalars) GetSFixed32() uint32 {
	if x != nil {
		return x.SFixed32
	}
	return 0
}

func (x *Scalars) GetSFixed64() uint64 {
	if x != nil {
		return x.SFixed64
	}
	return 0
}

func (x *Scalars) GetSSfixed32() int32 {
	if x != nil {
		return x.SSfixed32
	}
	return 0
}

func (x *Scalars) GetSSfixed64() int64 {
	if x != nil {
		return x.SSfixed64
	}
	return 0
}

func (x *Scalars) GetSFloat() float32 {
	if x != nil {
		return x.SFloat
	}
	return 0
}

func (x *Scalars) GetSDouble() float64 {
	if x != nil {
		return x.SDouble
	}
	return 0
}

func (x *Scalars) GetSBytes() []byte {
	if x != nil {
		return x.SBytes
	}
	return nil
}

func (x *Scalars) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

// Message contains enum fields.
type Enums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SEnum       Enum             `protobuf:"varint,1,opt,name=s_enum,json=sEnum,proto3,enum=pb3.Enum" json:"s_enum,omitempty"`
	SNestedEnum Enums_NestedEnum `protobuf:"varint,3,opt,name=s_nested_enum,json=sNestedEnum,proto3,enum=pb3.Enums_NestedEnum" json:"s_nested_enum,omitempty"`
}

func (x *Enums) Reset() {
	*x = Enums{}
}

func (x *Enums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enums) ProtoMessage() {}

func (x *Enums) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enums.ProtoReflect.Descriptor instead.
func (*Enums) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{1}
}

func (x *Enums) GetSEnum() Enum {
	if x != nil {
		return x.SEnum
	}
	return Enum_ZERO
}

func (x *Enums) GetSNestedEnum() Enums_NestedEnum {
	if x != nil {
		return x.SNestedEnum
	}
	return Enums_CERO
}

// Message contains nested message field.
type Nests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SNested *Nested `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
}

func (x *Nests) Reset() {
	*x = Nests{}
}

func (x *Nests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nests) ProtoMessage() {}

func (x *Nests) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nests.ProtoReflect.Descriptor instead.
func (*Nests) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{2}
}

func (x *Nests) GetSNested() *Nested {
	if x != nil {
		return x.SNested
	}
	return nil
}

// Message type used as submessage.
type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SString string  `protobuf:"bytes,1,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	SNested *Nested `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
}

func (x *Nested) Reset() {
	*x = Nested{}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{3}
}

func (x *Nested) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

func (x *Nested) GetSNested() *Nested {
	if x != nil {
		return x.SNested
	}
	return nil
}

// Message contains oneof field.
type Oneofs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Union:
	//	*Oneofs_OneofEnum
	//	*Oneofs_OneofString
	//	*Oneofs_OneofNested
	Union isOneofs_Union `protobuf_oneof:"union"`
}

func (x *Oneofs) Reset() {
	*x = Oneofs{}
}

func (x *Oneofs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oneofs) ProtoMessage() {}

func (x *Oneofs) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oneofs.ProtoReflect.Descriptor instead.
func (*Oneofs) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{4}
}

func (m *Oneofs) GetUnion() isOneofs_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (x *Oneofs) GetOneofEnum() Enum {
	if x, ok := x.GetUnion().(*Oneofs_OneofEnum); ok {
		return x.OneofEnum
	}
	return Enum_ZERO
}

func (x *Oneofs) GetOneofString() string {
	if x, ok := x.GetUnion().(*Oneofs_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (x *Oneofs) GetOneofNested() *Nested {
	if x, ok := x.GetUnion().(*Oneofs_OneofNested); ok {
		return x.OneofNested
	}
	return nil
}

type isOneofs_Union interface {
	isOneofs_Union()
}

type Oneofs_OneofEnum struct {
	OneofEnum Enum `protobuf:"varint,1,opt,name=oneof_enum,json=oneofEnum,proto3,enum=pb3.Enum,oneof"`
}

type Oneofs_OneofString struct {
	OneofString string `protobuf:"bytes,2,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type Oneofs_OneofNested struct {
	OneofNested *Nested `protobuf:"bytes,3,opt,name=oneof_nested,json=oneofNested,proto3,oneof"`
}

func (*Oneofs_OneofEnum) isOneofs_Union() {}

func (*Oneofs_OneofString) isOneofs_Union() {}

func (*Oneofs_OneofNested) isOneofs_Union() {}

// Message contains map fields.
type Maps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32ToStr   map[int32]string   `protobuf:"bytes,1,rep,name=int32_to_str,json=int32ToStr,proto3" json:"int32_to_str,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BoolToUint32 map[bool]uint32    `protobuf:"bytes,2,rep,name=bool_to_uint32,json=boolToUint32,proto3" json:"bool_to_uint32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Uint64ToEnum map[uint64]Enum    `protobuf:"bytes,3,rep,name=uint64_to_enum,json=uint64ToEnum,proto3" json:"uint64_to_enum,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=pb3.Enum"`
	StrToNested  map[string]*Nested `protobuf:"bytes,4,rep,name=str_to_nested,json=strToNested,proto3" json:"str_to_nested,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StrToOneofs  map[string]*Oneofs `protobuf:"bytes,5,rep,name=str_to_oneofs,json=strToOneofs,proto3" json:"str_to_oneofs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Maps) Reset() {
	*x = Maps{}
}

func (x *Maps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maps) ProtoMessage() {}

func (x *Maps) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maps.ProtoReflect.Descriptor instead.
func (*Maps) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{5}
}

func (x *Maps) GetInt32ToStr() map[int32]string {
	if x != nil {
		return x.Int32ToStr
	}
	return nil
}

func (x *Maps) GetBoolToUint32() map[bool]uint32 {
	if x != nil {
		return x.BoolToUint32
	}
	return nil
}

func (x *Maps) GetUint64ToEnum() map[uint64]Enum {
	if x != nil {
		return x.Uint64ToEnum
	}
	return nil
}

func (x *Maps) GetStrToNested() map[string]*Nested {
	if x != nil {
		return x.StrToNested
	}
	return nil
}

func (x *Maps) GetStrToOneofs() map[string]*Oneofs {
	if x != nil {
		return x.StrToOneofs
	}
	return nil
}

// Message for testing json_name option.
type JSONNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SString string `protobuf:"bytes,1,opt,name=s_string,json=foo_bar,proto3" json:"s_string,omitempty"`
}

func (x *JSONNames) Reset() {
	*x = JSONNames{}
}

func (x *JSONNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONNames) ProtoMessage() {}

func (x *JSONNames) ProtoReflect() protoreflect.Message {
	mi := &file_pb3_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONNames.ProtoReflect.Descriptor instead.
func (*JSONNames) Descriptor() ([]byte, []int) {
	return file_pb3_test_proto_rawDescGZIP(), []int{6}
}

func (x *JSONNames) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

var File_pb3_test_proto protoreflect.FileDescriptor

var file_pb3_test_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x62, 0x33, 0x22, 0x9e, 0x03, 0x0a, 0x07, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x73, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x73, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x07, 0x73, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x73,
	0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x07, 0x52, 0x08, 0x73, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x73, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x09, 0x73, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x10, 0x52, 0x09, 0x73, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x73, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x98, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x12, 0x20, 0x0a, 0x06, 0x73, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x39, 0x0a, 0x0d, 0x73, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x0b, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x32, 0x0a,
	0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x43,
	0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x4e, 0x4f, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x44, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x45, 0x5a, 0x10,
	0x0a, 0x22, 0x2f, 0x0a, 0x05, 0x4e, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x73, 0x5f,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x07, 0x73, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x22, 0x4b, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x08, 0x73, 0x5f, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x07, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22,
	0x94, 0x01, 0x0a, 0x06, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x0c, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x05, 0x0a, 0x04, 0x4d, 0x61, 0x70, 0x73, 0x12,
	0x3b, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4d, 0x61, 0x70, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x12, 0x41, 0x0a, 0x0e,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x41, 0x0a, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4d, 0x61,
	0x70, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x11, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x45, 0x6e,
	0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b,
	0x0a, 0x10, 0x53, 0x74, 0x72, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x10, 0x53,
	0x74, 0x72, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x09, 0x4a, 0x53, 0x4f, 0x4e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x6f, 0x5f, 0x62, 0x61, 0x72,
	0x2a, 0x2b, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x57, 0x4f, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x45, 0x4e, 0x10, 0x0a, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x70, 0x62, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb3_test_proto_rawDescOnce sync.Once
	file_pb3_test_proto_rawDescData = file_pb3_test_proto_rawDesc
)

func file_pb3_test_proto_rawDescGZIP() []byte {
	file_pb3_test_proto_rawDescOnce.Do(func() {
		file_pb3_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb3_test_proto_rawDescData)
	})
	return file_pb3_test_proto_rawDescData
}

var file_pb3_test_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb3_test_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pb3_test_proto_goTypes = []interface{}{
	(Enum)(0),             // 0: pb3.Enum
	(Enums_NestedEnum)(0), // 1: pb3.Enums.NestedEnum
	(*Scalars)(nil),       // 2: pb3.Scalars
	(*Enums)(nil),         // 3: pb3.Enums
	(*Nests)(nil),         // 4: pb3.Nests
	(*Nested)(nil),        // 5: pb3.Nested
	(*Oneofs)(nil),        // 6: pb3.Oneofs
	(*Maps)(nil),          // 7: pb3.Maps
	(*JSONNames)(nil),     // 8: pb3.JSONNames
	nil,                   // 9: pb3.Maps.Int32ToStrEntry
	nil,                   // 10: pb3.Maps.BoolToUint32Entry
	nil,                   // 11: pb3.Maps.Uint64ToEnumEntry
	nil,                   // 12: pb3.Maps.StrToNestedEntry
	nil,                   // 13: pb3.Maps.StrToOneofsEntry
}
var file_pb3_test_proto_depIdxs = []int32{
	0,  // 0: pb3.Enums.s_enum:type_name -> pb3.Enum
	1,  // 1: pb3.Enums.s_nested_enum:type_name -> pb3.Enums.NestedEnum
	5,  // 2: pb3.Nests.s_nested:type_name -> pb3.Nested
	5,  // 3: pb3.Nested.s_nested:type_name -> pb3.Nested
	0,  // 4: pb3.Oneofs.oneof_enum:type_name -> pb3.Enum
	5,  // 5: pb3.Oneofs.oneof_nested:type_name -> pb3.Nested
	9,  // 6: pb3.Maps.int32_to_str:type_name -> pb3.Maps.Int32ToStrEntry
	10, // 7: pb3.Maps.bool_to_uint32:type_name -> pb3.Maps.BoolToUint32Entry
	11, // 8: pb3.Maps.uint64_to_enum:type_name -> pb3.Maps.Uint64ToEnumEntry
	12, // 9: pb3.Maps.str_to_nested:type_name -> pb3.Maps.StrToNestedEntry
	13, // 10: pb3.Maps.str_to_oneofs:type_name -> pb3.Maps.StrToOneofsEntry
	0,  // 11: pb3.Maps.Uint64ToEnumEntry.value:type_name -> pb3.Enum
	5,  // 12: pb3.Maps.StrToNestedEntry.value:type_name -> pb3.Nested
	6,  // 13: pb3.Maps.StrToOneofsEntry.value:type_name -> pb3.Oneofs
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_pb3_test_proto_init() }
func file_pb3_test_proto_init() {
	if File_pb3_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb3_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalars); i {
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
		file_pb3_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enums); i {
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
		file_pb3_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nests); i {
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
		file_pb3_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
		file_pb3_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oneofs); i {
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
		file_pb3_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maps); i {
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
		file_pb3_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONNames); i {
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
	file_pb3_test_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Oneofs_OneofEnum)(nil),
		(*Oneofs_OneofString)(nil),
		(*Oneofs_OneofNested)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb3_test_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb3_test_proto_goTypes,
		DependencyIndexes: file_pb3_test_proto_depIdxs,
		EnumInfos:         file_pb3_test_proto_enumTypes,
		MessageInfos:      file_pb3_test_proto_msgTypes,
	}.Build()
	File_pb3_test_proto = out.File
	file_pb3_test_proto_rawDesc = nil
	file_pb3_test_proto_goTypes = nil
	file_pb3_test_proto_depIdxs = nil
}
