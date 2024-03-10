// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.2
// source: uploads.proto

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

type UploadType int32

const (
	UploadType_UT_FILE  UploadType = 0
	UploadType_UT_IMAGE UploadType = 1
	UploadType_UT_VIDEO UploadType = 2
	UploadType_UT_AUDIO UploadType = 3
)

// Enum value maps for UploadType.
var (
	UploadType_name = map[int32]string{
		0: "UT_FILE",
		1: "UT_IMAGE",
		2: "UT_VIDEO",
		3: "UT_AUDIO",
	}
	UploadType_value = map[string]int32{
		"UT_FILE":  0,
		"UT_IMAGE": 1,
		"UT_VIDEO": 2,
		"UT_AUDIO": 3,
	}
)

func (x UploadType) Enum() *UploadType {
	p := new(UploadType)
	*p = x
	return p
}

func (x UploadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadType) Descriptor() protoreflect.EnumDescriptor {
	return file_uploads_proto_enumTypes[0].Descriptor()
}

func (UploadType) Type() protoreflect.EnumType {
	return &file_uploads_proto_enumTypes[0]
}

func (x UploadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadType.Descriptor instead.
func (UploadType) EnumDescriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{0}
}

type UploadResponse_Error int32

const (
	UploadResponse_UNKNOWN_ERROR          UploadResponse_Error = 0
	UploadResponse_FILE_TOO_BIG_ERROR     UploadResponse_Error = 1 // see StateResponse.max_file_size
	UploadResponse_INVALID_FILENAME_ERROR UploadResponse_Error = 2 // filename length must be between 1 and 255 and must not contain character /
)

// Enum value maps for UploadResponse_Error.
var (
	UploadResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "FILE_TOO_BIG_ERROR",
		2: "INVALID_FILENAME_ERROR",
	}
	UploadResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":          0,
		"FILE_TOO_BIG_ERROR":     1,
		"INVALID_FILENAME_ERROR": 2,
	}
)

func (x UploadResponse_Error) Enum() *UploadResponse_Error {
	p := new(UploadResponse_Error)
	*p = x
	return p
}

func (x UploadResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_uploads_proto_enumTypes[1].Descriptor()
}

func (UploadResponse_Error) Type() protoreflect.EnumType {
	return &file_uploads_proto_enumTypes[1]
}

func (x UploadResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadResponse_Error.Descriptor instead.
func (UploadResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{2, 0}
}

type ChatUploadsResponse_Error int32

const (
	ChatUploadsResponse_UNKNOWN_ERROR       ChatUploadsResponse_Error = 0
	ChatUploadsResponse_INVALID_CHAT_ERROR  ChatUploadsResponse_Error = 2
	ChatUploadsResponse_ACCESS_DENIED_ERROR ChatUploadsResponse_Error = 3
	ChatUploadsResponse_INVALID_LIMIT_ERROR ChatUploadsResponse_Error = 1
)

// Enum value maps for ChatUploadsResponse_Error.
var (
	ChatUploadsResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		2: "INVALID_CHAT_ERROR",
		3: "ACCESS_DENIED_ERROR",
		1: "INVALID_LIMIT_ERROR",
	}
	ChatUploadsResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":       0,
		"INVALID_CHAT_ERROR":  2,
		"ACCESS_DENIED_ERROR": 3,
		"INVALID_LIMIT_ERROR": 1,
	}
)

func (x ChatUploadsResponse_Error) Enum() *ChatUploadsResponse_Error {
	p := new(ChatUploadsResponse_Error)
	*p = x
	return p
}

func (x ChatUploadsResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatUploadsResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_uploads_proto_enumTypes[2].Descriptor()
}

func (ChatUploadsResponse_Error) Type() protoreflect.EnumType {
	return &file_uploads_proto_enumTypes[2]
}

