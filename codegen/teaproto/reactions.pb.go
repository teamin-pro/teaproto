// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: reactions.proto

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

type SetMessageReactionResponse_Error int32

const (
	SetMessageReactionResponse_UNKNOWN_ERROR            SetMessageReactionResponse_Error = 0
	SetMessageReactionResponse_INVALID_MESSAGE_ID_ERROR SetMessageReactionResponse_Error = 1
	SetMessageReactionResponse_INVALID_EMOJI_ERROR      SetMessageReactionResponse_Error = 2
)

// Enum value maps for SetMessageReactionResponse_Error.
var (
	SetMessageReactionResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_MESSAGE_ID_ERROR",
		2: "INVALID_EMOJI_ERROR",
	}
	SetMessageReactionResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":            0,
		"INVALID_MESSAGE_ID_ERROR": 1,
		"INVALID_EMOJI_ERROR":      2,
	}
)

func (x SetMessageReactionResponse_Error) Enum() *SetMessageReactionResponse_Error {
	p := new(SetMessageReactionResponse_Error)
	*p = x
	return p
}

func (x SetMessageReactionResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetMessageReactionResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_reactions_proto_enumTypes[0].Descriptor()
}

func (SetMessageReactionResponse_Error) Type() protoreflect.EnumType {
	return &file_reactions_proto_enumTypes[0]
}

func (x SetMessageReactionResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetMessageReactionResponse_Error.Descriptor instead.
func (SetMessageReactionResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{4, 0}
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emoji    string `protobuf:"bytes,1,opt,name=emoji,proto3" json:"emoji,omitempty"`                       // ☘️
	Slug     string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`                         // shamrock
	Group    string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`                       // Animals & Nature
	SubGroup string `protobuf:"bytes,4,opt,name=sub_group,json=subGroup,proto3" json:"sub_group,omitempty"` // plant-other
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{0}
}

func (x *Reaction) GetEmoji() string {
	if x != nil {
		return x.Emoji
	}
	return ""
}

func (x *Reaction) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Reaction) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Reaction) GetSubGroup() string {
	if x != nil {
		return x.SubGroup
	}
	return ""
}

type ReactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReactionsRequest) Reset() {
	*x = ReactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionsRequest) ProtoMessage() {}

func (x *ReactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionsRequest.ProtoReflect.Descriptor instead.
func (*ReactionsRequest) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{1}
}

type ReactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recent          []string    `protobuf:"bytes,1,rep,name=recent,proto3" json:"recent,omitempty"`
	All             []*Reaction `protobuf:"bytes,2,rep,name=all,proto3" json:"all,omitempty"`
	DefaultReaction string      `protobuf:"bytes,3,opt,name=default_reaction,json=defaultReaction,proto3" json:"default_reaction,omitempty"` // for double-tap action
}

func (x *ReactionsResponse) Reset() {
	*x = ReactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionsResponse) ProtoMessage() {}

func (x *ReactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionsResponse.ProtoReflect.Descriptor instead.
func (*ReactionsResponse) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{2}
}

func (x *ReactionsResponse) GetRecent() []string {
	if x != nil {
		return x.Recent
	}
	return nil
}

func (x *ReactionsResponse) GetAll() []*Reaction {
	if x != nil {
		return x.All
	}
	return nil
}

func (x *ReactionsResponse) GetDefaultReaction() string {
	if x != nil {
		return x.DefaultReaction
	}
	return ""
}

type SetMessageReactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Emoji     string `protobuf:"bytes,2,opt,name=emoji,proto3" json:"emoji,omitempty"`
	Remove    bool   `protobuf:"varint,3,opt,name=remove,proto3" json:"remove,omitempty"`
}

func (x *SetMessageReactionRequest) Reset() {
	*x = SetMessageReactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMessageReactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMessageReactionRequest) ProtoMessage() {}

func (x *SetMessageReactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMessageReactionRequest.ProtoReflect.Descriptor instead.
func (*SetMessageReactionRequest) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{3}
}

func (x *SetMessageReactionRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SetMessageReactionRequest) GetEmoji() string {
	if x != nil {
		return x.Emoji
	}
	return ""
}

