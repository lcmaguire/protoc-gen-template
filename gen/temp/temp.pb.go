// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: temp/temp.proto

package temp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data int32

const (
	Data_DATA_UNSPECIFIED Data = 0
	Data_DATA_SPECIFIED   Data = 1
)

// Enum value maps for Data.
var (
	Data_name = map[int32]string{
		0: "DATA_UNSPECIFIED",
		1: "DATA_SPECIFIED",
	}
	Data_value = map[string]int32{
		"DATA_UNSPECIFIED": 0,
		"DATA_SPECIFIED":   1,
	}
)

func (x Data) Enum() *Data {
	p := new(Data)
	*p = x
	return p
}

func (x Data) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Data) Descriptor() protoreflect.EnumDescriptor {
	return file_temp_temp_proto_enumTypes[0].Descriptor()
}

func (Data) Type() protoreflect.EnumType {
	return &file_temp_temp_proto_enumTypes[0]
}

func (x Data) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Data.Descriptor instead.
func (Data) EnumDescriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{0}
}

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// standard types
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count  int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Active bool   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	// repeated example (consider an append function)
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// extra message
	Foo *Foo         `protobuf:"bytes,5,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar *Example_Bar `protobuf:"bytes,6,opt,name=bar,proto3" json:"bar,omitempty"`
	// imported message
	Any *anypb.Any `protobuf:"bytes,7,opt,name=any,proto3" json:"any,omitempty"`
	// enum
	Data Data `protobuf:"varint,8,opt,name=data,proto3,enum=proto.Data" json:"data,omitempty"`
	// optional field
	ExtraComments *string `protobuf:"bytes,9,opt,name=extra_comments,json=extraComments,proto3,oneof" json:"extra_comments,omitempty"`
	// map (consider a SetKey, Value)
	FooMap map[string]*Foo `protobuf:"bytes,10,rep,name=foo_map,json=fooMap,proto3" json:"foo_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// oneofs
	Sample *SampleMessage `protobuf:"bytes,11,opt,name=sample,proto3" json:"sample,omitempty"`
	// Types that are assignable to AbcOneof:
	//
	//	*Example_Abc
	//	*Example_Far_
	AbcOneof isExample_AbcOneof `protobuf_oneof:"abc_oneof"`
	Bites    [][]byte           `protobuf:"bytes,14,rep,name=bites,proto3" json:"bites,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_temp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_temp_temp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Example) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Example) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Example) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Example) GetFoo() *Foo {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *Example) GetBar() *Example_Bar {
	if x != nil {
		return x.Bar
	}
	return nil
}

func (x *Example) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *Example) GetData() Data {
	if x != nil {
		return x.Data
	}
	return Data_DATA_UNSPECIFIED
}

func (x *Example) GetExtraComments() string {
	if x != nil && x.ExtraComments != nil {
		return *x.ExtraComments
	}
	return ""
}

func (x *Example) GetFooMap() map[string]*Foo {
	if x != nil {
		return x.FooMap
	}
	return nil
}

func (x *Example) GetSample() *SampleMessage {
	if x != nil {
		return x.Sample
	}
	return nil
}

func (m *Example) GetAbcOneof() isExample_AbcOneof {
	if m != nil {
		return m.AbcOneof
	}
	return nil
}

func (x *Example) GetAbc() string {
	if x, ok := x.GetAbcOneof().(*Example_Abc); ok {
		return x.Abc
	}
	return ""
}

func (x *Example) GetFar() *Example_Far {
	if x, ok := x.GetAbcOneof().(*Example_Far_); ok {
		return x.Far
	}
	return nil
}

func (x *Example) GetBites() [][]byte {
	if x != nil {
		return x.Bites
	}
	return nil
}

type isExample_AbcOneof interface {
	isExample_AbcOneof()
}

type Example_Abc struct {
	Abc string `protobuf:"bytes,12,opt,name=abc,proto3,oneof"`
}

type Example_Far_ struct {
	Far *Example_Far `protobuf:"bytes,13,opt,name=far,proto3,oneof"`
}

func (*Example_Abc) isExample_AbcOneof() {}

func (*Example_Far_) isExample_AbcOneof() {}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_temp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_temp_temp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{1}
}

func (x *Foo) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Funk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Funk) Reset() {
	*x = Funk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_temp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Funk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Funk) ProtoMessage() {}

func (x *Funk) ProtoReflect() protoreflect.Message {
	mi := &file_temp_temp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Funk.ProtoReflect.Descriptor instead.
func (*Funk) Descriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{2}
}

func (x *Funk) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TestOneof:
	//
	//	*SampleMessage_Name
	//	*SampleMessage_Foo
	//	*SampleMessage_Funk
	TestOneof isSampleMessage_TestOneof `protobuf_oneof:"test_oneof"`
}

func (x *SampleMessage) Reset() {
	*x = SampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_temp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleMessage) ProtoMessage() {}

func (x *SampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_temp_temp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleMessage.ProtoReflect.Descriptor instead.
func (*SampleMessage) Descriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{3}
}

func (m *SampleMessage) GetTestOneof() isSampleMessage_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *SampleMessage) GetName() string {
	if x, ok := x.GetTestOneof().(*SampleMessage_Name); ok {
		return x.Name
	}
	return ""
}

func (x *SampleMessage) GetFoo() *Foo {
	if x, ok := x.GetTestOneof().(*SampleMessage_Foo); ok {
		return x.Foo
	}
	return nil
}

func (x *SampleMessage) GetFunk() *Funk {
	if x, ok := x.GetTestOneof().(*SampleMessage_Funk); ok {
		return x.Funk
	}
	return nil
}

type isSampleMessage_TestOneof interface {
	isSampleMessage_TestOneof()
}

type SampleMessage_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type SampleMessage_Foo struct {
	Foo *Foo `protobuf:"bytes,2,opt,name=foo,proto3,oneof"`
}

type SampleMessage_Funk struct {
	Funk *Funk `protobuf:"bytes,3,opt,name=funk,proto3,oneof"`
}

func (*SampleMessage_Name) isSampleMessage_TestOneof() {}

func (*SampleMessage_Foo) isSampleMessage_TestOneof() {}

func (*SampleMessage_Funk) isSampleMessage_TestOneof() {}

// nested message.
type Example_Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nested string `protobuf:"bytes,1,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *Example_Bar) Reset() {
	*x = Example_Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_temp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example_Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example_Bar) ProtoMessage() {}