func (x ChatUploadsResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatUploadsResponse_Error.Descriptor instead.
func (ChatUploadsResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{4, 0}
}

type UploadIconResponse_Error int32

const (
	UploadIconResponse_UNKNOWN_ERROR              UploadIconResponse_Error = 0
	UploadIconResponse_FILE_TOO_BIG_ERROR         UploadIconResponse_Error = 1 // see StateResponse.max_file_size
	UploadIconResponse_INVALID_IMAGE_FORMAT_ERROR UploadIconResponse_Error = 2
	UploadIconResponse_IMAGE_TOO_SMALL_ERROR      UploadIconResponse_Error = 3 // see StateResponse.min_icon_size
	UploadIconResponse_NOT_ALLOWED_ERROR          UploadIconResponse_Error = 4
)

// Enum value maps for UploadIconResponse_Error.
var (
	UploadIconResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "FILE_TOO_BIG_ERROR",
		2: "INVALID_IMAGE_FORMAT_ERROR",
		3: "IMAGE_TOO_SMALL_ERROR",
		4: "NOT_ALLOWED_ERROR",
	}
	UploadIconResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":              0,
		"FILE_TOO_BIG_ERROR":         1,
		"INVALID_IMAGE_FORMAT_ERROR": 2,
		"IMAGE_TOO_SMALL_ERROR":      3,
		"NOT_ALLOWED_ERROR":          4,
	}
)

func (x UploadIconResponse_Error) Enum() *UploadIconResponse_Error {
	p := new(UploadIconResponse_Error)
	*p = x
	return p
}

func (x UploadIconResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadIconResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_uploads_proto_enumTypes[3].Descriptor()
}

func (UploadIconResponse_Error) Type() protoreflect.EnumType {
	return &file_uploads_proto_enumTypes[3]
}

func (x UploadIconResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadIconResponse_Error.Descriptor instead.
func (UploadIconResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{6, 0}
}

type Upload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt    uint64     `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Url          string     `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Filename     string     `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Type         UploadType `protobuf:"varint,4,opt,name=type,proto3,enum=teaproto.UploadType" json:"type,omitempty"`
	Size         uint32     `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	LargePreview *Image     `protobuf:"bytes,6,opt,name=large_preview,json=largePreview,proto3" json:"large_preview,omitempty"`  // optional large preview: for UT_IMAGE and UT_VIDEO
	SmallPreview *Image     `protobuf:"bytes,10,opt,name=small_preview,json=smallPreview,proto3" json:"small_preview,omitempty"` // optional small square preview: for UT_IMAGE and UT_VIDEO
	Duration     uint32     `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`                             // non-zero for UT_AUDIO and UT_VIDEO
	ContentType  string     `protobuf:"bytes,9,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`     // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
}

func (x *Upload) Reset() {
	*x = Upload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upload) ProtoMessage() {}

func (x *Upload) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upload.ProtoReflect.Descriptor instead.
func (*Upload) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{0}
}

func (x *Upload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Upload) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Upload) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Upload) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Upload) GetType() UploadType {
	if x != nil {
		return x.Type
	}
	return UploadType_UT_FILE
}

func (x *Upload) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Upload) GetLargePreview() *Image {
	if x != nil {
		return x.LargePreview
	}
	return nil
}

func (x *Upload) GetSmallPreview() *Image {
	if x != nil {
		return x.SmallPreview
	}
	return nil
}

func (x *Upload) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Upload) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content              []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Filename             string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	DisableModifications bool   `protobuf:"varint,3,opt,name=disable_modifications,json=disableModifications,proto3" json:"disable_modifications,omitempty"` // don't make any modifications to uploaded file: resize, rotate, convert, etc.
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{1}
}

