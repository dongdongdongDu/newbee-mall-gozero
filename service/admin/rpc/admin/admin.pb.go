// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: admin.proto

package admin

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

// 管理员token
type AdminToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminUserId int64  `protobuf:"varint,1,opt,name=adminUserId,proto3" json:"adminUserId,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	UpdateTime  int64  `protobuf:"varint,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	ExpireTime  int64  `protobuf:"varint,4,opt,name=expireTime,proto3" json:"expireTime,omitempty"`
}

func (x *AdminToken) Reset() {
	*x = AdminToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminToken) ProtoMessage() {}

func (x *AdminToken) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminToken.ProtoReflect.Descriptor instead.
func (*AdminToken) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AdminToken) GetAdminUserId() int64 {
	if x != nil {
		return x.AdminUserId
	}
	return 0
}

func (x *AdminToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AdminToken) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *AdminToken) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

// 管理员
type AdminModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminUserId   int64  `protobuf:"varint,1,opt,name=adminUserId,proto3" json:"adminUserId,omitempty"`
	LoginUserName string `protobuf:"bytes,2,opt,name=loginUserName,proto3" json:"loginUserName,omitempty"`
	LoginPassword string `protobuf:"bytes,3,opt,name=loginPassword,proto3" json:"loginPassword,omitempty"`
	NickName      string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Locked        int64  `protobuf:"varint,5,opt,name=locked,proto3" json:"locked,omitempty"`
}

func (x *AdminModel) Reset() {
	*x = AdminModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminModel) ProtoMessage() {}

func (x *AdminModel) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminModel.ProtoReflect.Descriptor instead.
func (*AdminModel) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminModel) GetAdminUserId() int64 {
	if x != nil {
		return x.AdminUserId
	}
	return 0
}

func (x *AdminModel) GetLoginUserName() string {
	if x != nil {
		return x.LoginUserName
	}
	return ""
}

func (x *AdminModel) GetLoginPassword() string {
	if x != nil {
		return x.LoginPassword
	}
	return ""
}

func (x *AdminModel) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *AdminModel) GetLocked() int64 {
	if x != nil {
		return x.Locked
	}
	return 0
}

// 管理员登录
type AdminLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName    string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	PasswordMd5 string `protobuf:"bytes,2,opt,name=passwordMd5,proto3" json:"passwordMd5,omitempty"`
}

func (x *AdminLoginRequest) Reset() {
	*x = AdminLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginRequest) ProtoMessage() {}

func (x *AdminLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginRequest.ProtoReflect.Descriptor instead.
func (*AdminLoginRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdminLoginRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AdminLoginRequest) GetPasswordMd5() string {
	if x != nil {
		return x.PasswordMd5
	}
	return ""
}

type AdminLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminToken *AdminToken `protobuf:"bytes,1,opt,name=adminToken,proto3" json:"adminToken,omitempty"`
}

func (x *AdminLoginResponse) Reset() {
	*x = AdminLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginResponse) ProtoMessage() {}

func (x *AdminLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginResponse.ProtoReflect.Descriptor instead.
func (*AdminLoginResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AdminLoginResponse) GetAdminToken() *AdminToken {
	if x != nil {
		return x.AdminToken
	}
	return nil
}

// 管理员更改昵称
type UpdateAdminNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	LoginUserName string `protobuf:"bytes,2,opt,name=loginUserName,proto3" json:"loginUserName,omitempty"`
	NickName      string `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
}

func (x *UpdateAdminNameRequest) Reset() {
	*x = UpdateAdminNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminNameRequest) ProtoMessage() {}

func (x *UpdateAdminNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminNameRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminNameRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAdminNameRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateAdminNameRequest) GetLoginUserName() string {
	if x != nil {
		return x.LoginUserName
	}
	return ""
}

func (x *UpdateAdminNameRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

// 管理员更改密码
type UpdateAdminPwdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token            string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OriginalPassword string `protobuf:"bytes,2,opt,name=originalPassword,proto3" json:"originalPassword,omitempty"`
	NewPassword      string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *UpdateAdminPwdRequest) Reset() {
	*x = UpdateAdminPwdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminPwdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminPwdRequest) ProtoMessage() {}

func (x *UpdateAdminPwdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminPwdRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminPwdRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAdminPwdRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateAdminPwdRequest) GetOriginalPassword() string {
	if x != nil {
		return x.OriginalPassword
	}
	return ""
}

func (x *UpdateAdminPwdRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

// 查看管理员信息
type GetAdminProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAdminProfileRequest) Reset() {
	*x = GetAdminProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminProfileRequest) ProtoMessage() {}

func (x *GetAdminProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminProfileRequest.ProtoReflect.Descriptor instead.
func (*GetAdminProfileRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{6}
}

func (x *GetAdminProfileRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetAdminProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *AdminModel `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *GetAdminProfileResponse) Reset() {
	*x = GetAdminProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminProfileResponse) ProtoMessage() {}

