// Code generated by protoc-gen-go. DO NOT EDIT.

package protocol

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Properties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Properties) Reset() {
	*x = Properties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Properties) ProtoMessage() {}

func (x *Properties) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Properties.ProtoReflect.Descriptor instead.
func (*Properties) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{0}
}

func (x *Properties) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_ObjectValue
	//	*Value_ListValue
	//	*Value_DateValue
	//	*Value_UuidValue
	//	*Value_IntValue
	//	*Value_GeoValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{1}
}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Value) GetNumberValue() float64 {
	if x, ok := x.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetBoolValue() bool {
	if x, ok := x.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Value) GetObjectValue() *Properties {
	if x, ok := x.GetKind().(*Value_ObjectValue); ok {
		return x.ObjectValue
	}
	return nil
}

func (x *Value) GetListValue() *ListValue {
	if x, ok := x.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (x *Value) GetDateValue() string {
	if x, ok := x.GetKind().(*Value_DateValue); ok {
		return x.DateValue
	}
	return ""
}

func (x *Value) GetUuidValue() string {
	if x, ok := x.GetKind().(*Value_UuidValue); ok {
		return x.UuidValue
	}
	return ""
}

func (x *Value) GetIntValue() int64 {
	if x, ok := x.GetKind().(*Value_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *Value) GetGeoValue() *GeoCoordinate {
	if x, ok := x.GetKind().(*Value_GeoValue); ok {
		return x.GeoValue
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,1,opt,name=number_value,json=numberValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_ObjectValue struct {
	ObjectValue *Properties `protobuf:"bytes,4,opt,name=object_value,json=objectValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,5,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_DateValue struct {
	DateValue string `protobuf:"bytes,6,opt,name=date_value,json=dateValue,proto3,oneof"`
}

type Value_UuidValue struct {
	UuidValue string `protobuf:"bytes,7,opt,name=uuid_value,json=uuidValue,proto3,oneof"`
}

type Value_IntValue struct {
	IntValue int64 `protobuf:"varint,8,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Value_GeoValue struct {
	GeoValue *GeoCoordinate `protobuf:"bytes,9,opt,name=geo_value,json=geoValue,proto3,oneof"`
}

func (*Value_NumberValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_ObjectValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_DateValue) isValue_Kind() {}

func (*Value_UuidValue) isValue_Kind() {}

func (*Value_IntValue) isValue_Kind() {}

func (*Value_GeoValue) isValue_Kind() {}

type ListValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{2}
}

func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type GeoCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float32 `protobuf:"fixed32,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *GeoCoordinate) Reset() {
	*x = GeoCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_properties_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCoordinate) ProtoMessage() {}

func (x *GeoCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_v1_properties_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoCoordinate.ProtoReflect.Descriptor instead.
func (*GeoCoordinate) Descriptor() ([]byte, []int) {
	return file_v1_properties_proto_rawDescGZIP(), []int{3}
}

func (x *GeoCoordinate) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *GeoCoordinate) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

var File_v1_properties_proto protoreflect.FileDescriptor

var file_v1_properties_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x4d,
	0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8d, 0x03,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52,
	0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x09,
	0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x75, 0x75,
	0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x75, 0x75, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x67, 0x65,
	0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x08, 0x67, 0x65, 0x6f,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x37, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x65, 0x61,
	0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x42, 0x74, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_properties_proto_rawDescOnce sync.Once
	file_v1_properties_proto_rawDescData = file_v1_properties_proto_rawDesc
)

func file_v1_properties_proto_rawDescGZIP() []byte {
	file_v1_properties_proto_rawDescOnce.Do(func() {
		file_v1_properties_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_properties_proto_rawDescData)
	})
	return file_v1_properties_proto_rawDescData
}

var file_v1_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_properties_proto_goTypes = []interface{}{
	(*Properties)(nil),    // 0: weaviate.v1.Properties
	(*Value)(nil),         // 1: weaviate.v1.Value
	(*ListValue)(nil),     // 2: weaviate.v1.ListValue
	(*GeoCoordinate)(nil), // 3: weaviate.v1.GeoCoordinate
	nil,                   // 4: weaviate.v1.Properties.FieldsEntry
}
var file_v1_properties_proto_depIdxs = []int32{
	4, // 0: weaviate.v1.Properties.fields:type_name -> weaviate.v1.Properties.FieldsEntry
	0, // 1: weaviate.v1.Value.object_value:type_name -> weaviate.v1.Properties
	2, // 2: weaviate.v1.Value.list_value:type_name -> weaviate.v1.ListValue
	3, // 3: weaviate.v1.Value.geo_value:type_name -> weaviate.v1.GeoCoordinate
	1, // 4: weaviate.v1.ListValue.values:type_name -> weaviate.v1.Value
	1, // 5: weaviate.v1.Properties.FieldsEntry.value:type_name -> weaviate.v1.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_properties_proto_init() }
func file_v1_properties_proto_init() {
	if File_v1_properties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_properties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Properties); i {
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
		file_v1_properties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_v1_properties_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListValue); i {
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
		file_v1_properties_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCoordinate); i {
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
	file_v1_properties_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_ObjectValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_DateValue)(nil),
		(*Value_UuidValue)(nil),
		(*Value_IntValue)(nil),
		(*Value_GeoValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_properties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_properties_proto_goTypes,
		DependencyIndexes: file_v1_properties_proto_depIdxs,
		MessageInfos:      file_v1_properties_proto_msgTypes,
	}.Build()
	File_v1_properties_proto = out.File
	file_v1_properties_proto_rawDesc = nil
	file_v1_properties_proto_goTypes = nil
	file_v1_properties_proto_depIdxs = nil
}
