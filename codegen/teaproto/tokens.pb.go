// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.2
// source: tokens.proto

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

type DeleteTokenResponse_Error int32

const (
	DeleteTokenResponse_UNKNOWN_ERROR         DeleteTokenResponse_Error = 0
	DeleteTokenResponse_TOKEN_NOT_FOUND_ERROR DeleteTokenResponse_Error = 2
)

// Enum value maps for DeleteTokenResponse_Error.
var (
	DeleteTokenResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		2: "TOKEN_NOT_FOUND_ERROR",
	}
	DeleteTokenResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":         0,
		"TOKEN_NOT_FOUND_ERROR": 2,
	}
)

func (x DeleteTokenResponse_Error) Enum() *DeleteTokenResponse_Error {
	p := new(DeleteTokenResponse_Error)
	*p = x
	return p
}

func (x DeleteTokenResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteTokenResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_tokens_proto_enumTypes[0].Descriptor()
}

func (DeleteTokenResponse_Error) Type() protoreflect.EnumType {
	return &file_tokens_proto_enumTypes[0]
}

func (x DeleteTokenResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteTokenResponse_Error.Descriptor instead.
func (DeleteTokenResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{4, 0}
}

// Auth token. TODO: rename to session?
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr       string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`           // ip address
	Useragent  string `protobuf:"bytes,3,opt,name=useragent,proto3" json:"useragent,omitempty"` // example: Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0
	LastUsedAt uint64 `protobuf:"varint,4,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Token) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Token) GetUseragent() string {
	if x != nil {
		return x.Useragent
	}
	return ""
}

func (x *Token) GetLastUsedAt() uint64 {
	if x != nil {
		return x.LastUsedAt
	}
	return 0
}

type TokensListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TokensListRequest) Reset() {
	*x = TokensListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensListRequest) ProtoMessage() {}

func (x *TokensListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensListRequest.ProtoReflect.Descriptor instead.
func (*TokensListRequest) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{1}
}

type TokensListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *TokensListResponse) Reset() {
	*x = TokensListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensListResponse) ProtoMessage() {}

func (x *TokensListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensListResponse.ProtoReflect.Descriptor instead.
func (*TokensListResponse) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{2}
}

func (x *TokensListResponse) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTokenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []DeleteTokenResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.DeleteTokenResponse_Error" json:"errors,omitempty"`
}

func (x *DeleteTokenResponse) Reset() {
	*x = DeleteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokens_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenResponse) ProtoMessage() {}

func (x *DeleteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokens_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenResponse.ProtoReflect.Descriptor instead.
func (*DeleteTokenResponse) Descriptor() ([]byte, []int) {
	return file_tokens_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTokenResponse) GetErrors() []DeleteTokenResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_tokens_proto protoreflect.FileDescriptor

var file_tokens_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x64, 0x41, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x12, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x89, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tokens_proto_rawDescOnce sync.Once
	file_tokens_proto_rawDescData = file_tokens_proto_rawDesc
)

func file_tokens_proto_rawDescGZIP() []byte {
	file_tokens_proto_rawDescOnce.Do(func() {
		file_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_tokens_proto_rawDescData)
	})
	return file_tokens_proto_rawDescData
}

var file_tokens_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tokens_proto_goTypes = []interface{}{
	(DeleteTokenResponse_Error)(0), // 0: teaproto.DeleteTokenResponse.Error
	(*Token)(nil),                  // 1: teaproto.Token
	(*TokensListRequest)(nil),      // 2: teaproto.TokensListRequest
	(*TokensListResponse)(nil),     // 3: teaproto.TokensListResponse
	(*DeleteTokenRequest)(nil),     // 4: teaproto.DeleteTokenRequest
	(*DeleteTokenResponse)(nil),    // 5: teaproto.DeleteTokenResponse
}
var file_tokens_proto_depIdxs = []int32{
	1, // 0: teaproto.TokensListResponse.tokens:type_name -> teaproto.Token
	0, // 1: teaproto.DeleteTokenResponse.errors:type_name -> teaproto.DeleteTokenResponse.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tokens_proto_init() }
func file_tokens_proto_init() {
	if File_tokens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensListRequest); i {
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
		file_tokens_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensListResponse); i {
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
		file_tokens_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTokenRequest); i {
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
		file_tokens_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTokenResponse); i {
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
			RawDescriptor: file_tokens_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tokens_proto_goTypes,
		DependencyIndexes: file_tokens_proto_depIdxs,
		EnumInfos:         file_tokens_proto_enumTypes,
		MessageInfos:      file_tokens_proto_msgTypes,
	}.Build()
	File_tokens_proto = out.File
	file_tokens_proto_rawDesc = nil
	file_tokens_proto_goTypes = nil
	file_tokens_proto_depIdxs = nil
}