func (x *Example_Bar) ProtoReflect() protoreflect.Message {
	mi := &file_temp_temp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example_Bar.ProtoReflect.Descriptor instead.
func (*Example_Bar) Descriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Example_Bar) GetNested() string {
	if x != nil {
		return x.Nested
	}
	return ""
}

type Example_Far struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Example_Far) Reset() {
	*x = Example_Far{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_temp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example_Far) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example_Far) ProtoMessage() {}

func (x *Example_Far) ProtoReflect() protoreflect.Message {
	mi := &file_temp_temp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example_Far.ProtoReflect.Descriptor instead.
func (*Example_Far) Descriptor() ([]byte, []int) {
	return file_temp_temp_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Example_Far) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

var File_temp_temp_proto protoreflect.FileDescriptor

var file_temp_temp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x04, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x03,
	0x66, 0x6f, 0x6f, 0x12, 0x24, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x42, 0x61, 0x72, 0x52, 0x03, 0x62, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x6e, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x61, 0x6e,
	0x79, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2a, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x33,
	0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x46, 0x6f, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x6f, 0x6f,
	0x4d, 0x61, 0x70, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x03, 0x61, 0x62, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x62, 0x63, 0x12, 0x26, 0x0a, 0x03, 0x66, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x46, 0x61, 0x72, 0x48, 0x00, 0x52, 0x03, 0x66, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x69,
	0x74, 0x65, 0x73, 0x1a, 0x1d, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x1a, 0x45, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x1d, 0x0a, 0x03, 0x46, 0x61, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x61, 0x62, 0x63, 0x5f,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x1b, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1c, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x66, 0x6f,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x6f, 0x6f, 0x48, 0x00, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x75,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x6b, 0x42, 0x0c, 0x0a,
	0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2a, 0x30, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x01, 0x32, 0x9c, 0x02,
	0x0a, 0x0a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x50, 0x49, 0x12, 0x2c, 0x0a, 0x0a,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x70, 0x63, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x6e, 0x79, 0x52, 0x70, 0x63, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x12, 0x37, 0x0a, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x28, 0x01, 0x12, 0x37, 0x0a, 0x13, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x69,
	0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x7d, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x63, 0x6d, 0x61, 0x67, 0x75, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2,
	0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_temp_temp_proto_rawDescOnce sync.Once
	file_temp_temp_proto_rawDescData = file_temp_temp_proto_rawDesc
)

