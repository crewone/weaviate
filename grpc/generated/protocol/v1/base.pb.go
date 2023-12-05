//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by protoc-gen-go. DO NOT EDIT.

package protocol

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConsistencyLevel int32

const (
	ConsistencyLevel_CONSISTENCY_LEVEL_UNSPECIFIED ConsistencyLevel = 0
	ConsistencyLevel_CONSISTENCY_LEVEL_ONE         ConsistencyLevel = 1
	ConsistencyLevel_CONSISTENCY_LEVEL_QUORUM      ConsistencyLevel = 2
	ConsistencyLevel_CONSISTENCY_LEVEL_ALL         ConsistencyLevel = 3
)

// Enum value maps for ConsistencyLevel.
var (
	ConsistencyLevel_name = map[int32]string{
		0: "CONSISTENCY_LEVEL_UNSPECIFIED",
		1: "CONSISTENCY_LEVEL_ONE",
		2: "CONSISTENCY_LEVEL_QUORUM",
		3: "CONSISTENCY_LEVEL_ALL",
	}
	ConsistencyLevel_value = map[string]int32{
		"CONSISTENCY_LEVEL_UNSPECIFIED": 0,
		"CONSISTENCY_LEVEL_ONE":         1,
		"CONSISTENCY_LEVEL_QUORUM":      2,
		"CONSISTENCY_LEVEL_ALL":         3,
	}
)

func (x ConsistencyLevel) Enum() *ConsistencyLevel {
	p := new(ConsistencyLevel)
	*p = x
	return p
}

func (x ConsistencyLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsistencyLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_base_proto_enumTypes[0].Descriptor()
}

func (ConsistencyLevel) Type() protoreflect.EnumType {
	return &file_v1_base_proto_enumTypes[0]
}

func (x ConsistencyLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsistencyLevel.Descriptor instead.
func (ConsistencyLevel) EnumDescriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{0}
}

type NumberArrayProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in v1/base.proto.
	Values      []float64 `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"` // will be removed in the future, use vector_bytes
	PropName    string    `protobuf:"bytes,2,opt,name=prop_name,json=propName,proto3" json:"prop_name,omitempty"`
	ValuesBytes []byte    `protobuf:"bytes,3,opt,name=values_bytes,json=valuesBytes,proto3" json:"values_bytes,omitempty"`
}

func (x *NumberArrayProperties) Reset() {
	*x = NumberArrayProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberArrayProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberArrayProperties) ProtoMessage() {}

func (x *NumberArrayProperties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberArrayProperties.ProtoReflect.Descriptor instead.
func (*NumberArrayProperties) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in v1/base.proto.
func (x *NumberArrayProperties) GetValues() []float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *NumberArrayProperties) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

func (x *NumberArrayProperties) GetValuesBytes() []byte {
	if x != nil {
		return x.ValuesBytes
	}
	return nil
}

type IntArrayProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values   []int64 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	PropName string  `protobuf:"bytes,2,opt,name=prop_name,json=propName,proto3" json:"prop_name,omitempty"`
}

func (x *IntArrayProperties) Reset() {
	*x = IntArrayProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntArrayProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntArrayProperties) ProtoMessage() {}

func (x *IntArrayProperties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntArrayProperties.ProtoReflect.Descriptor instead.
func (*IntArrayProperties) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{1}
}

func (x *IntArrayProperties) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *IntArrayProperties) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

type TextArrayProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values   []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	PropName string   `protobuf:"bytes,2,opt,name=prop_name,json=propName,proto3" json:"prop_name,omitempty"`
}

func (x *TextArrayProperties) Reset() {
	*x = TextArrayProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextArrayProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextArrayProperties) ProtoMessage() {}

func (x *TextArrayProperties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextArrayProperties.ProtoReflect.Descriptor instead.
func (*TextArrayProperties) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{2}
}

func (x *TextArrayProperties) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *TextArrayProperties) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

type BooleanArrayProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values   []bool `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	PropName string `protobuf:"bytes,2,opt,name=prop_name,json=propName,proto3" json:"prop_name,omitempty"`
}

func (x *BooleanArrayProperties) Reset() {
	*x = BooleanArrayProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanArrayProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanArrayProperties) ProtoMessage() {}

func (x *BooleanArrayProperties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanArrayProperties.ProtoReflect.Descriptor instead.
func (*BooleanArrayProperties) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{3}
}

func (x *BooleanArrayProperties) GetValues() []bool {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *BooleanArrayProperties) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

type ObjectPropertiesValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NonRefProperties       *structpb.Struct          `protobuf:"bytes,1,opt,name=non_ref_properties,json=nonRefProperties,proto3" json:"non_ref_properties,omitempty"`
	NumberArrayProperties  []*NumberArrayProperties  `protobuf:"bytes,2,rep,name=number_array_properties,json=numberArrayProperties,proto3" json:"number_array_properties,omitempty"`
	IntArrayProperties     []*IntArrayProperties     `protobuf:"bytes,3,rep,name=int_array_properties,json=intArrayProperties,proto3" json:"int_array_properties,omitempty"`
	TextArrayProperties    []*TextArrayProperties    `protobuf:"bytes,4,rep,name=text_array_properties,json=textArrayProperties,proto3" json:"text_array_properties,omitempty"`
	BooleanArrayProperties []*BooleanArrayProperties `protobuf:"bytes,5,rep,name=boolean_array_properties,json=booleanArrayProperties,proto3" json:"boolean_array_properties,omitempty"`
	ObjectProperties       []*ObjectProperties       `protobuf:"bytes,6,rep,name=object_properties,json=objectProperties,proto3" json:"object_properties,omitempty"`
	ObjectArrayProperties  []*ObjectArrayProperties  `protobuf:"bytes,7,rep,name=object_array_properties,json=objectArrayProperties,proto3" json:"object_array_properties,omitempty"`
}

func (x *ObjectPropertiesValue) Reset() {
	*x = ObjectPropertiesValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectPropertiesValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectPropertiesValue) ProtoMessage() {}

func (x *ObjectPropertiesValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectPropertiesValue.ProtoReflect.Descriptor instead.
func (*ObjectPropertiesValue) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectPropertiesValue) GetNonRefProperties() *structpb.Struct {
	if x != nil {
		return x.NonRefProperties
	}
	return nil
}

func (x *ObjectPropertiesValue) GetNumberArrayProperties() []*NumberArrayProperties {
	if x != nil {
		return x.NumberArrayProperties
	}
	return nil
}

func (x *ObjectPropertiesValue) GetIntArrayProperties() []*IntArrayProperties {
	if x != nil {
		return x.IntArrayProperties
	}
	return nil
}

func (x *ObjectPropertiesValue) GetTextArrayProperties() []*TextArrayProperties {
	if x != nil {
		return x.TextArrayProperties
	}
	return nil
}

func (x *ObjectPropertiesValue) GetBooleanArrayProperties() []*BooleanArrayProperties {
	if x != nil {
		return x.BooleanArrayProperties
	}
	return nil
}

func (x *ObjectPropertiesValue) GetObjectProperties() []*ObjectProperties {
	if x != nil {
		return x.ObjectProperties
	}
	return nil
}

func (x *ObjectPropertiesValue) GetObjectArrayProperties() []*ObjectArrayProperties {
	if x != nil {
		return x.ObjectArrayProperties
	}
	return nil
}

type ObjectArrayProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values   []*ObjectPropertiesValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	PropName string                   `protobuf:"bytes,2,opt,name=prop_name,json=propName,proto3" json:"prop_name,omitempty"`
}

func (x *ObjectArrayProperties) Reset() {
	*x = ObjectArrayProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectArrayProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectArrayProperties) ProtoMessage() {}

func (x *ObjectArrayProperties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectArrayProperties.ProtoReflect.Descriptor instead.
func (*ObjectArrayProperties) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{5}
}

func (x *ObjectArrayProperties) GetValues() []*ObjectPropertiesValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ObjectArrayProperties) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

type ObjectProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    *ObjectPropertiesValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	PropName string                 `protobuf:"bytes,2,opt,name=prop_name,json=propName,proto3" json:"prop_name,omitempty"`
}

func (x *ObjectProperties) Reset() {
	*x = ObjectProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectProperties) ProtoMessage() {}

func (x *ObjectProperties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectProperties.ProtoReflect.Descriptor instead.
func (*ObjectProperties) Descriptor() ([]byte, []int) {
	return file_v1_base_proto_rawDescGZIP(), []int{6}
}

func (x *ObjectProperties) GetValue() *ObjectPropertiesValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ObjectProperties) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

var File_v1_base_proto protoreflect.FileDescriptor

var file_v1_base_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x15, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x01, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22,
	0x49, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x13, 0x54, 0x65,
	0x78, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x16, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xea, 0x04, 0x0a, 0x15, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x45, 0x0a, 0x12, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x10, 0x6e, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x17, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x15, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x51, 0x0a, 0x14, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x15, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x13, 0x74, 0x65, 0x78, 0x74, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x18, 0x62,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x16, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x11, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x17, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x15, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x70, 0x0a, 0x15, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x65,
	0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x2a,
	0x89, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45,
	0x4e, 0x43, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x53, 0x49,
	0x53, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x43,
	0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x51, 0x55, 0x4f, 0x52, 0x55, 0x4d, 0x10, 0x02,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x42, 0x6e, 0x0a, 0x23, 0x69,
	0x6f, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x42, 0x11, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x61, 0x73, 0x65, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_base_proto_rawDescOnce sync.Once
	file_v1_base_proto_rawDescData = file_v1_base_proto_rawDesc
)

func file_v1_base_proto_rawDescGZIP() []byte {
	file_v1_base_proto_rawDescOnce.Do(func() {
		file_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_base_proto_rawDescData)
	})
	return file_v1_base_proto_rawDescData
}

var file_v1_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_base_proto_goTypes = []interface{}{
	(ConsistencyLevel)(0),          // 0: weaviate.v1.ConsistencyLevel
	(*NumberArrayProperties)(nil),  // 1: weaviate.v1.NumberArrayProperties
	(*IntArrayProperties)(nil),     // 2: weaviate.v1.IntArrayProperties
	(*TextArrayProperties)(nil),    // 3: weaviate.v1.TextArrayProperties
	(*BooleanArrayProperties)(nil), // 4: weaviate.v1.BooleanArrayProperties
	(*ObjectPropertiesValue)(nil),  // 5: weaviate.v1.ObjectPropertiesValue
	(*ObjectArrayProperties)(nil),  // 6: weaviate.v1.ObjectArrayProperties
	(*ObjectProperties)(nil),       // 7: weaviate.v1.ObjectProperties
	(*structpb.Struct)(nil),        // 8: google.protobuf.Struct
}
var file_v1_base_proto_depIdxs = []int32{
	8, // 0: weaviate.v1.ObjectPropertiesValue.non_ref_properties:type_name -> google.protobuf.Struct
	1, // 1: weaviate.v1.ObjectPropertiesValue.number_array_properties:type_name -> weaviate.v1.NumberArrayProperties
	2, // 2: weaviate.v1.ObjectPropertiesValue.int_array_properties:type_name -> weaviate.v1.IntArrayProperties
	3, // 3: weaviate.v1.ObjectPropertiesValue.text_array_properties:type_name -> weaviate.v1.TextArrayProperties
	4, // 4: weaviate.v1.ObjectPropertiesValue.boolean_array_properties:type_name -> weaviate.v1.BooleanArrayProperties
	7, // 5: weaviate.v1.ObjectPropertiesValue.object_properties:type_name -> weaviate.v1.ObjectProperties
	6, // 6: weaviate.v1.ObjectPropertiesValue.object_array_properties:type_name -> weaviate.v1.ObjectArrayProperties
	5, // 7: weaviate.v1.ObjectArrayProperties.values:type_name -> weaviate.v1.ObjectPropertiesValue
	5, // 8: weaviate.v1.ObjectProperties.value:type_name -> weaviate.v1.ObjectPropertiesValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_v1_base_proto_init() }
func file_v1_base_proto_init() {
	if File_v1_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberArrayProperties); i {
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
		file_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntArrayProperties); i {
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
		file_v1_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextArrayProperties); i {
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
		file_v1_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanArrayProperties); i {
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
		file_v1_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectPropertiesValue); i {
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
		file_v1_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectArrayProperties); i {
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
		file_v1_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectProperties); i {
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
			RawDescriptor: file_v1_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_base_proto_goTypes,
		DependencyIndexes: file_v1_base_proto_depIdxs,
		EnumInfos:         file_v1_base_proto_enumTypes,
		MessageInfos:      file_v1_base_proto_msgTypes,
	}.Build()
	File_v1_base_proto = out.File
	file_v1_base_proto_rawDesc = nil
	file_v1_base_proto_goTypes = nil
	file_v1_base_proto_depIdxs = nil
}
