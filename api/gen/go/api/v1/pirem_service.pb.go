// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/v1/pirem_service.proto

package v1

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

// ---------- Send IR ----------
type SendIrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string  `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	IrData   *IrData `protobuf:"bytes,2,opt,name=ir_data,json=irData,proto3" json:"ir_data,omitempty"`
}

func (x *SendIrRequest) Reset() {
	*x = SendIrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendIrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendIrRequest) ProtoMessage() {}

func (x *SendIrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendIrRequest.ProtoReflect.Descriptor instead.
func (*SendIrRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendIrRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *SendIrRequest) GetIrData() *IrData {
	if x != nil {
		return x.IrData
	}
	return nil
}

type SendIrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendIrResponse) Reset() {
	*x = SendIrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendIrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendIrResponse) ProtoMessage() {}

func (x *SendIrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendIrResponse.ProtoReflect.Descriptor instead.
func (*SendIrResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{1}
}

// ---------- Receive IR ----------
type ReceiveIrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *ReceiveIrRequest) Reset() {
	*x = ReceiveIrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveIrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveIrRequest) ProtoMessage() {}

func (x *ReceiveIrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveIrRequest.ProtoReflect.Descriptor instead.
func (*ReceiveIrRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReceiveIrRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type ReceiveIrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IrData *IrData `protobuf:"bytes,1,opt,name=ir_data,json=irData,proto3" json:"ir_data,omitempty"`
}

func (x *ReceiveIrResponse) Reset() {
	*x = ReceiveIrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveIrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveIrResponse) ProtoMessage() {}

func (x *ReceiveIrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveIrResponse.ProtoReflect.Descriptor instead.
func (*ReceiveIrResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveIrResponse) GetIrData() *IrData {
	if x != nil {
		return x.IrData
	}
	return nil
}

// ---------- Get all device information -----------
type ListDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDevicesRequest) Reset() {
	*x = ListDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDevicesRequest) ProtoMessage() {}

func (x *ListDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDevicesRequest.ProtoReflect.Descriptor instead.
func (*ListDevicesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{4}
}

type ListDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *ListDevicesResponse) Reset() {
	*x = ListDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDevicesResponse) ProtoMessage() {}

func (x *ListDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDevicesResponse.ProtoReflect.Descriptor instead.
func (*ListDevicesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListDevicesResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

// ----------- Get a device information ----------
type GetDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetDeviceRequest) Reset() {
	*x = GetDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceRequest) ProtoMessage() {}

func (x *GetDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetDeviceRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type ListRemotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRemotesRequest) Reset() {
	*x = ListRemotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRemotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRemotesRequest) ProtoMessage() {}

func (x *ListRemotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRemotesRequest.ProtoReflect.Descriptor instead.
func (*ListRemotesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{7}
}

type ListRemotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remotes []*Remote `protobuf:"bytes,1,rep,name=remotes,proto3" json:"remotes,omitempty"`
}

func (x *ListRemotesResponse) Reset() {
	*x = ListRemotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRemotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRemotesResponse) ProtoMessage() {}

func (x *ListRemotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRemotesResponse.ProtoReflect.Descriptor instead.
func (*ListRemotesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListRemotesResponse) GetRemotes() []*Remote {
	if x != nil {
		return x.Remotes
	}
	return nil
}

type GetRemoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteId string `protobuf:"bytes,1,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
}

func (x *GetRemoteRequest) Reset() {
	*x = GetRemoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteRequest) ProtoMessage() {}

func (x *GetRemoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteRequest.ProtoReflect.Descriptor instead.
func (*GetRemoteRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetRemoteRequest) GetRemoteId() string {
	if x != nil {
		return x.RemoteId
	}
	return ""
}

type UpdateRemoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteId string `protobuf:"bytes,1,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *UpdateRemoteRequest) Reset() {
	*x = UpdateRemoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRemoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRemoteRequest) ProtoMessage() {}

func (x *UpdateRemoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRemoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateRemoteRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRemoteRequest) GetRemoteId() string {
	if x != nil {
		return x.RemoteId
	}
	return ""
}

func (x *UpdateRemoteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRemoteRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type DeleteRemoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteId string `protobuf:"bytes,1,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
}

func (x *DeleteRemoteRequest) Reset() {
	*x = DeleteRemoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRemoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRemoteRequest) ProtoMessage() {}

func (x *DeleteRemoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRemoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRemoteRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteRemoteRequest) GetRemoteId() string {
	if x != nil {
		return x.RemoteId
	}
	return ""
}

type GetButtonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButtonId string `protobuf:"bytes,1,opt,name=button_id,json=buttonId,proto3" json:"button_id,omitempty"`
}

func (x *GetButtonRequest) Reset() {
	*x = GetButtonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetButtonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetButtonRequest) ProtoMessage() {}

