// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: devices.proto

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

type CreateOrUpdateDeviceResponse_Error int32

const (
	CreateOrUpdateDeviceResponse_UNKNOWN_ERROR              CreateOrUpdateDeviceResponse_Error = 0
	CreateOrUpdateDeviceResponse_INVALID_DEVICE_TOKEN_ERROR CreateOrUpdateDeviceResponse_Error = 1
)

// Enum value maps for CreateOrUpdateDeviceResponse_Error.
var (
	CreateOrUpdateDeviceResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "INVALID_DEVICE_TOKEN_ERROR",
	}
	CreateOrUpdateDeviceResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":              0,
		"INVALID_DEVICE_TOKEN_ERROR": 1,
	}
)

func (x CreateOrUpdateDeviceResponse_Error) Enum() *CreateOrUpdateDeviceResponse_Error {
	p := new(CreateOrUpdateDeviceResponse_Error)
	*p = x
	return p
}

func (x CreateOrUpdateDeviceResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateOrUpdateDeviceResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_devices_proto_enumTypes[0].Descriptor()
}

func (CreateOrUpdateDeviceResponse_Error) Type() protoreflect.EnumType {
	return &file_devices_proto_enumTypes[0]
}

func (x CreateOrUpdateDeviceResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateOrUpdateDeviceResponse_Error.Descriptor instead.
func (CreateOrUpdateDeviceResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{2, 0}
}

type DeleteDeviceResponse_Error int32

const (
	DeleteDeviceResponse_UNKNOWN_ERROR          DeleteDeviceResponse_Error = 0
	DeleteDeviceResponse_DEVICE_NOT_FOUND_ERROR DeleteDeviceResponse_Error = 1
)

// Enum value maps for DeleteDeviceResponse_Error.
var (
	DeleteDeviceResponse_Error_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "DEVICE_NOT_FOUND_ERROR",
	}
	DeleteDeviceResponse_Error_value = map[string]int32{
		"UNKNOWN_ERROR":          0,
		"DEVICE_NOT_FOUND_ERROR": 1,
	}
)

func (x DeleteDeviceResponse_Error) Enum() *DeleteDeviceResponse_Error {
	p := new(DeleteDeviceResponse_Error)
	*p = x
	return p
}

func (x DeleteDeviceResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteDeviceResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_devices_proto_enumTypes[1].Descriptor()
}

func (DeleteDeviceResponse_Error) Type() protoreflect.EnumType {
	return &file_devices_proto_enumTypes[1]
}

