// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.19.0-devel
// 	protoc        v3.11.3
// source: github.com/afking/graphpb/testpb/test.proto

package testpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(19 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 19)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"` // The resource content
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_github_com_afking_graphpb_testpb_test_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetMessageRequestOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Mapped to URL path
}

func (x *GetMessageRequestOne) Reset() {
	*x = GetMessageRequestOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequestOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequestOne) ProtoMessage() {}

func (x *GetMessageRequestOne) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequestOne.ProtoReflect.Descriptor instead.
func (*GetMessageRequestOne) Descriptor() ([]byte, []int) {
	return file_github_com_afking_graphpb_testpb_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessageRequestOne) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMessageRequestTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string                           `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // Mapped to URL path
	Revision  int64                            `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`                   // Mapped to URL query parameter `revision`
	Sub       *GetMessageRequestTwo_SubMessage `protobuf:"bytes,3,opt,name=sub,proto3" json:"sub,omitempty"`                              // Mapped to URL query parameter `sub.subfield`
	UserId    string                           `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // Additional binding
}

func (x *GetMessageRequestTwo) Reset() {
	*x = GetMessageRequestTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequestTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequestTwo) ProtoMessage() {}

func (x *GetMessageRequestTwo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequestTwo.ProtoReflect.Descriptor instead.
func (*GetMessageRequestTwo) Descriptor() ([]byte, []int) {
	return file_github_com_afking_graphpb_testpb_test_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequestTwo) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *GetMessageRequestTwo) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *GetMessageRequestTwo) GetSub() *GetMessageRequestTwo_SubMessage {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *GetMessageRequestTwo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateMessageRequestOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string   `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // Mapped to the URL
	Message   *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                      // Mapped to the body
}

func (x *UpdateMessageRequestOne) Reset() {
	*x = UpdateMessageRequestOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRequestOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRequestOne) ProtoMessage() {}

func (x *UpdateMessageRequestOne) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRequestOne.ProtoReflect.Descriptor instead.
func (*UpdateMessageRequestOne) Descriptor() ([]byte, []int) {
	return file_github_com_afking_graphpb_testpb_test_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMessageRequestOne) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *UpdateMessageRequestOne) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type GetMessageRequestTwo_SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subfield string `protobuf:"bytes,1,opt,name=subfield,proto3" json:"subfield,omitempty"`
}

func (x *GetMessageRequestTwo_SubMessage) Reset() {
	*x = GetMessageRequestTwo_SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequestTwo_SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequestTwo_SubMessage) ProtoMessage() {}

func (x *GetMessageRequestTwo_SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_afking_graphpb_testpb_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequestTwo_SubMessage.ProtoReflect.Descriptor instead.
func (*GetMessageRequestTwo_SubMessage) Descriptor() ([]byte, []int) {
	return file_github_com_afking_graphpb_testpb_test_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetMessageRequestTwo_SubMessage) GetSubfield() string {
	if x != nil {
		return x.Subfield
	}
	return ""
}

var File_github_com_afking_graphpb_testpb_test_proto protoreflect.FileDescriptor

var file_github_com_afking_graphpb_testpb_test_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x66, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd7,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x77, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x28,
	0x0a, 0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x6b, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x91, 0x04, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x12, 0x72, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x6e, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x2a, 0x7d, 0x12, 0x9e, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x77, 0x6f, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x1a,
	0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48,
	0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x2b, 0x12, 0x29, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f,
	0x6e, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x24, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x17, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x32, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x66, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x3b, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_afking_graphpb_testpb_test_proto_rawDescOnce sync.Once
	file_github_com_afking_graphpb_testpb_test_proto_rawDescData = file_github_com_afking_graphpb_testpb_test_proto_rawDesc
)

func file_github_com_afking_graphpb_testpb_test_proto_rawDescGZIP() []byte {
	file_github_com_afking_graphpb_testpb_test_proto_rawDescOnce.Do(func() {
		file_github_com_afking_graphpb_testpb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_afking_graphpb_testpb_test_proto_rawDescData)
	})
	return file_github_com_afking_graphpb_testpb_test_proto_rawDescData
}

var file_github_com_afking_graphpb_testpb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_afking_graphpb_testpb_test_proto_goTypes = []interface{}{
	(*Message)(nil),                         // 0: graphpb.testpb.Message
	(*GetMessageRequestOne)(nil),            // 1: graphpb.testpb.GetMessageRequestOne
	(*GetMessageRequestTwo)(nil),            // 2: graphpb.testpb.GetMessageRequestTwo
	(*UpdateMessageRequestOne)(nil),         // 3: graphpb.testpb.UpdateMessageRequestOne
	(*GetMessageRequestTwo_SubMessage)(nil), // 4: graphpb.testpb.GetMessageRequestTwo.SubMessage
}
var file_github_com_afking_graphpb_testpb_test_proto_depIdxs = []int32{
	4, // 0: graphpb.testpb.GetMessageRequestTwo.sub:type_name -> graphpb.testpb.GetMessageRequestTwo.SubMessage
	0, // 1: graphpb.testpb.UpdateMessageRequestOne.message:type_name -> graphpb.testpb.Message
	1, // 2: graphpb.testpb.Messaging.GetMessageOne:input_type -> graphpb.testpb.GetMessageRequestOne
	2, // 3: graphpb.testpb.Messaging.GetMessageTwo:input_type -> graphpb.testpb.GetMessageRequestTwo
	3, // 4: graphpb.testpb.Messaging.UpdateMessage:input_type -> graphpb.testpb.UpdateMessageRequestOne
	0, // 5: graphpb.testpb.Messaging.UpdateMessageBody:input_type -> graphpb.testpb.Message
	0, // 6: graphpb.testpb.Messaging.GetMessageOne:output_type -> graphpb.testpb.Message
	0, // 7: graphpb.testpb.Messaging.GetMessageTwo:output_type -> graphpb.testpb.Message
	0, // 8: graphpb.testpb.Messaging.UpdateMessage:output_type -> graphpb.testpb.Message
	0, // 9: graphpb.testpb.Messaging.UpdateMessageBody:output_type -> graphpb.testpb.Message
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_afking_graphpb_testpb_test_proto_init() }
func file_github_com_afking_graphpb_testpb_test_proto_init() {
	if File_github_com_afking_graphpb_testpb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_afking_graphpb_testpb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_github_com_afking_graphpb_testpb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequestOne); i {
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
		file_github_com_afking_graphpb_testpb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequestTwo); i {
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
		file_github_com_afking_graphpb_testpb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRequestOne); i {
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
		file_github_com_afking_graphpb_testpb_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequestTwo_SubMessage); i {
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
			RawDescriptor: file_github_com_afking_graphpb_testpb_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_afking_graphpb_testpb_test_proto_goTypes,
		DependencyIndexes: file_github_com_afking_graphpb_testpb_test_proto_depIdxs,
		MessageInfos:      file_github_com_afking_graphpb_testpb_test_proto_msgTypes,
	}.Build()
	File_github_com_afking_graphpb_testpb_test_proto = out.File
	file_github_com_afking_graphpb_testpb_test_proto_rawDesc = nil
	file_github_com_afking_graphpb_testpb_test_proto_goTypes = nil
	file_github_com_afking_graphpb_testpb_test_proto_depIdxs = nil
}
