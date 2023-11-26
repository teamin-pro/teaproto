// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: tasks.proto

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

type TaskListResponse_Error int32

const (
	TaskListResponse_UNKNOWN_ERROR       TaskListResponse_Error = 0
	TaskListResponse_INVALID_CHAT_ERROR  TaskListResponse_Error = 1
	TaskListResponse_ACCESS_DENIED_ERROR TaskListResponse_Error = 3
	TaskListResponse_INVALID_LIMIT_ERROR TaskListResponse_Error = 2
)

// Enum value maps for TaskListResponse_Error.
var (
	TaskListResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		3: "ACCESS_DENIED_ERROR",
		2: "INVALID_LIMIT_ERROR",
	}
	TaskListResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":       0,
		"INVALID_CHAT_ERROR":  1,
		"ACCESS_DENIED_ERROR": 3,
		"INVALID_LIMIT_ERROR": 2,
	}
)

func (x TaskListResponse_Error) Enum() *TaskListResponse_Error {
	p := new(TaskListResponse_Error)
	*p = x
	return p
}

func (x TaskListResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskListResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_tasks_proto_enumTypes[0].Descriptor()
}

func (TaskListResponse_Error) Type() protoreflect.EnumType {
	return &file_tasks_proto_enumTypes[0]
}

func (x TaskListResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskListResponse_Error.Descriptor instead.
func (TaskListResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{1, 0}
}

type CreateTaskResponse_Error int32

const (
	CreateTaskResponse_UNKNOWN_ERROR                     CreateTaskResponse_Error = 0
	CreateTaskResponse_INVALID_PARENT_CHAT_ERROR         CreateTaskResponse_Error = 1
	CreateTaskResponse_TITLE_AND_DESCRIPTION_EMPTY_ERROR CreateTaskResponse_Error = 2
	CreateTaskResponse_TITLE_IS_TOO_LONG_ERROR           CreateTaskResponse_Error = 3 // see: StateResponse.max_task_title_length
	CreateTaskResponse_DESCRIPTION_IS_TOO_LONG_ERROR     CreateTaskResponse_Error = 4 // see: StateResponse.max_task_description_length
	CreateTaskResponse_TOO_MANY_TASK_MEMBERS             CreateTaskResponse_Error = 5 // see: StateResponse.max_task_members
	CreateTaskResponse_NOT_ALLOWED_TO_CREATE_TASK_ERROR  CreateTaskResponse_Error = 6
)

// Enum value maps for CreateTaskResponse_Error.
var (
	CreateTaskResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_PARENT_CHAT_ERROR",
		2: "TITLE_AND_DESCRIPTION_EMPTY_ERROR",
		3: "TITLE_IS_TOO_LONG_ERROR",
		4: "DESCRIPTION_IS_TOO_LONG_ERROR",
		5: "TOO_MANY_TASK_MEMBERS",
		6: "NOT_ALLOWED_TO_CREATE_TASK_ERROR",
	}
	CreateTaskResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":                     0,
		"INVALID_PARENT_CHAT_ERROR":         1,
		"TITLE_AND_DESCRIPTION_EMPTY_ERROR": 2,
		"TITLE_IS_TOO_LONG_ERROR":           3,
		"DESCRIPTION_IS_TOO_LONG_ERROR":     4,
		"TOO_MANY_TASK_MEMBERS":             5,
		"NOT_ALLOWED_TO_CREATE_TASK_ERROR":  6,
	}
)

func (x CreateTaskResponse_Error) Enum() *CreateTaskResponse_Error {
	p := new(CreateTaskResponse_Error)
	*p = x
	return p
}

func (x CreateTaskResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateTaskResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_tasks_proto_enumTypes[1].Descriptor()
}

func (CreateTaskResponse_Error) Type() protoreflect.EnumType {
	return &file_tasks_proto_enumTypes[1]
}

func (x CreateTaskResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateTaskResponse_Error.Descriptor instead.
func (CreateTaskResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{3, 0}
}

type UpdateTaskResponse_Error int32

const (
	UpdateTaskResponse_UNKNOWN_ERROR                     UpdateTaskResponse_Error = 0
	UpdateTaskResponse_INVALID_CHAT_ERROR                UpdateTaskResponse_Error = 1
	UpdateTaskResponse_ACCESS_DENIED_ERROR               UpdateTaskResponse_Error = 2
	UpdateTaskResponse_TITLE_IS_TOO_LONG_ERROR           UpdateTaskResponse_Error = 3 // see: StateResponse.max_task_title_length
	UpdateTaskResponse_DESCRIPTION_IS_TOO_LONG_ERROR     UpdateTaskResponse_Error = 4 // see: StateResponse.max_task_description_length
	UpdateTaskResponse_INVALID_TASK_STATUS_ID_ERROR      UpdateTaskResponse_Error = 7
	UpdateTaskResponse_TITLE_AND_DESCRIPTION_EMPTY_ERROR UpdateTaskResponse_Error = 5
)

// Enum value maps for UpdateTaskResponse_Error.
var (
	UpdateTaskResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "ACCESS_DENIED_ERROR",
		3: "TITLE_IS_TOO_LONG_ERROR",
		4: "DESCRIPTION_IS_TOO_LONG_ERROR",
		7: "INVALID_TASK_STATUS_ID_ERROR",
		5: "TITLE_AND_DESCRIPTION_EMPTY_ERROR",
	}
	UpdateTaskResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":                     0,
		"INVALID_CHAT_ERROR":                1,
		"ACCESS_DENIED_ERROR":               2,
		"TITLE_IS_TOO_LONG_ERROR":           3,
		"DESCRIPTION_IS_TOO_LONG_ERROR":     4,
		"INVALID_TASK_STATUS_ID_ERROR":      7,
		"TITLE_AND_DESCRIPTION_EMPTY_ERROR": 5,
	}
)