func (x DeleteDeviceResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteDeviceResponse_Error.Descriptor instead.
func (DeleteDeviceResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{6, 0}
}

// Device for push notifications
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"` // device token from Firebase / APNs
	Useragent               string `protobuf:"bytes,3,opt,name=useragent,proto3" json:"useragent,omitempty"`
	CreatedAt               uint64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DisablePushesWhenOnline bool   `protobuf:"varint,5,opt,name=disable_pushes_when_online,json=disablePushesWhenOnline,proto3" json:"disable_pushes_when_online,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Device) GetUseragent() string {
	if x != nil {
		return x.Useragent
	}
	return ""
}

func (x *Device) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Device) GetDisablePushesWhenOnline() bool {
	if x != nil {
		return x.DisablePushesWhenOnline
	}
	return false
}

type CreateOrUpdateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token                   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TestPushValue           string `protobuf:"bytes,2,opt,name=test_push_value,json=testPushValue,proto3" json:"test_push_value,omitempty"` // send data push with payload: {"type": "test", "value": <given value>"}
	DisablePushesWhenOnline bool   `protobuf:"varint,5,opt,name=disable_pushes_when_online,json=disablePushesWhenOnline,proto3" json:"disable_pushes_when_online,omitempty"`
}

func (x *CreateOrUpdateDeviceRequest) Reset() {
	*x = CreateOrUpdateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateDeviceRequest) ProtoMessage() {}

func (x *CreateOrUpdateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrUpdateDeviceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateOrUpdateDeviceRequest) GetTestPushValue() string {
	if x != nil {
		return x.TestPushValue
	}
	return ""
}

func (x *CreateOrUpdateDeviceRequest) GetDisablePushesWhenOnline() bool {
	if x != nil {
		return x.DisablePushesWhenOnline
	}
	return false
}

type CreateOrUpdateDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []CreateOrUpdateDeviceResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.CreateOrUpdateDeviceResponse_Error" json:"errors,omitempty"`
	Device *Device                              `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateOrUpdateDeviceResponse) Reset() {
	*x = CreateOrUpdateDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateDeviceResponse) ProtoMessage() {}

func (x *CreateOrUpdateDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateDeviceResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateDeviceResponse) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrUpdateDeviceResponse) GetErrors() []CreateOrUpdateDeviceResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *CreateOrUpdateDeviceResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type DevicesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DevicesListRequest) Reset() {
	*x = DevicesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicesListRequest) ProtoMessage() {}

func (x *DevicesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicesListRequest.ProtoReflect.Descriptor instead.
func (*DevicesListRequest) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{3}
}

type DevicesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *DevicesListResponse) Reset() {
	*x = DevicesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicesListResponse) ProtoMessage() {}

func (x *DevicesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicesListResponse.ProtoReflect.Descriptor instead.
func (*DevicesListResponse) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{4}
}

func (x *DevicesListResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type DeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDeviceRequest) Reset() {
	*x = DeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceRequest) ProtoMessage() {}

func (x *DeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []DeleteDeviceResponse_Error `protobuf:"varint,1,rep,packed,name=errors,proto3,enum=teaproto.DeleteDeviceResponse_Error" json:"errors,omitempty"`
}

func (x *DeleteDeviceResponse) Reset() {
	*x = DeleteDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devices_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceResponse) ProtoMessage() {}

func (x *DeleteDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devices_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeviceResponse) Descriptor() ([]byte, []int) {
	return file_devices_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDeviceResponse) GetErrors() []DeleteDeviceResponse_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_devices_proto protoreflect.FileDescriptor

var file_devices_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x06, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x65, 0x73, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x75, 0x73, 0x68, 0x65, 0x73, 0x57, 0x68, 0x65, 0x6e, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x75,
	0x73, 0x68, 0x65, 0x73, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x75, 0x73, 0x68, 0x65, 0x73, 0x57, 0x68, 0x65, 0x6e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22,
	0xca, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x3a, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x61,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8c, 0x01, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0x36, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_devices_proto_rawDescOnce sync.Once
	file_devices_proto_rawDescData = file_devices_proto_rawDesc
)

func file_devices_proto_rawDescGZIP() []byte {
	file_devices_proto_rawDescOnce.Do(func() {
		file_devices_proto_rawDescData = protoimpl.X.CompressGZIP(file_devices_proto_rawDescData)
	})
	return file_devices_proto_rawDescData
}

var file_devices_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_devices_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_devices_proto_goTypes = []interface{}{
	(CreateOrUpdateDeviceResponse_Error)(0), // 0: teaproto.CreateOrUpdateDeviceResponse.Error
	(DeleteDeviceResponse_Error)(0),         // 1: teaproto.DeleteDeviceResponse.Error
	(*Device)(nil),                          // 2: teaproto.Device
	(*CreateOrUpdateDeviceRequest)(nil),     // 3: teaproto.CreateOrUpdateDeviceRequest
	(*CreateOrUpdateDeviceResponse)(nil),    // 4: teaproto.CreateOrUpdateDeviceResponse
	(*DevicesListRequest)(nil),              // 5: teaproto.DevicesListRequest
	(*DevicesListResponse)(nil),             // 6: teaproto.DevicesListResponse
	(*DeleteDeviceRequest)(nil),             // 7: teaproto.DeleteDeviceRequest
	(*DeleteDeviceResponse)(nil),            // 8: teaproto.DeleteDeviceResponse
}
var file_devices_proto_depIdxs = []int32{
	0, // 0: teaproto.CreateOrUpdateDeviceResponse.errors:type_name -> teaproto.CreateOrUpdateDeviceResponse.Error
	2, // 1: teaproto.CreateOrUpdateDeviceResponse.device:type_name -> teaproto.Device
	2, // 2: teaproto.DevicesListResponse.devices:type_name -> teaproto.Device
	1, // 3: teaproto.DeleteDeviceResponse.errors:type_name -> teaproto.DeleteDeviceResponse.Error
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_devices_proto_init() }
func file_devices_proto_init() {
	if File_devices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_devices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateDeviceRequest); i {
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
		file_devices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateDeviceResponse); i {
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
		file_devices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicesListRequest); i {
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
		file_devices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicesListResponse); i {
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
		file_devices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceRequest); i {
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
		file_devices_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceResponse); i {
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
			RawDescriptor: file_devices_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_devices_proto_goTypes,
		DependencyIndexes: file_devices_proto_depIdxs,
		EnumInfos:         file_devices_proto_enumTypes,
		MessageInfos:      file_devices_proto_msgTypes,
	}.Build()
	File_devices_proto = out.File
	file_devices_proto_rawDesc = nil
	file_devices_proto_goTypes = nil
	file_devices_proto_depIdxs = nil
}