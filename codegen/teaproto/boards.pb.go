// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: boards.proto

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

type BoardDetailsResponse_Error int32

const (
	BoardDetailsResponse_UNKNOWN_ERROR       BoardDetailsResponse_Error = 0
	BoardDetailsResponse_INVALID_CHAT_ERROR  BoardDetailsResponse_Error = 1
	BoardDetailsResponse_ACCESS_DENIED_ERROR BoardDetailsResponse_Error = 2
)

// Enum value maps for BoardDetailsResponse_Error.
var (
	BoardDetailsResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "ACCESS_DENIED_ERROR",
	}
	BoardDetailsResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":       0,
		"INVALID_CHAT_ERROR":  1,
		"ACCESS_DENIED_ERROR": 2,
	}
)

func (x BoardDetailsResponse_Error) Enum() *BoardDetailsResponse_Error {
	p := new(BoardDetailsResponse_Error)
	*p = x
	return p
}

func (x BoardDetailsResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BoardDetailsResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_boards_proto_enumTypes[0].Descriptor()
}

func (BoardDetailsResponse_Error) Type() protoreflect.EnumType {
	return &file_boards_proto_enumTypes[0]
}

func (x BoardDetailsResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BoardDetailsResponse_Error.Descriptor instead.
func (BoardDetailsResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{1, 0}
}

type CreateBoardResponse_Error int32

const (
	CreateBoardResponse_UNKNOWN_ERROR        CreateBoardResponse_Error = 0
	CreateBoardResponse_EMPTY_TITLE_ERROR    CreateBoardResponse_Error = 1
	CreateBoardResponse_TITLE_TOO_LONG_ERROR CreateBoardResponse_Error = 2 // see StateResponse.max_board_title_length
)

// Enum value maps for CreateBoardResponse_Error.
var (
	CreateBoardResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "EMPTY_TITLE_ERROR",
		2: "TITLE_TOO_LONG_ERROR",
	}
	CreateBoardResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":        0,
		"EMPTY_TITLE_ERROR":    1,
		"TITLE_TOO_LONG_ERROR": 2,
	}
)

func (x CreateBoardResponse_Error) Enum() *CreateBoardResponse_Error {
	p := new(CreateBoardResponse_Error)
	*p = x
	return p
}

func (x CreateBoardResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateBoardResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_boards_proto_enumTypes[1].Descriptor()
}

func (CreateBoardResponse_Error) Type() protoreflect.EnumType {
	return &file_boards_proto_enumTypes[1]
}

func (x CreateBoardResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateBoardResponse_Error.Descriptor instead.
func (CreateBoardResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{3, 0}
}

type UpdateBoardResponse_Error int32

const (
	UpdateBoardResponse_UNKNOWN_ERROR                  UpdateBoardResponse_Error = 0
	UpdateBoardResponse_INVALID_CHAT_ERROR             UpdateBoardResponse_Error = 1
	UpdateBoardResponse_EMPTY_TITLE_ERROR              UpdateBoardResponse_Error = 2
	UpdateBoardResponse_TITLE_TOO_LONG_ERROR           UpdateBoardResponse_Error = 3 // see StateResponse.max_board_title_length
	UpdateBoardResponse_UPDATE_BOARD_NOT_ALLOWED_ERROR UpdateBoardResponse_Error = 4
)

// Enum value maps for UpdateBoardResponse_Error.
var (
	UpdateBoardResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "EMPTY_TITLE_ERROR",
		3: "TITLE_TOO_LONG_ERROR",
		4: "UPDATE_BOARD_NOT_ALLOWED_ERROR",
	}
	UpdateBoardResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":                  0,
		"INVALID_CHAT_ERROR":             1,
		"EMPTY_TITLE_ERROR":              2,
		"TITLE_TOO_LONG_ERROR":           3,
		"UPDATE_BOARD_NOT_ALLOWED_ERROR": 4,
	}
)