func (x UpdateTaskResponse_Error) Enum() *UpdateTaskResponse_Error {
	p := new(UpdateTaskResponse_Error)
	*p = x
	return p
}

func (x UpdateTaskResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateTaskResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_tasks_proto_enumTypes[2].Descriptor()
}

func (UpdateTaskResponse_Error) Type() protoreflect.EnumType {
	return &file_tasks_proto_enumTypes[2]
}

func (x UpdateTaskResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateTaskResponse_Error.Descriptor instead.
func (UpdateTaskResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{5, 0}
}

type DeleteTaskResponse_Error int32

const (
	DeleteTaskResponse_UNKNOWN_ERROR       DeleteTaskResponse_Error = 0
	DeleteTaskResponse_INVALID_CHAT_ERROR  DeleteTaskResponse_Error = 1
	DeleteTaskResponse_ACCESS_DENIED_ERROR DeleteTaskResponse_Error = 2
)

// Enum value maps for DeleteTaskResponse_Error.
var (
	DeleteTaskResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_CHAT_ERROR",
		2: "ACCESS_DENIED_ERROR",
	}
	DeleteTaskResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":       0,
		"INVALID_CHAT_ERROR":  1,
		"ACCESS_DENIED_ERROR": 2,
	}
)

func (x DeleteTaskResponse_Error) Enum() *DeleteTaskResponse_Error {
	p := new(DeleteTaskResponse_Error)
	*p = x
	return p
}

func (x DeleteTaskResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteTaskResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_tasks_proto_enumTypes[3].Descriptor()
}

func (DeleteTaskResponse_Error) Type() protoreflect.EnumType {
	return &file_tasks_proto_enumTypes[3]
}

func (x DeleteTaskResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteTaskResponse_Error.Descriptor instead.
func (DeleteTaskResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{7, 0}
}

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentChatId  string   `protobuf:"bytes,1,opt,name=parent_chat_id,json=parentChatId,proto3" json:"parent_chat_id,omitempty"`            // any task should have a parent chat for now
	TaskStatusIds []uint32 `protobuf:"varint,2,rep,packed,name=task_status_ids,json=taskStatusIds,proto3" json:"task_status_ids,omitempty"` // if empty, all statuses are returned
	Limit         uint32   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`                                               // default: StateResponse.max_results_on_page
	Offset        uint32   `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *TaskListRequest) GetParentChatId() string {
	if x != nil {
		return x.ParentChatId
	}
	return ""
}

func (x *TaskListRequest) GetTaskStatusIds() []uint32 {
	if x != nil {
		return x.TaskStatusIds
	}
	return nil
}

func (x *TaskListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TaskListRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []TaskListResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.TaskListResponse_Error" json:"errors,omitempty"`
	Chats  []*Chat                  `protobuf:"bytes,2,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *TaskListResponse) GetErrors() []TaskListResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TaskListResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentChatId string              `protobuf:"bytes,1,opt,name=parent_chat_id,json=parentChatId,proto3" json:"parent_chat_id,omitempty"`
	Title        string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string              `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Members      []*TaskMemberUpdate `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskRequest) GetParentChatId() string {
	if x != nil {
		return x.ParentChatId
	}
	return ""
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetMembers() []*TaskMemberUpdate {
	if x != nil {
		return x.Members
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []CreateTaskResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.CreateTaskResponse_Error" json:"errors,omitempty"`
	Chat   *Chat                      `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"` // TODO: create from message
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTaskResponse) GetErrors() []CreateTaskResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *CreateTaskResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId            string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UpdateTitle       bool   `protobuf:"varint,2,opt,name=update_title,json=updateTitle,proto3" json:"update_title,omitempty"`
	Title             string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	UpdateDescription bool   `protobuf:"varint,4,opt,name=update_description,json=updateDescription,proto3" json:"update_description,omitempty"`
	Description       string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	TaskStatusId      uint32 `protobuf:"varint,7,opt,name=task_status_id,json=taskStatusId,proto3" json:"task_status_id,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTaskRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *UpdateTaskRequest) GetUpdateTitle() bool {
	if x != nil {
		return x.UpdateTitle
	}
	return false
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetUpdateDescription() bool {
	if x != nil {
		return x.UpdateDescription
	}
	return false
}

func (x *UpdateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTaskRequest) GetTaskStatusId() uint32 {
	if x != nil {
		return x.TaskStatusId
	}
	return 0
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []UpdateTaskResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.UpdateTaskResponse_Error" json:"errors,omitempty"`
	Chat   *Chat                      `protobuf:"bytes,2,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTaskResponse) GetErrors() []UpdateTaskResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *UpdateTaskResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTaskRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []DeleteTaskResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.DeleteTaskResponse_Error" json:"errors,omitempty"`
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTaskResponse) GetErrors() []DeleteTaskResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_tasks_proto protoreflect.FileDescriptor