func (x *GetButtonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetButtonRequest.ProtoReflect.Descriptor instead.
func (*GetButtonRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{12}
}

func (x *GetButtonRequest) GetButtonId() string {
	if x != nil {
		return x.ButtonId
	}
	return ""
}

type LearnIrDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButtonId string  `protobuf:"bytes,1,opt,name=button_id,json=buttonId,proto3" json:"button_id,omitempty"`
	IrData   *IrData `protobuf:"bytes,2,opt,name=ir_data,json=irData,proto3" json:"ir_data,omitempty"`
}

func (x *LearnIrDataRequest) Reset() {
	*x = LearnIrDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LearnIrDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LearnIrDataRequest) ProtoMessage() {}

func (x *LearnIrDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LearnIrDataRequest.ProtoReflect.Descriptor instead.
func (*LearnIrDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{13}
}

func (x *LearnIrDataRequest) GetButtonId() string {
	if x != nil {
		return x.ButtonId
	}
	return ""
}

func (x *LearnIrDataRequest) GetIrData() *IrData {
	if x != nil {
		return x.IrData
	}
	return nil
}

type PushButtonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButtonId string `protobuf:"bytes,1,opt,name=button_id,json=buttonId,proto3" json:"button_id,omitempty"`
}

func (x *PushButtonRequest) Reset() {
	*x = PushButtonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_pirem_service_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushButtonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushButtonRequest) ProtoMessage() {}

func (x *PushButtonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_pirem_service_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushButtonRequest.ProtoReflect.Descriptor instead.
func (*PushButtonRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_pirem_service_proto_rawDescGZIP(), []int{14}
}

func (x *PushButtonRequest) GetButtonId() string {
	if x != nil {
		return x.ButtonId
	}
	return ""
}

var File_api_v1_pirem_service_proto protoreflect.FileDescriptor

var file_api_v1_pirem_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x69,
	0x72, 0x65, 0x6d, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0d, 0x53,
	0x65, 0x6e, 0x64, 0x49, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x69, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x69, 0x72,
	0x65, 0x6d, 0x2e, 0x49, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x69, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x69, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x69, 0x72,
	0x65, 0x6d, 0x2e, 0x49, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x69, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x2f,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x63, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x12, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x49, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07,
	0x69, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x49, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x69, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xcd, 0x05, 0x0a, 0x0c, 0x50, 0x69, 0x52, 0x65, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x49,
	0x72, 0x12, 0x14, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x49, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x72, 0x12, 0x17, 0x2e,
	0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x49,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x70,
	0x69, 0x72, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x69,
	0x72, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x69, 0x72,
	0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x49, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x4c, 0x65, 0x61, 0x72,
	0x6e, 0x49, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70,
	0x69, 0x72, 0x65, 0x6d, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x69, 0x72, 0x65, 0x6d, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x61, 0x4b, 0x61, 0x32, 0x33, 0x35, 0x35, 0x2f, 0x69, 0x72,
	0x64, 0x65, 0x63, 0x6b, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_pirem_service_proto_rawDescOnce sync.Once
	file_api_v1_pirem_service_proto_rawDescData = file_api_v1_pirem_service_proto_rawDesc
)

func file_api_v1_pirem_service_proto_rawDescGZIP() []byte {
	file_api_v1_pirem_service_proto_rawDescOnce.Do(func() {
		file_api_v1_pirem_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_pirem_service_proto_rawDescData)
	})
	return file_api_v1_pirem_service_proto_rawDescData
}

