// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: sso/apps.proto

package ssov2

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

type GetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *GetAppRequest) Reset() {
	*x = GetAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRequest) ProtoMessage() {}

func (x *GetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRequest.ProtoReflect.Descriptor instead.
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{0}
}

func (x *GetAppRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId   int32  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppName string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{1}
}

func (x *GetAppResponse) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *GetAppResponse) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type SetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName   string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
}

func (x *SetAppRequest) Reset() {
	*x = SetAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAppRequest) ProtoMessage() {}

func (x *SetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAppRequest.ProtoReflect.Descriptor instead.
func (*SetAppRequest) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{2}
}

func (x *SetAppRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *SetAppRequest) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

type SetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID int32 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *SetAppResponse) Reset() {
	*x = SetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAppResponse) ProtoMessage() {}

func (x *SetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAppResponse.ProtoReflect.Descriptor instead.
func (*SetAppResponse) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{3}
}

func (x *SetAppResponse) GetAppID() int32 {
	if x != nil {
		return x.AppID
	}
	return 0
}

type UpdAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID        int32  `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	NewAppName   string `protobuf:"bytes,2,opt,name=new_app_name,json=newAppName,proto3" json:"new_app_name,omitempty"`
	NewAppSecret string `protobuf:"bytes,3,opt,name=new_app_secret,json=newAppSecret,proto3" json:"new_app_secret,omitempty"`
}

func (x *UpdAppRequest) Reset() {
	*x = UpdAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdAppRequest) ProtoMessage() {}

func (x *UpdAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdAppRequest.ProtoReflect.Descriptor instead.
func (*UpdAppRequest) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{4}
}

func (x *UpdAppRequest) GetAppID() int32 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *UpdAppRequest) GetNewAppName() string {
	if x != nil {
		return x.NewAppName
	}
	return ""
}

func (x *UpdAppRequest) GetNewAppSecret() string {
	if x != nil {
		return x.NewAppSecret
	}
	return ""
}

type UpdAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUpdApp bool `protobuf:"varint,1,opt,name=is_upd_app,json=isUpdApp,proto3" json:"is_upd_app,omitempty"`
}

func (x *UpdAppResponse) Reset() {
	*x = UpdAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdAppResponse) ProtoMessage() {}

func (x *UpdAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdAppResponse.ProtoReflect.Descriptor instead.
func (*UpdAppResponse) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{5}
}

func (x *UpdAppResponse) GetIsUpdApp() bool {
	if x != nil {
		return x.IsUpdApp
	}
	return false
}

type DelAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *DelAppRequest) Reset() {
	*x = DelAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelAppRequest) ProtoMessage() {}

func (x *DelAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelAppRequest.ProtoReflect.Descriptor instead.
func (*DelAppRequest) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{6}
}

func (x *DelAppRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type DelAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDelApp bool `protobuf:"varint,1,opt,name=is_del_app,json=isDelApp,proto3" json:"is_del_app,omitempty"`
}

func (x *DelAppResponse) Reset() {
	*x = DelAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelAppResponse) ProtoMessage() {}

func (x *DelAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelAppResponse.ProtoReflect.Descriptor instead.
func (*DelAppResponse) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{7}
}

func (x *DelAppResponse) GetIsDelApp() bool {
	if x != nil {
		return x.IsDelApp
	}
	return false
}

var File_sso_apps_proto protoreflect.FileDescriptor

var file_sso_apps_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x73, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x2a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x22, 0x6d, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x41, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x41,
	0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x75, 0x70, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x55, 0x70, 0x64, 0x41, 0x70, 0x70, 0x22, 0x2a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x5f, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65,
	0x6c, 0x41, 0x70, 0x70, 0x32, 0xdc, 0x01, 0x0a, 0x04, 0x41, 0x70, 0x70, 0x73, 0x12, 0x35, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x13,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x41, 0x70, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x41, 0x70, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x6e, 0x65, 0x65, 0x70, 0x6f, 0x6f, 0x68, 0x61, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x76, 0x32, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_apps_proto_rawDescOnce sync.Once
	file_sso_apps_proto_rawDescData = file_sso_apps_proto_rawDesc
)

func file_sso_apps_proto_rawDescGZIP() []byte {
	file_sso_apps_proto_rawDescOnce.Do(func() {
		file_sso_apps_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_apps_proto_rawDescData)
	})
	return file_sso_apps_proto_rawDescData
}

var file_sso_apps_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sso_apps_proto_goTypes = []interface{}{
	(*GetAppRequest)(nil),  // 0: apps.GetAppRequest
	(*GetAppResponse)(nil), // 1: apps.GetAppResponse
	(*SetAppRequest)(nil),  // 2: apps.SetAppRequest
	(*SetAppResponse)(nil), // 3: apps.SetAppResponse
	(*UpdAppRequest)(nil),  // 4: apps.UpdAppRequest
	(*UpdAppResponse)(nil), // 5: apps.UpdAppResponse
	(*DelAppRequest)(nil),  // 6: apps.DelAppRequest
	(*DelAppResponse)(nil), // 7: apps.DelAppResponse
}
var file_sso_apps_proto_depIdxs = []int32{
	0, // 0: apps.Apps.GetAppID:input_type -> apps.GetAppRequest
	2, // 1: apps.Apps.SetApp:input_type -> apps.SetAppRequest
	4, // 2: apps.Apps.UpdApp:input_type -> apps.UpdAppRequest
	6, // 3: apps.Apps.DelApp:input_type -> apps.DelAppRequest
	1, // 4: apps.Apps.GetAppID:output_type -> apps.GetAppResponse
	3, // 5: apps.Apps.SetApp:output_type -> apps.SetAppResponse
	5, // 6: apps.Apps.UpdApp:output_type -> apps.UpdAppResponse
	7, // 7: apps.Apps.DelApp:output_type -> apps.DelAppResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_apps_proto_init() }
func file_sso_apps_proto_init() {
	if File_sso_apps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_apps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRequest); i {
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
		file_sso_apps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppResponse); i {
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
		file_sso_apps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAppRequest); i {
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
		file_sso_apps_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAppResponse); i {
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
		file_sso_apps_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdAppRequest); i {
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
		file_sso_apps_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdAppResponse); i {
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
		file_sso_apps_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelAppRequest); i {
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
		file_sso_apps_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelAppResponse); i {
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
			RawDescriptor: file_sso_apps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_apps_proto_goTypes,
		DependencyIndexes: file_sso_apps_proto_depIdxs,
		MessageInfos:      file_sso_apps_proto_msgTypes,
	}.Build()
	File_sso_apps_proto = out.File
	file_sso_apps_proto_rawDesc = nil
	file_sso_apps_proto_goTypes = nil
	file_sso_apps_proto_depIdxs = nil
}
