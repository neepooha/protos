// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: sso/permissions.proto

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

type SetAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SetAdminRequest) Reset() {
	*x = SetAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAdminRequest) ProtoMessage() {}

func (x *SetAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAdminRequest.ProtoReflect.Descriptor instead.
func (*SetAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *SetAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SetAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SetAdmin bool `protobuf:"varint,1,opt,name=set_admin,json=setAdmin,proto3" json:"set_admin,omitempty"`
}

func (x *SetAdminResponse) Reset() {
	*x = SetAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAdminResponse) ProtoMessage() {}

func (x *SetAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAdminResponse.ProtoReflect.Descriptor instead.
func (*SetAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *SetAdminResponse) GetSetAdmin() bool {
	if x != nil {
		return x.SetAdmin
	}
	return false
}

type DelAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DelAdminRequest) Reset() {
	*x = DelAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelAdminRequest) ProtoMessage() {}

func (x *DelAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelAdminRequest.ProtoReflect.Descriptor instead.
func (*DelAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *DelAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DelAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelAdmin bool `protobuf:"varint,1,opt,name=del_admin,json=delAdmin,proto3" json:"del_admin,omitempty"`
}

func (x *DelAdminResponse) Reset() {
	*x = DelAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelAdminResponse) ProtoMessage() {}

func (x *DelAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelAdminResponse.ProtoReflect.Descriptor instead.
func (*DelAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *DelAdminResponse) GetDelAdmin() bool {
	if x != nil {
		return x.DelAdmin
	}
	return false
}

type IsAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IsAdminRequest) Reset() {
	*x = IsAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminRequest) ProtoMessage() {}

func (x *IsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminRequest.ProtoReflect.Descriptor instead.
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{4}
}

func (x *IsAdminRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *IsAdminResponse) Reset() {
	*x = IsAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminResponse) ProtoMessage() {}

func (x *IsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminResponse.ProtoReflect.Descriptor instead.
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{5}
}

func (x *IsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type IsCreatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IsCreatorRequest) Reset() {
	*x = IsCreatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsCreatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsCreatorRequest) ProtoMessage() {}

func (x *IsCreatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsCreatorRequest.ProtoReflect.Descriptor instead.
func (*IsCreatorRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{6}
}

func (x *IsCreatorRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsCreatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCreator bool `protobuf:"varint,1,opt,name=is_creator,json=isCreator,proto3" json:"is_creator,omitempty"`
}

func (x *IsCreatorResponse) Reset() {
	*x = IsCreatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsCreatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsCreatorResponse) ProtoMessage() {}

func (x *IsCreatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsCreatorResponse.ProtoReflect.Descriptor instead.
func (*IsCreatorResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{7}
}

func (x *IsCreatorResponse) GetIsCreator() bool {
	if x != nil {
		return x.IsCreator
	}
	return false
}

var File_sso_permissions_proto protoreflect.FileDescriptor

var file_sso_permissions_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x73, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x65, 0x72, 0x6d, 0x22, 0x27, 0x0a,
	0x0f, 0x53, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x2f, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x22, 0x29, 0x0a, 0x0e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f,
	0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x2b, 0x0a, 0x10, 0x49, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x49, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x32, 0xf9, 0x01, 0x0a, 0x0b,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x2e, 0x53,
	0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x2e, 0x44, 0x65, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x2e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x2e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x49, 0x73, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x2e, 0x49, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x2e, 0x49, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x6e, 0x65, 0x65, 0x70, 0x6f,
	0x6f, 0x68, 0x61, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x32, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_permissions_proto_rawDescOnce sync.Once
	file_sso_permissions_proto_rawDescData = file_sso_permissions_proto_rawDesc
)

func file_sso_permissions_proto_rawDescGZIP() []byte {
	file_sso_permissions_proto_rawDescOnce.Do(func() {
		file_sso_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_permissions_proto_rawDescData)
	})
	return file_sso_permissions_proto_rawDescData
}

var file_sso_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sso_permissions_proto_goTypes = []interface{}{
	(*SetAdminRequest)(nil),   // 0: perm.SetAdminRequest
	(*SetAdminResponse)(nil),  // 1: perm.SetAdminResponse
	(*DelAdminRequest)(nil),   // 2: perm.DelAdminRequest
	(*DelAdminResponse)(nil),  // 3: perm.DelAdminResponse
	(*IsAdminRequest)(nil),    // 4: perm.IsAdminRequest
	(*IsAdminResponse)(nil),   // 5: perm.IsAdminResponse
	(*IsCreatorRequest)(nil),  // 6: perm.IsCreatorRequest
	(*IsCreatorResponse)(nil), // 7: perm.IsCreatorResponse
}
var file_sso_permissions_proto_depIdxs = []int32{
	0, // 0: perm.Permissions.SetAdmin:input_type -> perm.SetAdminRequest
	2, // 1: perm.Permissions.DelAdmin:input_type -> perm.DelAdminRequest
	4, // 2: perm.Permissions.IsAdmin:input_type -> perm.IsAdminRequest
	6, // 3: perm.Permissions.IsCreator:input_type -> perm.IsCreatorRequest
	1, // 4: perm.Permissions.SetAdmin:output_type -> perm.SetAdminResponse
	3, // 5: perm.Permissions.DelAdmin:output_type -> perm.DelAdminResponse
	5, // 6: perm.Permissions.IsAdmin:output_type -> perm.IsAdminResponse
	7, // 7: perm.Permissions.IsCreator:output_type -> perm.IsCreatorResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_permissions_proto_init() }
func file_sso_permissions_proto_init() {
	if File_sso_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAdminRequest); i {
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
		file_sso_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAdminResponse); i {
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
		file_sso_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelAdminRequest); i {
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
		file_sso_permissions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelAdminResponse); i {
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
		file_sso_permissions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAdminRequest); i {
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
		file_sso_permissions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAdminResponse); i {
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
		file_sso_permissions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsCreatorRequest); i {
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
		file_sso_permissions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsCreatorResponse); i {
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
			RawDescriptor: file_sso_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_permissions_proto_goTypes,
		DependencyIndexes: file_sso_permissions_proto_depIdxs,
		MessageInfos:      file_sso_permissions_proto_msgTypes,
	}.Build()
	File_sso_permissions_proto = out.File
	file_sso_permissions_proto_rawDesc = nil
	file_sso_permissions_proto_goTypes = nil
	file_sso_permissions_proto_depIdxs = nil
}