var file_tasks_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74,
	0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x61,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x22, 0x64, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x22, 0xa7, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0xd8, 0x02, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65,
	0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61,
	0x74, 0x22, 0xe1, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x25, 0x0a,
	0x21, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52,
	0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x49, 0x53,
	0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x49, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x10, 0x05, 0x12,
	0x24, 0x0a, 0x20, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x54,
	0x4f, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x06, 0x22, 0xdc, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a,
	0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x49, 0x64, 0x22, 0xcb, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65,
	0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0xd4, 0x01, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45,
	0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x49, 0x54,
	0x4c, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49,
	0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e,
	0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x49, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21, 0x54,
	0x49, 0x54, 0x4c, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x05, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x22, 0x9d, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tasks_proto_rawDescOnce sync.Once
	file_tasks_proto_rawDescData = file_tasks_proto_rawDesc
)

func file_tasks_proto_rawDescGZIP() []byte {
	file_tasks_proto_rawDescOnce.Do(func() {
		file_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_tasks_proto_rawDescData)
	})
	return file_tasks_proto_rawDescData
}

var file_tasks_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tasks_proto_goTypes = []interface{}{
	(TaskListResponse_Error)(0),   // 0: teaproto.TaskListResponse.Error
	(CreateTaskResponse_Error)(0), // 1: teaproto.CreateTaskResponse.Error
	(UpdateTaskResponse_Error)(0), // 2: teaproto.UpdateTaskResponse.Error
	(DeleteTaskResponse_Error)(0), // 3: teaproto.DeleteTaskResponse.Error
	(*TaskListRequest)(nil),       // 4: teaproto.TaskListRequest
	(*TaskListResponse)(nil),      // 5: teaproto.TaskListResponse
	(*CreateTaskRequest)(nil),     // 6: teaproto.CreateTaskRequest
	(*CreateTaskResponse)(nil),    // 7: teaproto.CreateTaskResponse
	(*UpdateTaskRequest)(nil),     // 8: teaproto.UpdateTaskRequest
	(*UpdateTaskResponse)(nil),    // 9: teaproto.UpdateTaskResponse
	(*DeleteTaskRequest)(nil),     // 10: teaproto.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),    // 11: teaproto.DeleteTaskResponse
	(*Chat)(nil),                  // 12: teaproto.Chat
	(*TaskMemberUpdate)(nil),      // 13: teaproto.TaskMemberUpdate
}
var file_tasks_proto_depIdxs = []int32{
	0,  // 0: teaproto.TaskListResponse.errors:type_name -> teaproto.TaskListResponse.Error
	12, // 1: teaproto.TaskListResponse.chats:type_name -> teaproto.Chat
	13, // 2: teaproto.CreateTaskRequest.members:type_name -> teaproto.TaskMemberUpdate
	1,  // 3: teaproto.CreateTaskResponse.errors:type_name -> teaproto.CreateTaskResponse.Error
	12, // 4: teaproto.CreateTaskResponse.chat:type_name -> teaproto.Chat
	2,  // 5: teaproto.UpdateTaskResponse.errors:type_name -> teaproto.UpdateTaskResponse.Error
	12, // 6: teaproto.UpdateTaskResponse.chat:type_name -> teaproto.Chat
	3,  // 7: teaproto.DeleteTaskResponse.errors:type_name -> teaproto.DeleteTaskResponse.Error
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_tasks_proto_init() }
func file_tasks_proto_init() {
	if File_tasks_proto != nil {
		return
	}
	file_core_proto_init()
	file_core_tasks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListRequest); i {
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
		file_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse); i {
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
		file_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
		file_tasks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_tasks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskResponse); i {
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
		file_tasks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_tasks_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskResponse); i {
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
			RawDescriptor: file_tasks_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tasks_proto_goTypes,
		DependencyIndexes: file_tasks_proto_depIdxs,
		EnumInfos:         file_tasks_proto_enumTypes,
		MessageInfos:      file_tasks_proto_msgTypes,
	}.Build()
	File_tasks_proto = out.File
	file_tasks_proto_rawDesc = nil
	file_tasks_proto_goTypes = nil
	file_tasks_proto_depIdxs = nil
}