func (x *GetAdminProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminProfileResponse.ProtoReflect.Descriptor instead.
func (*GetAdminProfileResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{7}
}

func (x *GetAdminProfileResponse) GetAdmin() *AdminModel {
	if x != nil {
		return x.Admin
	}
	return nil
}

// 管理员退出登录
type AdminLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminToken string `protobuf:"bytes,1,opt,name=adminToken,proto3" json:"adminToken,omitempty"`
}

func (x *AdminLogoutRequest) Reset() {
	*x = AdminLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLogoutRequest) ProtoMessage() {}

func (x *AdminLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLogoutRequest.ProtoReflect.Descriptor instead.
func (*AdminLogoutRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{8}
}

func (x *AdminLogoutRequest) GetAdminToken() string {
	if x != nil {
		return x.AdminToken
	}
	return ""
}

// 空响应
type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{9}
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x0a,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x11,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x64, 0x35, 0x22,
	0x47, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x34, 0x0a, 0x12, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xea, 0x02, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x41, 0x0a, 0x0a,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50,
	0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x19,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_admin_proto_goTypes = []interface{}{
	(*AdminToken)(nil),              // 0: admin.AdminToken
	(*AdminModel)(nil),              // 1: admin.AdminModel
	(*AdminLoginRequest)(nil),       // 2: admin.AdminLoginRequest
	(*AdminLoginResponse)(nil),      // 3: admin.AdminLoginResponse
	(*UpdateAdminNameRequest)(nil),  // 4: admin.UpdateAdminNameRequest
	(*UpdateAdminPwdRequest)(nil),   // 5: admin.UpdateAdminPwdRequest
	(*GetAdminProfileRequest)(nil),  // 6: admin.GetAdminProfileRequest
	(*GetAdminProfileResponse)(nil), // 7: admin.GetAdminProfileResponse
	(*AdminLogoutRequest)(nil),      // 8: admin.AdminLogoutRequest
	(*EmptyResponse)(nil),           // 9: admin.EmptyResponse
}
var file_admin_proto_depIdxs = []int32{
	0, // 0: admin.AdminLoginResponse.adminToken:type_name -> admin.AdminToken
	1, // 1: admin.GetAdminProfileResponse.admin:type_name -> admin.AdminModel
	2, // 2: admin.admin.adminLogin:input_type -> admin.AdminLoginRequest
	6, // 3: admin.admin.getAdminProfile:input_type -> admin.GetAdminProfileRequest
	4, // 4: admin.admin.updateAdminName:input_type -> admin.UpdateAdminNameRequest
	5, // 5: admin.admin.updateAdminPwd:input_type -> admin.UpdateAdminPwdRequest
	8, // 6: admin.admin.adminLogout:input_type -> admin.AdminLogoutRequest
	3, // 7: admin.admin.adminLogin:output_type -> admin.AdminLoginResponse
	7, // 8: admin.admin.getAdminProfile:output_type -> admin.GetAdminProfileResponse
	9, // 9: admin.admin.updateAdminName:output_type -> admin.EmptyResponse
	9, // 10: admin.admin.updateAdminPwd:output_type -> admin.EmptyResponse
	9, // 11: admin.admin.adminLogout:output_type -> admin.EmptyResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminToken); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminModel); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginRequest); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLoginResponse); i {
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
		file_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminNameRequest); i {
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
		file_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminPwdRequest); i {
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
		file_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminProfileRequest); i {
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
		file_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminProfileResponse); i {
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
		file_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminLogoutRequest); i {
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
		file_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