func (x UpdateBoardResponse_Error) Enum() *UpdateBoardResponse_Error {
	p := new(UpdateBoardResponse_Error)
	*p = x
	return p
}

func (x UpdateBoardResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateBoardResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_boards_proto_enumTypes[2].Descriptor()
}

func (UpdateBoardResponse_Error) Type() protoreflect.EnumType {
	return &file_boards_proto_enumTypes[2]
}

func (x UpdateBoardResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateBoardResponse_Error.Descriptor instead.
func (UpdateBoardResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{5, 0}
}

type BoardDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *BoardDetailsRequest) Reset() {
	*x = BoardDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardDetailsRequest) ProtoMessage() {}

func (x *BoardDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardDetailsRequest.ProtoReflect.Descriptor instead.
func (*BoardDetailsRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{0}
}

func (x *BoardDetailsRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type BoardDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors       []BoardDetailsResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.BoardDetailsResponse_Error" json:"errors,omitempty"`
	Chat         *Chat                        `protobuf:"bytes,4,opt,name=chat,proto3" json:"chat,omitempty"`
	BoardRoles   []*BoardRole                 `protobuf:"bytes,2,rep,name=board_roles,json=boardRoles,proto3" json:"board_roles,omitempty"`
	TaskRoles    []*TaskRole                  `protobuf:"bytes,5,rep,name=task_roles,json=taskRoles,proto3" json:"task_roles,omitempty"`
	TaskStatuses []*TaskStatus                `protobuf:"bytes,3,rep,name=task_statuses,json=taskStatuses,proto3" json:"task_statuses,omitempty"`
}

func (x *BoardDetailsResponse) Reset() {
	*x = BoardDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardDetailsResponse) ProtoMessage() {}

func (x *BoardDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardDetailsResponse.ProtoReflect.Descriptor instead.
func (*BoardDetailsResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{1}
}

func (x *BoardDetailsResponse) GetErrors() []BoardDetailsResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *BoardDetailsResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *BoardDetailsResponse) GetBoardRoles() []*BoardRole {
	if x != nil {
		return x.BoardRoles
	}
	return nil
}

func (x *BoardDetailsResponse) GetTaskRoles() []*TaskRole {
	if x != nil {
		return x.TaskRoles
	}
	return nil
}

func (x *BoardDetailsResponse) GetTaskStatuses() []*TaskStatus {
	if x != nil {
		return x.TaskStatuses
	}
	return nil
}

type CreateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *CreateBoardRequest) Reset() {
	*x = CreateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest) ProtoMessage() {}

