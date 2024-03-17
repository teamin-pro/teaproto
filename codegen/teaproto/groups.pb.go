// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.2
// source: groups.proto

package teaproto

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

type CreateGroupResponse_Error int32

const (
	CreateGroupResponse_UNKNOWN_ERROR                    CreateGroupResponse_Error = 0
	CreateGroupResponse_EMPTY_TITLE_ERROR                CreateGroupResponse_Error = 1
	CreateGroupResponse_TITLE_TOO_LONG_ERROR             CreateGroupResponse_Error = 2
	CreateGroupResponse_GROUP_CREATION_NOT_ALLOWED_ERROR CreateGroupResponse_Error = 3
)

// Enum value maps for CreateGroupResponse_Error.
var (
	CreateGroupResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "EMPTY_TITLE_ERROR",
		2: "TITLE_TOO_LONG_ERROR",
		3: "GROUP_CREATION_NOT_ALLOWED_ERROR",
	}
	CreateGroupResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":                    0,
		"EMPTY_TITLE_ERROR":                1,
		"TITLE_TOO_LONG_ERROR":             2,
		"GROUP_CREATION_NOT_ALLOWED_ERROR": 3,
	}
)

func (x CreateGroupResponse_Error) Enum() *CreateGroupResponse_Error {
	p := new(CreateGroupResponse_Error)
	*p = x
	return p
}

func (x CreateGroupResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateGroupResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_proto_enumTypes[0].Descriptor()
}

func (CreateGroupResponse_Error) Type() protoreflect.EnumType {
	return &file_groups_proto_enumTypes[0]
}

func (x CreateGroupResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateGroupResponse_Error.Descriptor instead.
func (CreateGroupResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{1, 0}
}

type UpdateGroupResponse_Error int32

const (
	UpdateGroupResponse_UNKNOWN_ERROR               UpdateGroupResponse_Error = 0
	UpdateGroupResponse_INVALID_CHAT_ERROR          UpdateGroupResponse_Error = 1
	UpdateGroupResponse_NOT_ALLOWED_TO_UPDATE_ERROR UpdateGroupResponse_Error = 2
	UpdateGroupResponse_EMPTY_TITLE_ERROR           UpdateGroupResponse_Error = 3
	UpdateGroupResponse_TITLE_TOO_LONG_ERROR        UpdateGroupResponse_Error = 4
)

// Enum value maps for UpdateGroupResponse_Error.
var (
	UpdateGroupResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "NOT_ALLOWED_TO_UPDATE_ERROR",
		3: "EMPTY_TITLE_ERROR",
		4: "TITLE_TOO_LONG_ERROR",
	}
	UpdateGroupResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":               0,
		"INVALID_CHAT_ERROR":          1,
		"NOT_ALLOWED_TO_UPDATE_ERROR": 2,
		"EMPTY_TITLE_ERROR":           3,
		"TITLE_TOO_LONG_ERROR":        4,
	}
)

func (x UpdateGroupResponse_Error) Enum() *UpdateGroupResponse_Error {
	p := new(UpdateGroupResponse_Error)
	*p = x
	return p
}

func (x UpdateGroupResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateGroupResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_proto_enumTypes[1].Descriptor()
}

func (UpdateGroupResponse_Error) Type() protoreflect.EnumType {
	return &file_groups_proto_enumTypes[1]
}

func (x UpdateGroupResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateGroupResponse_Error.Descriptor instead.
func (UpdateGroupResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{3, 0}
}

type GroupDetailsResponse_Error int32

const (
	GroupDetailsResponse_UNKNOWN_ERROR       GroupDetailsResponse_Error = 0
	GroupDetailsResponse_INVALID_CHAT_ERROR  GroupDetailsResponse_Error = 1
	GroupDetailsResponse_ACCESS_DENIED_ERROR GroupDetailsResponse_Error = 2
)

// Enum value maps for GroupDetailsResponse_Error.
var (
	GroupDetailsResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "ACCESS_DENIED_ERROR",
	}
	GroupDetailsResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":       0,
		"INVALID_CHAT_ERROR":  1,
		"ACCESS_DENIED_ERROR": 2,
	}
)

