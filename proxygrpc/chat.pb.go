// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: chat.proto

package proxygrpc

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

type ChannelInfo_TypeChannel int32

const (
	ChannelInfo_CHAT    ChannelInfo_TypeChannel = 0
	ChannelInfo_CHANNEL ChannelInfo_TypeChannel = 1
)

// Enum value maps for ChannelInfo_TypeChannel.
var (
	ChannelInfo_TypeChannel_name = map[int32]string{
		0: "CHAT",
		1: "CHANNEL",
	}
	ChannelInfo_TypeChannel_value = map[string]int32{
		"CHAT":    0,
		"CHANNEL": 1,
	}
)

func (x ChannelInfo_TypeChannel) Enum() *ChannelInfo_TypeChannel {
	p := new(ChannelInfo_TypeChannel)
	*p = x
	return p
}

func (x ChannelInfo_TypeChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelInfo_TypeChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (ChannelInfo_TypeChannel) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x ChannelInfo_TypeChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelInfo_TypeChannel.Descriptor instead.
func (ChannelInfo_TypeChannel) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2, 0}
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *GetInfoRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type GetInfoResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelInfo []*ChannelInfo `protobuf:"bytes,1,rep,name=channelInfo,proto3" json:"channelInfo,omitempty"`
}

func (x *GetInfoResponce) Reset() {
	*x = GetInfoResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponce) ProtoMessage() {}

func (x *GetInfoResponce) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponce.ProtoReflect.Descriptor instead.
func (*GetInfoResponce) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoResponce) GetChannelInfo() []*ChannelInfo {
	if x != nil {
		return x.ChannelInfo
	}
	return nil
}

type ChannelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdChannel   int32                   `protobuf:"varint,1,opt,name=idChannel,proto3" json:"idChannel,omitempty"`
	NameChannel int32                   `protobuf:"varint,2,opt,name=nameChannel,proto3" json:"nameChannel,omitempty"`
	TypeChannel ChannelInfo_TypeChannel `protobuf:"varint,3,opt,name=typeChannel,proto3,enum=messanger.ChannelInfo_TypeChannel" json:"typeChannel,omitempty"`
}

func (x *ChannelInfo) Reset() {
	*x = ChannelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelInfo) ProtoMessage() {}

func (x *ChannelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelInfo.ProtoReflect.Descriptor instead.
func (*ChannelInfo) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelInfo) GetIdChannel() int32 {
	if x != nil {
		return x.IdChannel
	}
	return 0
}

func (x *ChannelInfo) GetNameChannel() int32 {
	if x != nil {
		return x.NameChannel
	}
	return 0
}

func (x *ChannelInfo) GetTypeChannel() ChannelInfo_TypeChannel {
	if x != nil {
		return x.TypeChannel
	}
	return ChannelInfo_CHAT
}

type StartChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ChannelInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Text *ChannelMessage `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *StartChannelRequest) Reset() {
	*x = StartChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartChannelRequest) ProtoMessage() {}

func (x *StartChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartChannelRequest.ProtoReflect.Descriptor instead.
func (*StartChannelRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *StartChannelRequest) GetInfo() *ChannelInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *StartChannelRequest) GetText() *ChannelMessage {
	if x != nil {
		return x.Text
	}
	return nil
}

type JoinChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdChannel int32 `protobuf:"varint,1,opt,name=idChannel,proto3" json:"idChannel,omitempty"`
}

func (x *JoinChannelRequest) Reset() {
	*x = JoinChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChannelRequest) ProtoMessage() {}

func (x *JoinChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChannelRequest.ProtoReflect.Descriptor instead.
func (*JoinChannelRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *JoinChannelRequest) GetIdChannel() int32 {
	if x != nil {
		return x.IdChannel
	}
	return 0
}

type StartChannelResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StartChannelResponce) Reset() {
	*x = StartChannelResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartChannelResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartChannelResponce) ProtoMessage() {}

func (x *StartChannelResponce) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartChannelResponce.ProtoReflect.Descriptor instead.
func (*StartChannelResponce) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *StartChannelResponce) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ChannelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChannelMessage) Reset() {
	*x = ChannelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessage) ProtoMessage() {}

func (x *ChannelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessage.ProtoReflect.Descriptor instead.
func (*ChannelMessage) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *ChannelMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ChannelInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Name string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Text string       `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *Chat) GetInfo() *ChannelInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Chat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chat) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x44, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x24, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x48, 0x41, 0x54, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x01, 0x22, 0x70, 0x0a, 0x13,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x2d, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x32,
	0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x22, 0x2e, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5a,
	0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xa5, 0x02, 0x0a, 0x09, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0b,
	0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1d, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x32,
	0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0f, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x2e, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x75, 0x62, 0x61, 0x6e,
	0x6f, 0x76, 0x61, 0x42, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chat_proto_goTypes = []interface{}{
	(ChannelInfo_TypeChannel)(0), // 0: messanger.ChannelInfo.TypeChannel
	(*GetInfoRequest)(nil),       // 1: messanger.GetInfoRequest
	(*GetInfoResponce)(nil),      // 2: messanger.GetInfoResponce
	(*ChannelInfo)(nil),          // 3: messanger.ChannelInfo
	(*StartChannelRequest)(nil),  // 4: messanger.StartChannelRequest
	(*JoinChannelRequest)(nil),   // 5: messanger.JoinChannelRequest
	(*StartChannelResponce)(nil), // 6: messanger.StartChannelResponce
	(*ChannelMessage)(nil),       // 7: messanger.ChannelMessage
	(*Chat)(nil),                 // 8: messanger.Chat
}
var file_chat_proto_depIdxs = []int32{
	3, // 0: messanger.GetInfoResponce.channelInfo:type_name -> messanger.ChannelInfo
	0, // 1: messanger.ChannelInfo.typeChannel:type_name -> messanger.ChannelInfo.TypeChannel
	3, // 2: messanger.StartChannelRequest.info:type_name -> messanger.ChannelInfo
	7, // 3: messanger.StartChannelRequest.text:type_name -> messanger.ChannelMessage
	3, // 4: messanger.Chat.info:type_name -> messanger.ChannelInfo
	1, // 5: messanger.Messanger.GetInfo:input_type -> messanger.GetInfoRequest
	5, // 6: messanger.Messanger.JoinChannel:input_type -> messanger.JoinChannelRequest
	4, // 7: messanger.Messanger.StartChannel:input_type -> messanger.StartChannelRequest
	8, // 8: messanger.Messanger.JoinChat:input_type -> messanger.Chat
	2, // 9: messanger.Messanger.GetInfo:output_type -> messanger.GetInfoResponce
	7, // 10: messanger.Messanger.JoinChannel:output_type -> messanger.ChannelMessage
	6, // 11: messanger.Messanger.StartChannel:output_type -> messanger.StartChannelResponce
	8, // 12: messanger.Messanger.JoinChat:output_type -> messanger.Chat
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponce); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelInfo); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartChannelRequest); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinChannelRequest); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartChannelResponce); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelMessage); i {
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
		file_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
