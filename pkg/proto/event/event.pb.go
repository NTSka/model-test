// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: event.proto

package event

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

type Format int32

const (
	Format_XML  Format = 0
	Format_JSON Format = 1
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "XML",
		1: "JSON",
	}
	Format_value = map[string]int32{
		"XML":  0,
		"JSON": 1,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[0].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[0]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

type Importance int32

const (
	Importance_Info   Importance = 0
	Importance_Low    Importance = 1
	Importance_High   Importance = 2
	Importance_Medium Importance = 3
)

// Enum value maps for Importance.
var (
	Importance_name = map[int32]string{
		0: "Info",
		1: "Low",
		2: "High",
		3: "Medium",
	}
	Importance_value = map[string]int32{
		"Info":   0,
		"Low":    1,
		"High":   2,
		"Medium": 3,
	}
)

func (x Importance) Enum() *Importance {
	p := new(Importance)
	*p = x
	return p
}

func (x Importance) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Importance) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[1].Descriptor()
}

func (Importance) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[1]
}

func (x Importance) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Importance.Descriptor instead.
func (Importance) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

type RawEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format Format `protobuf:"varint,1,opt,name=format,proto3,enum=Format" json:"format,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RawEvent) Reset() {
	*x = RawEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawEvent) ProtoMessage() {}

func (x *RawEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawEvent.ProtoReflect.Descriptor instead.
func (*RawEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *RawEvent) GetFormat() Format {
	if x != nil {
		return x.Format
	}
	return Format_XML
}

func (x *RawEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventSrc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *EventSrc) Reset() {
	*x = EventSrc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSrc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSrc) ProtoMessage() {}

func (x *EventSrc) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSrc.ProtoReflect.Descriptor instead.
func (*EventSrc) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventSrc) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *EventSrc) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assets     []string   `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	EventSrc   *EventSrc  `protobuf:"bytes,2,opt,name=event_src,json=eventSrc,proto3" json:"event_src,omitempty"`
	Action     string     `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Importance Importance `protobuf:"varint,4,opt,name=importance,proto3,enum=Importance" json:"importance,omitempty"`
	Object     string     `protobuf:"bytes,5,opt,name=Object,proto3" json:"Object,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetAssets() []string {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *Event) GetEventSrc() *EventSrc {
	if x != nil {
		return x.EventSrc
	}
	return nil
}

func (x *Event) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Event) GetImportance() Importance {
	if x != nil {
		return x.Importance
	}
	return Importance_Info
}

func (x *Event) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type EventStep1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event     *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Meta1     string `protobuf:"bytes,3,opt,name=meta1,proto3" json:"meta1,omitempty"`
}

func (x *EventStep1) Reset() {
	*x = EventStep1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStep1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStep1) ProtoMessage() {}

func (x *EventStep1) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStep1.ProtoReflect.Descriptor instead.
func (*EventStep1) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{3}
}

func (x *EventStep1) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *EventStep1) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *EventStep1) GetMeta1() string {
	if x != nil {
		return x.Meta1
	}
	return ""
}

type EventStep2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *EventStep1 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Meta2 string      `protobuf:"bytes,2,opt,name=meta2,proto3" json:"meta2,omitempty"`
	Meta3 string      `protobuf:"bytes,3,opt,name=meta3,proto3" json:"meta3,omitempty"`
}

func (x *EventStep2) Reset() {
	*x = EventStep2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStep2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStep2) ProtoMessage() {}

func (x *EventStep2) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStep2.ProtoReflect.Descriptor instead.
func (*EventStep2) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{4}
}

func (x *EventStep2) GetEvent() *EventStep1 {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *EventStep2) GetMeta2() string {
	if x != nil {
		return x.Meta2
	}
	return ""
}

func (x *EventStep2) GetMeta3() string {
	if x != nil {
		return x.Meta3
	}
	return ""
}

type EventStep3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *EventStep2 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Meta4 string      `protobuf:"bytes,2,opt,name=meta4,proto3" json:"meta4,omitempty"`
	Meta5 string      `protobuf:"bytes,3,opt,name=meta5,proto3" json:"meta5,omitempty"`
}

func (x *EventStep3) Reset() {
	*x = EventStep3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStep3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStep3) ProtoMessage() {}

func (x *EventStep3) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStep3.ProtoReflect.Descriptor instead.
func (*EventStep3) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5}
}

func (x *EventStep3) GetEvent() *EventStep2 {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *EventStep3) GetMeta4() string {
	if x != nil {
		return x.Meta4
	}
	return ""
}

func (x *EventStep3) GetMeta5() string {
	if x != nil {
		return x.Meta5
	}
	return ""
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x08, 0x52, 0x61, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x32,
	0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x72, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x72,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x72, 0x63, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x72, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5e, 0x0a, 0x0a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x31, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x31, 0x22, 0x5b, 0x0a, 0x0a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x32, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x65, 0x70, 0x31, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65,
	0x74, 0x61, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x32,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x65, 0x74, 0x61, 0x33, 0x22, 0x5b, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x65, 0x70, 0x33, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x32,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x34,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x34, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x65, 0x74, 0x61, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65,
	0x74, 0x61, 0x35, 0x2a, 0x1b, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x07, 0x0a,
	0x03, 0x58, 0x4d, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01,
	0x2a, 0x35, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_event_proto_goTypes = []interface{}{
	(Format)(0),        // 0: Format
	(Importance)(0),    // 1: Importance
	(*RawEvent)(nil),   // 2: RawEvent
	(*EventSrc)(nil),   // 3: EventSrc
	(*Event)(nil),      // 4: Event
	(*EventStep1)(nil), // 5: EventStep1
	(*EventStep2)(nil), // 6: EventStep2
	(*EventStep3)(nil), // 7: EventStep3
}
var file_event_proto_depIdxs = []int32{
	0, // 0: RawEvent.format:type_name -> Format
	3, // 1: Event.event_src:type_name -> EventSrc
	1, // 2: Event.importance:type_name -> Importance
	4, // 3: EventStep1.event:type_name -> Event
	5, // 4: EventStep2.event:type_name -> EventStep1
	6, // 5: EventStep3.event:type_name -> EventStep2
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawEvent); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSrc); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStep1); i {
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
		file_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStep2); i {
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
		file_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStep3); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		EnumInfos:         file_event_proto_enumTypes,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