func (x *SetMessageReactionRequest) GetRemove() bool {
	if x != nil {
		return x.Remove
	}
	return false
}

type SetMessageReactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors    []SetMessageReactionResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.SetMessageReactionResponse_Error" json:"errors,omitempty"`
	Reactions []*MessageReaction                 `protobuf:"bytes,2,rep,name=reactions,proto3" json:"reactions,omitempty"`
	// update
	Recent          []string `protobuf:"bytes,3,rep,name=recent,proto3" json:"recent,omitempty"`
	DefaultReaction string   `protobuf:"bytes,4,opt,name=default_reaction,json=defaultReaction,proto3" json:"default_reaction,omitempty"`
}

func (x *SetMessageReactionResponse) Reset() {
	*x = SetMessageReactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reactions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMessageReactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMessageReactionResponse) ProtoMessage() {}

func (x *SetMessageReactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reactions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMessageReactionResponse.ProtoReflect.Descriptor instead.
func (*SetMessageReactionResponse) Descriptor() ([]byte, []int) {
	return file_reactions_proto_rawDescGZIP(), []int{4}
}

func (x *SetMessageReactionResponse) GetErrors() []SetMessageReactionResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *SetMessageReactionResponse) GetReactions() []*MessageReaction {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *SetMessageReactionResponse) GetRecent() []string {
	if x != nil {
		return x.Recent
	}
	return nil
}

func (x *SetMessageReactionResponse) GetDefaultReaction() string {
	if x != nil {
		return x.DefaultReaction
	}
	return ""
}

var File_reactions_proto protoreflect.FileDescriptor

var file_reactions_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x7c, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x24, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0xaf, 0x02, 0x0a,
	0x1a, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x74, 0x65,
	0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12,
	0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x45, 0x4d, 0x4f, 0x4a, 0x49, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reactions_proto_rawDescOnce sync.Once
	file_reactions_proto_rawDescData = file_reactions_proto_rawDesc
)

func file_reactions_proto_rawDescGZIP() []byte {
	file_reactions_proto_rawDescOnce.Do(func() {
		file_reactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_reactions_proto_rawDescData)
	})
	return file_reactions_proto_rawDescData
}

var file_reactions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_reactions_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reactions_proto_goTypes = []interface{}{
	(SetMessageReactionResponse_Error)(0), // 0: teaproto.SetMessageReactionResponse.Error
	(*Reaction)(nil),                      // 1: teaproto.Reaction
	(*ReactionsRequest)(nil),              // 2: teaproto.ReactionsRequest
	(*ReactionsResponse)(nil),             // 3: teaproto.ReactionsResponse
	(*SetMessageReactionRequest)(nil),     // 4: teaproto.SetMessageReactionRequest
	(*SetMessageReactionResponse)(nil),    // 5: teaproto.SetMessageReactionResponse
	(*MessageReaction)(nil),               // 6: teaproto.MessageReaction
}
var file_reactions_proto_depIdxs = []int32{
	1, // 0: teaproto.ReactionsResponse.all:type_name -> teaproto.Reaction
	0, // 1: teaproto.SetMessageReactionResponse.errors:type_name -> teaproto.SetMessageReactionResponse.Error
	6, // 2: teaproto.SetMessageReactionResponse.reactions:type_name -> teaproto.MessageReaction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_reactions_proto_init() }
func file_reactions_proto_init() {
	if File_reactions_proto != nil {
		return
	}
	file_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
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
		file_reactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionsRequest); i {
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
		file_reactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionsResponse); i {
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
		file_reactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMessageReactionRequest); i {
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
		file_reactions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMessageReactionResponse); i {
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
			RawDescriptor: file_reactions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reactions_proto_goTypes,
		DependencyIndexes: file_reactions_proto_depIdxs,
		EnumInfos:         file_reactions_proto_enumTypes,
		MessageInfos:      file_reactions_proto_msgTypes,
	}.Build()
	File_reactions_proto = out.File
	file_reactions_proto_rawDesc = nil
	file_reactions_proto_goTypes = nil
	file_reactions_proto_depIdxs = nil
}