func (x GroupDetailsResponse_Error) Enum() *GroupDetailsResponse_Error {
	p := new(GroupDetailsResponse_Error)
	*p = x
	return p
}

func (x GroupDetailsResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupDetailsResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_proto_enumTypes[2].Descriptor()
}

func (GroupDetailsResponse_Error) Type() protoreflect.EnumType {
	return &file_groups_proto_enumTypes[2]
}

func (x GroupDetailsResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupDetailsResponse_Error.Descriptor instead.
func (GroupDetailsResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{5, 0}
}

type DeleteGroupResponse_Error int32

const (
	DeleteGroupResponse_UNKNOWN_ERROR       DeleteGroupResponse_Error = 0
	DeleteGroupResponse_INVALID_CHAT_ERROR  DeleteGroupResponse_Error = 1
	DeleteGroupResponse_ACCESS_DENIED_ERROR DeleteGroupResponse_Error = 2
)

// Enum value maps for DeleteGroupResponse_Error.
var (
	DeleteGroupResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "ACCESS_DENIED_ERROR",
	}
	DeleteGroupResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":       0,
		"INVALID_CHAT_ERROR":  1,
		"ACCESS_DENIED_ERROR": 2,
	}
)

func (x DeleteGroupResponse_Error) Enum() *DeleteGroupResponse_Error {
	p := new(DeleteGroupResponse_Error)
	*p = x
	return p
}

func (x DeleteGroupResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteGroupResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_proto_enumTypes[3].Descriptor()
}

func (DeleteGroupResponse_Error) Type() protoreflect.EnumType {
	return &file_groups_proto_enumTypes[3]
}

func (x DeleteGroupResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteGroupResponse_Error.Descriptor instead.
func (DeleteGroupResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{7, 0}
}

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateGroupRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []CreateGroupResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.CreateGroupResponse_Error" json:"errors,omitempty"`
	Chat   *Chat                       `protobuf:"bytes,3,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupResponse) GetErrors() []CreateGroupResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *CreateGroupResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateGroupRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *UpdateGroupRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []UpdateGroupResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.UpdateGroupResponse_Error" json:"errors,omitempty"`
	Chat   *Chat                       `protobuf:"bytes,3,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *UpdateGroupResponse) Reset() {
	*x = UpdateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupResponse) ProtoMessage() {}

func (x *UpdateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGroupResponse) GetErrors() []UpdateGroupResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *UpdateGroupResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type GroupDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *GroupDetailsRequest) Reset() {
	*x = GroupDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDetailsRequest) ProtoMessage() {}

func (x *GroupDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDetailsRequest.ProtoReflect.Descriptor instead.
func (*GroupDetailsRequest) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{4}
}

func (x *GroupDetailsRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type GroupDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors     []GroupDetailsResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.GroupDetailsResponse_Error" json:"errors,omitempty"`
	Chat       *Chat                        `protobuf:"bytes,3,opt,name=chat,proto3" json:"chat,omitempty"`
	GroupRoles []*GroupRole                 `protobuf:"bytes,2,rep,name=group_roles,json=groupRoles,proto3" json:"group_roles,omitempty"`
}

func (x *GroupDetailsResponse) Reset() {
	*x = GroupDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDetailsResponse) ProtoMessage() {}

func (x *GroupDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDetailsResponse.ProtoReflect.Descriptor instead.
func (*GroupDetailsResponse) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{5}
}

func (x *GroupDetailsResponse) GetErrors() []GroupDetailsResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *GroupDetailsResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *GroupDetailsResponse) GetGroupRoles() []*GroupRole {
	if x != nil {
		return x.GroupRoles
	}
	return nil
}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteGroupRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type DeleteGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []DeleteGroupResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.DeleteGroupResponse_Error" json:"errors,omitempty"`
}

func (x *DeleteGroupResponse) Reset() {
	*x = DeleteGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResponse) ProtoMessage() {}

func (x *DeleteGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteGroupResponse) GetErrors() []DeleteGroupResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_groups_proto protoreflect.FileDescriptor