var file_api_v1_pirem_service_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_api_v1_pirem_service_proto_goTypes = []interface{}{
	(*SendIrRequest)(nil),       // 0: pirem.SendIrRequest
	(*SendIrResponse)(nil),      // 1: pirem.SendIrResponse
	(*ReceiveIrRequest)(nil),    // 2: pirem.ReceiveIrRequest
	(*ReceiveIrResponse)(nil),   // 3: pirem.ReceiveIrResponse
	(*ListDevicesRequest)(nil),  // 4: pirem.ListDevicesRequest
	(*ListDevicesResponse)(nil), // 5: pirem.ListDevicesResponse
	(*GetDeviceRequest)(nil),    // 6: pirem.GetDeviceRequest
	(*ListRemotesRequest)(nil),  // 7: pirem.ListRemotesRequest
	(*ListRemotesResponse)(nil), // 8: pirem.ListRemotesResponse
	(*GetRemoteRequest)(nil),    // 9: pirem.GetRemoteRequest
	(*UpdateRemoteRequest)(nil), // 10: pirem.UpdateRemoteRequest
	(*DeleteRemoteRequest)(nil), // 11: pirem.DeleteRemoteRequest
	(*GetButtonRequest)(nil),    // 12: pirem.GetButtonRequest
	(*LearnIrDataRequest)(nil),  // 13: pirem.LearnIrDataRequest
	(*PushButtonRequest)(nil),   // 14: pirem.PushButtonRequest
	(*IrData)(nil),              // 15: pirem.IrData
	(*Device)(nil),              // 16: pirem.Device
	(*Remote)(nil),              // 17: pirem.Remote
	(*Empty)(nil),               // 18: pirem.Empty
	(*Button)(nil),              // 19: pirem.Button
}
var file_api_v1_pirem_service_proto_depIdxs = []int32{
	15, // 0: pirem.SendIrRequest.ir_data:type_name -> pirem.IrData
	15, // 1: pirem.ReceiveIrResponse.ir_data:type_name -> pirem.IrData
	16, // 2: pirem.ListDevicesResponse.devices:type_name -> pirem.Device
	17, // 3: pirem.ListRemotesResponse.remotes:type_name -> pirem.Remote
	15, // 4: pirem.LearnIrDataRequest.ir_data:type_name -> pirem.IrData
	0,  // 5: pirem.PiRemService.SendIr:input_type -> pirem.SendIrRequest
	2,  // 6: pirem.PiRemService.ReceiveIr:input_type -> pirem.ReceiveIrRequest
	4,  // 7: pirem.PiRemService.ListDevices:input_type -> pirem.ListDevicesRequest
	6,  // 8: pirem.PiRemService.GetDevice:input_type -> pirem.GetDeviceRequest
	17, // 9: pirem.PiRemService.CreateRemote:input_type -> pirem.Remote
	7,  // 10: pirem.PiRemService.ListRemotes:input_type -> pirem.ListRemotesRequest
	9,  // 11: pirem.PiRemService.GetRemote:input_type -> pirem.GetRemoteRequest
	10, // 12: pirem.PiRemService.UpdateRemote:input_type -> pirem.UpdateRemoteRequest
	11, // 13: pirem.PiRemService.DeleteRemote:input_type -> pirem.DeleteRemoteRequest
	12, // 14: pirem.PiRemService.GetButton:input_type -> pirem.GetButtonRequest
	13, // 15: pirem.PiRemService.LearnIrData:input_type -> pirem.LearnIrDataRequest
	14, // 16: pirem.PiRemService.PushButton:input_type -> pirem.PushButtonRequest
	1,  // 17: pirem.PiRemService.SendIr:output_type -> pirem.SendIrResponse
	15, // 18: pirem.PiRemService.ReceiveIr:output_type -> pirem.IrData
	5,  // 19: pirem.PiRemService.ListDevices:output_type -> pirem.ListDevicesResponse
	16, // 20: pirem.PiRemService.GetDevice:output_type -> pirem.Device
	17, // 21: pirem.PiRemService.CreateRemote:output_type -> pirem.Remote
	8,  // 22: pirem.PiRemService.ListRemotes:output_type -> pirem.ListRemotesResponse
	17, // 23: pirem.PiRemService.GetRemote:output_type -> pirem.Remote
	18, // 24: pirem.PiRemService.UpdateRemote:output_type -> pirem.Empty
	18, // 25: pirem.PiRemService.DeleteRemote:output_type -> pirem.Empty
	19, // 26: pirem.PiRemService.GetButton:output_type -> pirem.Button
	18, // 27: pirem.PiRemService.LearnIrData:output_type -> pirem.Empty
	18, // 28: pirem.PiRemService.PushButton:output_type -> pirem.Empty
	17, // [17:29] is the sub-list for method output_type
	5,  // [5:17] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_pirem_service_proto_init() }
func file_api_v1_pirem_service_proto_init() {
	if File_api_v1_pirem_service_proto != nil {
		return
	}
	file_api_v1_device_proto_init()
	file_api_v1_irdata_proto_init()
	file_api_v1_remote_proto_init()
	file_api_v1_button_proto_init()
	file_api_v1_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_pirem_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendIrRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendIrResponse); i {
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
		file_api_v1_pirem_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveIrRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveIrResponse); i {
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
		file_api_v1_pirem_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDevicesRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDevicesResponse); i {
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
		file_api_v1_pirem_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRemotesRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRemotesResponse); i {
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
		file_api_v1_pirem_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRemoteRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRemoteRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetButtonRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LearnIrDataRequest); i {
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
		file_api_v1_pirem_service_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushButtonRequest); i {
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
			RawDescriptor: file_api_v1_pirem_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_pirem_service_proto_goTypes,
		DependencyIndexes: file_api_v1_pirem_service_proto_depIdxs,
		MessageInfos:      file_api_v1_pirem_service_proto_msgTypes,
	}.Build()
	File_api_v1_pirem_service_proto = out.File
	file_api_v1_pirem_service_proto_rawDesc = nil
	file_api_v1_pirem_service_proto_goTypes = nil
	file_api_v1_pirem_service_proto_depIdxs = nil
}