func file_temp_temp_proto_rawDescGZIP() []byte {
	file_temp_temp_proto_rawDescOnce.Do(func() {
		file_temp_temp_proto_rawDescData = protoimpl.X.CompressGZIP(file_temp_temp_proto_rawDescData)
	})
	return file_temp_temp_proto_rawDescData
}

var file_temp_temp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temp_temp_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_temp_temp_proto_goTypes = []interface{}{
	(Data)(0),             // 0: proto.Data
	(*Example)(nil),       // 1: proto.Example
	(*Foo)(nil),           // 2: proto.Foo
	(*Funk)(nil),          // 3: proto.Funk
	(*SampleMessage)(nil), // 4: proto.SampleMessage
	(*Example_Bar)(nil),   // 5: proto.Example.Bar
	nil,                   // 6: proto.Example.FooMapEntry
	(*Example_Far)(nil),   // 7: proto.Example.Far
	(*anypb.Any)(nil),     // 8: google.protobuf.Any
}
var file_temp_temp_proto_depIdxs = []int32{
	2,  // 0: proto.Example.foo:type_name -> proto.Foo
	5,  // 1: proto.Example.bar:type_name -> proto.Example.Bar
	8,  // 2: proto.Example.any:type_name -> google.protobuf.Any
	0,  // 3: proto.Example.data:type_name -> proto.Data
	6,  // 4: proto.Example.foo_map:type_name -> proto.Example.FooMapEntry
	4,  // 5: proto.Example.sample:type_name -> proto.SampleMessage
	7,  // 6: proto.Example.far:type_name -> proto.Example.Far
	2,  // 7: proto.SampleMessage.foo:type_name -> proto.Foo
	3,  // 8: proto.SampleMessage.funk:type_name -> proto.Funk
	2,  // 9: proto.Example.FooMapEntry.value:type_name -> proto.Foo
	1,  // 10: proto.ExampleAPI.ExampleRpc:input_type -> proto.Example
	1,  // 11: proto.ExampleAPI.ExampleAnyRpc:input_type -> proto.Example
	1,  // 12: proto.ExampleAPI.ExampleClientStream:input_type -> proto.Example
	1,  // 13: proto.ExampleAPI.ExampleServerStream:input_type -> proto.Example
	1,  // 14: proto.ExampleAPI.ExampleBidiStream:input_type -> proto.Example
	1,  // 15: proto.ExampleAPI.ExampleRpc:output_type -> proto.Example
	8,  // 16: proto.ExampleAPI.ExampleAnyRpc:output_type -> google.protobuf.Any
	1,  // 17: proto.ExampleAPI.ExampleClientStream:output_type -> proto.Example
	1,  // 18: proto.ExampleAPI.ExampleServerStream:output_type -> proto.Example
	1,  // 19: proto.ExampleAPI.ExampleBidiStream:output_type -> proto.Example
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_temp_temp_proto_init() }
func file_temp_temp_proto_init() {
	if File_temp_temp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temp_temp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_temp_temp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
		file_temp_temp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Funk); i {
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
		file_temp_temp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleMessage); i {
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
		file_temp_temp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example_Bar); i {
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
		file_temp_temp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example_Far); i {
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
	file_temp_temp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Example_Abc)(nil),
		(*Example_Far_)(nil),
	}
	file_temp_temp_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SampleMessage_Name)(nil),
		(*SampleMessage_Foo)(nil),
		(*SampleMessage_Funk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temp_temp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_temp_temp_proto_goTypes,
		DependencyIndexes: file_temp_temp_proto_depIdxs,
		EnumInfos:         file_temp_temp_proto_enumTypes,
		MessageInfos:      file_temp_temp_proto_msgTypes,
	}.Build()
	File_temp_temp_proto = out.File
	file_temp_temp_proto_rawDesc = nil
	file_temp_temp_proto_goTypes = nil
	file_temp_temp_proto_depIdxs = nil
}