func (x *UploadRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadRequest) GetDisableModifications() bool {
	if x != nil {
		return x.DisableModifications
	}
	return false
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []UploadResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.UploadResponse_Error" json:"errors,omitempty"`
	Upload *Upload                `protobuf:"bytes,2,opt,name=upload,proto3" json:"upload,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetErrors() []UploadResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *UploadResponse) GetUpload() *Upload {
	if x != nil {
		return x.Upload
	}
	return nil
}

type ChatUploadsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string       `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Limit  uint32       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"` // default: StateResponse.max_results_on_page
	Offset uint32       `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Types  []UploadType `protobuf:"varint,5,rep,packed,name=types,proto3,enum=teaproto.UploadType" json:"types,omitempty"`
}

func (x *ChatUploadsRequest) Reset() {
	*x = ChatUploadsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatUploadsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatUploadsRequest) ProtoMessage() {}

func (x *ChatUploadsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatUploadsRequest.ProtoReflect.Descriptor instead.
func (*ChatUploadsRequest) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{3}
}

func (x *ChatUploadsRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ChatUploadsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ChatUploadsRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ChatUploadsRequest) GetTypes() []UploadType {
	if x != nil {
		return x.Types
	}
	return nil
}

type ChatUploadsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors  []ChatUploadsResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.ChatUploadsResponse_Error" json:"errors,omitempty"`
	Uploads []*Upload                   `protobuf:"bytes,2,rep,name=uploads,proto3" json:"uploads,omitempty"`
}

func (x *ChatUploadsResponse) Reset() {
	*x = ChatUploadsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatUploadsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatUploadsResponse) ProtoMessage() {}

func (x *ChatUploadsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatUploadsResponse.ProtoReflect.Descriptor instead.
func (*ChatUploadsResponse) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{4}
}

func (x *ChatUploadsResponse) GetErrors() []ChatUploadsResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ChatUploadsResponse) GetUploads() []*Upload {
	if x != nil {
		return x.Uploads
	}
	return nil
}

type UploadIconRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Subject:
	//
	//	*UploadIconRequest_UserId
	//	*UploadIconRequest_ChatId
	Subject isUploadIconRequest_Subject `protobuf_oneof:"subject"`
	Content []byte                      `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"` // delete when empty
}

func (x *UploadIconRequest) Reset() {
	*x = UploadIconRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadIconRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadIconRequest) ProtoMessage() {}

func (x *UploadIconRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadIconRequest.ProtoReflect.Descriptor instead.
func (*UploadIconRequest) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{5}
}

