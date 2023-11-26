// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: core_images.proto

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width    uint32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height   uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	BlurHash string `protobuf:"bytes,4,opt,name=blur_hash,json=blurHash,proto3" json:"blur_hash,omitempty"` // https://blurha.sh/ (optional)
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_core_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_core_images_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Image) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Image) GetBlurHash() string {
	if x != nil {
		return x.BlurHash
	}
	return ""
}

type Icon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Letters string `protobuf:"bytes,1,opt,name=letters,proto3" json:"letters,omitempty"` // example: AB
	Color   string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`     // example: #FF9671
	Image   *Image `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`     // may be empty
}

func (x *Icon) Reset() {
	*x = Icon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Icon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Icon) ProtoMessage() {}

func (x *Icon) ProtoReflect() protoreflect.Message {
	mi := &file_core_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Icon.ProtoReflect.Descriptor instead.
func (*Icon) Descriptor() ([]byte, []int) {
	return file_core_images_proto_rawDescGZIP(), []int{1}
}

func (x *Icon) GetLetters() string {
	if x != nil {
		return x.Letters
	}
	return ""
}

func (x *Icon) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Icon) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_core_images_proto protoreflect.FileDescriptor

var file_core_images_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x72, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6c, 0x75, 0x72, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x5d, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x61,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_images_proto_rawDescOnce sync.Once
	file_core_images_proto_rawDescData = file_core_images_proto_rawDesc
)

func file_core_images_proto_rawDescGZIP() []byte {
	file_core_images_proto_rawDescOnce.Do(func() {
		file_core_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_images_proto_rawDescData)
	})
	return file_core_images_proto_rawDescData
}

var file_core_images_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_images_proto_goTypes = []interface{}{
	(*Image)(nil), // 0: teaproto.Image
	(*Icon)(nil),  // 1: teaproto.Icon
}
var file_core_images_proto_depIdxs = []int32{
	0, // 0: teaproto.Icon.image:type_name -> teaproto.Image
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_images_proto_init() }
func file_core_images_proto_init() {
	if File_core_images_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_core_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Icon); i {
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
			RawDescriptor: file_core_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_images_proto_goTypes,
		DependencyIndexes: file_core_images_proto_depIdxs,
		MessageInfos:      file_core_images_proto_msgTypes,
	}.Build()
	File_core_images_proto = out.File
	file_core_images_proto_rawDesc = nil
	file_core_images_proto_goTypes = nil
	file_core_images_proto_depIdxs = nil
}