var file_groups_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x74, 0x65, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x74, 0x65, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xe9,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x71, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x54, 0x49, 0x54,
	0x4c, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49,
	0x54, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x43, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0xfd, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x1f, 0x0a,
	0x1b, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x22,
	0x2e, 0x0a, 0x13, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22,
	0xfb, 0x01, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x65, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x22, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44,
	0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x2d, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x9f, 0x01, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x22, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groups_proto_rawDescOnce sync.Once
	file_groups_proto_rawDescData = file_groups_proto_rawDesc
)

func file_groups_proto_rawDescGZIP() []byte {
	file_groups_proto_rawDescOnce.Do(func() {
		file_groups_proto_rawDescData = protoimpl.X.CompressGZIP(file_groups_proto_rawDescData)
	})
	return file_groups_proto_rawDescData
}

var file_groups_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_groups_proto_goTypes = []interface{}{
	(CreateGroupResponse_Error)(0),  // 0: teaproto.CreateGroupResponse.Error
	(UpdateGroupResponse_Error)(0),  // 1: teaproto.UpdateGroupResponse.Error
	(GroupDetailsResponse_Error)(0), // 2: teaproto.GroupDetailsResponse.Error
	(DeleteGroupResponse_Error)(0),  // 3: teaproto.DeleteGroupResponse.Error
	(*CreateGroupRequest)(nil),      // 4: teaproto.CreateGroupRequest
	(*CreateGroupResponse)(nil),     // 5: teaproto.CreateGroupResponse
	(*UpdateGroupRequest)(nil),      // 6: teaproto.UpdateGroupRequest
	(*UpdateGroupResponse)(nil),     // 7: teaproto.UpdateGroupResponse
	(*GroupDetailsRequest)(nil),     // 8: teaproto.GroupDetailsRequest
	(*GroupDetailsResponse)(nil),    // 9: teaproto.GroupDetailsResponse
	(*DeleteGroupRequest)(nil),      // 10: teaproto.DeleteGroupRequest
	(*DeleteGroupResponse)(nil),     // 11: teaproto.DeleteGroupResponse
	(*Chat)(nil),                    // 12: teatypes.Chat
	(*GroupRole)(nil),               // 13: teatypes.GroupRole
}
var file_groups_proto_depIdxs = []int32{
	0,  // 0: teaproto.CreateGroupResponse.errors:type_name -> teaproto.CreateGroupResponse.Error
	12, // 1: teaproto.CreateGroupResponse.chat:type_name -> teatypes.Chat
	1,  // 2: teaproto.UpdateGroupResponse.errors:type_name -> teaproto.UpdateGroupResponse.Error
	12, // 3: teaproto.UpdateGroupResponse.chat:type_name -> teatypes.Chat
	2,  // 4: teaproto.GroupDetailsResponse.errors:type_name -> teaproto.GroupDetailsResponse.Error
	12, // 5: teaproto.GroupDetailsResponse.chat:type_name -> teatypes.Chat
	13, // 6: teaproto.GroupDetailsResponse.group_roles:type_name -> teatypes.GroupRole
	3,  // 7: teaproto.DeleteGroupResponse.errors:type_name -> teaproto.DeleteGroupResponse.Error
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_groups_proto_init() }
func file_groups_proto_init() {
	if File_groups_proto != nil {
		return
	}
	file_teatypes_tchats_proto_init()
	file_teatypes_tgroups_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_groups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_groups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_groups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_groups_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupResponse); i {
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
		file_groups_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDetailsRequest); i {
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
		file_groups_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDetailsResponse); i {
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
		file_groups_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_groups_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupResponse); i {
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
			RawDescriptor: file_groups_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_groups_proto_goTypes,
		DependencyIndexes: file_groups_proto_depIdxs,
		EnumInfos:         file_groups_proto_enumTypes,
		MessageInfos:      file_groups_proto_msgTypes,
	}.Build()
	File_groups_proto = out.File
	file_groups_proto_rawDesc = nil
	file_groups_proto_goTypes = nil
	file_groups_proto_depIdxs = nil
}