func (m *UploadIconRequest) GetSubject() isUploadIconRequest_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *UploadIconRequest) GetUserId() string {
	if x, ok := x.GetSubject().(*UploadIconRequest_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *UploadIconRequest) GetChatId() string {
	if x, ok := x.GetSubject().(*UploadIconRequest_ChatId); ok {
		return x.ChatId
	}
	return ""
}

func (x *UploadIconRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type isUploadIconRequest_Subject interface {
	isUploadIconRequest_Subject()
}

type UploadIconRequest_UserId struct {
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3,oneof"` // uuid or "me"
}

type UploadIconRequest_ChatId struct {
	ChatId string `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3,oneof"`
}

func (*UploadIconRequest_UserId) isUploadIconRequest_Subject() {}

func (*UploadIconRequest_ChatId) isUploadIconRequest_Subject() {}

type UploadIconResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []UploadIconResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.UploadIconResponse_Error" json:"errors,omitempty"`
	Icon   *Icon                      `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *UploadIconResponse) Reset() {
	*x = UploadIconResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploads_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadIconResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadIconResponse) ProtoMessage() {}

func (x *UploadIconResponse) ProtoReflect() protoreflect.Message {
	mi := &file_uploads_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadIconResponse.ProtoReflect.Descriptor instead.
func (*UploadIconResponse) Descriptor() ([]byte, []int) {
	return file_uploads_proto_rawDescGZIP(), []int{6}
}

func (x *UploadIconResponse) GetErrors() []UploadIconResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *UploadIconResponse) GetIcon() *Icon {
	if x != nil {
		return x.Icon
	}
	return nil
}

var File_uploads_proto protoreflect.FileDescriptor

var file_uploads_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02,
	0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x61,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x34, 0x0a, 0x0d, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x0c, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7a,
	0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x0e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x4e, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x42, 0x49, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46,
	0x49, 0x4c, 0x45, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22,
	0x87, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2a, 0x0a,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74,
	0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x13, 0x43, 0x68,
	0x61, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x2a,
	0x0a, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x64, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01,
	0x22, 0x6e, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0xfb, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x63, 0x6f,
	0x6e, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4f,
	0x5f, 0x42, 0x49, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x4f, 0x54, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x2a, 0x43,
	0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x54, 0x5f,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x54, 0x5f, 0x56, 0x49,
	0x44, 0x45, 0x4f, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x54, 0x5f, 0x41, 0x55, 0x44, 0x49,
	0x4f, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uploads_proto_rawDescOnce sync.Once
	file_uploads_proto_rawDescData = file_uploads_proto_rawDesc
)

func file_uploads_proto_rawDescGZIP() []byte {
	file_uploads_proto_rawDescOnce.Do(func() {
		file_uploads_proto_rawDescData = protoimpl.X.CompressGZIP(file_uploads_proto_rawDescData)
	})
	return file_uploads_proto_rawDescData
}

var file_uploads_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_uploads_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_uploads_proto_goTypes = []interface{}{
	(UploadType)(0),                // 0: teaproto.UploadType
	(UploadResponse_Error)(0),      // 1: teaproto.UploadResponse.Error
	(ChatUploadsResponse_Error)(0), // 2: teaproto.ChatUploadsResponse.Error
	(UploadIconResponse_Error)(0),  // 3: teaproto.UploadIconResponse.Error
	(*Upload)(nil),                 // 4: teaproto.Upload
	(*UploadRequest)(nil),          // 5: teaproto.UploadRequest
	(*UploadResponse)(nil),         // 6: teaproto.UploadResponse
	(*ChatUploadsRequest)(nil),     // 7: teaproto.ChatUploadsRequest
	(*ChatUploadsResponse)(nil),    // 8: teaproto.ChatUploadsResponse
	(*UploadIconRequest)(nil),      // 9: teaproto.UploadIconRequest
	(*UploadIconResponse)(nil),     // 10: teaproto.UploadIconResponse
	(*Image)(nil),                  // 11: teaproto.Image
	(*Icon)(nil),                   // 12: teaproto.Icon
}
var file_uploads_proto_depIdxs = []int32{
	0,  // 0: teaproto.Upload.type:type_name -> teaproto.UploadType
	11, // 1: teaproto.Upload.large_preview:type_name -> teaproto.Image
	11, // 2: teaproto.Upload.small_preview:type_name -> teaproto.Image
	1,  // 3: teaproto.UploadResponse.errors:type_name -> teaproto.UploadResponse.Error
	4,  // 4: teaproto.UploadResponse.upload:type_name -> teaproto.Upload
	0,  // 5: teaproto.ChatUploadsRequest.types:type_name -> teaproto.UploadType
	2,  // 6: teaproto.ChatUploadsResponse.errors:type_name -> teaproto.ChatUploadsResponse.Error
	4,  // 7: teaproto.ChatUploadsResponse.uploads:type_name -> teaproto.Upload
	3,  // 8: teaproto.UploadIconResponse.errors:type_name -> teaproto.UploadIconResponse.Error
	12, // 9: teaproto.UploadIconResponse.icon:type_name -> teaproto.Icon
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_uploads_proto_init() }
func file_uploads_proto_init() {
	if File_uploads_proto != nil {
		return
	}
	file_types_images_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_uploads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upload); i {
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
		file_uploads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_uploads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_uploads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatUploadsRequest); i {
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
		file_uploads_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatUploadsResponse); i {
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
		file_uploads_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadIconRequest); i {
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
		file_uploads_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadIconResponse); i {
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
	file_uploads_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*UploadIconRequest_UserId)(nil),
		(*UploadIconRequest_ChatId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uploads_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uploads_proto_goTypes,
		DependencyIndexes: file_uploads_proto_depIdxs,
		EnumInfos:         file_uploads_proto_enumTypes,
		MessageInfos:      file_uploads_proto_msgTypes,
	}.Build()
	File_uploads_proto = out.File
	file_uploads_proto_rawDesc = nil
	file_uploads_proto_goTypes = nil
	file_uploads_proto_depIdxs = nil
}