func (x *CreateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBoardRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBoardRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type CreateBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []CreateBoardResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.CreateBoardResponse_Error" json:"errors,omitempty"`
	Chat   *Chat                       `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateBoardResponse) Reset() {
	*x = CreateBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardResponse) ProtoMessage() {}

func (x *CreateBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardResponse.ProtoReflect.Descriptor instead.
func (*CreateBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBoardResponse) GetErrors() []CreateBoardResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *CreateBoardResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type UpdateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateBoardRequest) Reset() {
	*x = UpdateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardRequest) ProtoMessage() {}

func (x *UpdateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoardRequest) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBoardRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *UpdateBoardRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []UpdateBoardResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.UpdateBoardResponse_Error" json:"errors,omitempty"`
	Chat   *Chat                       `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *UpdateBoardResponse) Reset() {
	*x = UpdateBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardResponse) ProtoMessage() {}

func (x *UpdateBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boards_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardResponse.ProtoReflect.Descriptor instead.
func (*UpdateBoardResponse) Descriptor() ([]byte, []int) {
	return file_boards_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBoardResponse) GetErrors() []UpdateBoardResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *UpdateBoardResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

var File_boards_proto protoreflect.FileDescriptor

var file_boards_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x13, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0xe9, 0x02, 0x0a, 0x14, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x22, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04,
	0x63, 0x68, 0x61, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x61, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0a,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x39, 0x0a,
	0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xc3, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49, 0x54, 0x4c,
	0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x02, 0x22, 0x43, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x04,
	0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74,
	0x22, 0x87, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x54,
	0x49, 0x54, 0x4c, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14,
	0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_boards_proto_rawDescOnce sync.Once
	file_boards_proto_rawDescData = file_boards_proto_rawDesc
)

func file_boards_proto_rawDescGZIP() []byte {
	file_boards_proto_rawDescOnce.Do(func() {
		file_boards_proto_rawDescData = protoimpl.X.CompressGZIP(file_boards_proto_rawDescData)
	})
	return file_boards_proto_rawDescData
}

var file_boards_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_boards_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_boards_proto_goTypes = []interface{}{
	(BoardDetailsResponse_Error)(0), // 0: teaproto.BoardDetailsResponse.Error
	(CreateBoardResponse_Error)(0),  // 1: teaproto.CreateBoardResponse.Error
	(UpdateBoardResponse_Error)(0),  // 2: teaproto.UpdateBoardResponse.Error
	(*BoardDetailsRequest)(nil),     // 3: teaproto.BoardDetailsRequest
	(*BoardDetailsResponse)(nil),    // 4: teaproto.BoardDetailsResponse
	(*CreateBoardRequest)(nil),      // 5: teaproto.CreateBoardRequest
	(*CreateBoardResponse)(nil),     // 6: teaproto.CreateBoardResponse
	(*UpdateBoardRequest)(nil),      // 7: teaproto.UpdateBoardRequest
	(*UpdateBoardResponse)(nil),     // 8: teaproto.UpdateBoardResponse
	(*Chat)(nil),                    // 9: teaproto.Chat
	(*BoardRole)(nil),               // 10: teaproto.BoardRole
	(*TaskRole)(nil),                // 11: teaproto.TaskRole
	(*TaskStatus)(nil),              // 12: teaproto.TaskStatus
}
var file_boards_proto_depIdxs = []int32{
	0,  // 0: teaproto.BoardDetailsResponse.errors:type_name -> teaproto.BoardDetailsResponse.Error
	9,  // 1: teaproto.BoardDetailsResponse.chat:type_name -> teaproto.Chat
	10, // 2: teaproto.BoardDetailsResponse.board_roles:type_name -> teaproto.BoardRole
	11, // 3: teaproto.BoardDetailsResponse.task_roles:type_name -> teaproto.TaskRole
	12, // 4: teaproto.BoardDetailsResponse.task_statuses:type_name -> teaproto.TaskStatus
	1,  // 5: teaproto.CreateBoardResponse.errors:type_name -> teaproto.CreateBoardResponse.Error
	9,  // 6: teaproto.CreateBoardResponse.chat:type_name -> teaproto.Chat
	2,  // 7: teaproto.UpdateBoardResponse.errors:type_name -> teaproto.UpdateBoardResponse.Error
	9,  // 8: teaproto.UpdateBoardResponse.chat:type_name -> teaproto.Chat
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_boards_proto_init() }
func file_boards_proto_init() {
	if File_boards_proto != nil {
		return
	}
	file_core_proto_init()
	file_core_boards_proto_init()
	file_core_tasks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_boards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardDetailsRequest); i {
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
		file_boards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardDetailsResponse); i {
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
		file_boards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardRequest); i {
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
		file_boards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardResponse); i {
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
		file_boards_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardRequest); i {
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
		file_boards_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardResponse); i {
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
			RawDescriptor: file_boards_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_boards_proto_goTypes,
		DependencyIndexes: file_boards_proto_depIdxs,
		EnumInfos:         file_boards_proto_enumTypes,
		MessageInfos:      file_boards_proto_msgTypes,
	}.Build()
	File_boards_proto = out.File
	file_boards_proto_rawDesc = nil
	file_boards_proto_goTypes = nil
	file_boards_proto_depIdxs = nil
}
