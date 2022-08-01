// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: admintoken.proto

package admintoken

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

// adminToken
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
		mi := &file_admintoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminToken) ProtoMessage() {}

func (x *AdminToken) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[0]
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
	return file_admintoken_proto_rawDescGZIP(), []int{0}
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

// 生成adminToken
type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminUserId int64 `protobuf:"varint,1,opt,name=adminUserId,proto3" json:"adminUserId,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admintoken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_admintoken_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenRequest) GetAdminUserId() int64 {
	if x != nil {
		return x.AdminUserId
	}
	return 0
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminToken *AdminToken `protobuf:"bytes,1,opt,name=adminToken,proto3" json:"adminToken,omitempty"`
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admintoken_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_admintoken_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateTokenResponse) GetAdminToken() *AdminToken {
	if x != nil {
		return x.AdminToken
	}
	return nil
}

// 获取已有adminToken
type GetExistTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetExistTokenRequest) Reset() {
	*x = GetExistTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admintoken_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExistTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExistTokenRequest) ProtoMessage() {}

func (x *GetExistTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExistTokenRequest.ProtoReflect.Descriptor instead.
func (*GetExistTokenRequest) Descriptor() ([]byte, []int) {
	return file_admintoken_proto_rawDescGZIP(), []int{3}
}

func (x *GetExistTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetExistTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminToken *AdminToken `protobuf:"bytes,1,opt,name=adminToken,proto3" json:"adminToken,omitempty"`
}

func (x *GetExistTokenResponse) Reset() {
	*x = GetExistTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admintoken_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExistTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExistTokenResponse) ProtoMessage() {}

func (x *GetExistTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExistTokenResponse.ProtoReflect.Descriptor instead.
func (*GetExistTokenResponse) Descriptor() ([]byte, []int) {
	return file_admintoken_proto_rawDescGZIP(), []int{4}
}

func (x *GetExistTokenResponse) GetAdminToken() *AdminToken {
	if x != nil {
		return x.AdminToken
	}
	return nil
}

// 删除adminToken
type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admintoken_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[5]
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
	return file_admintoken_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTokenResponse) Reset() {
	*x = DeleteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admintoken_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenResponse) ProtoMessage() {}

func (x *DeleteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admintoken_proto_msgTypes[6]
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
	return file_admintoken_proto_rawDescGZIP(), []int{6}
}

var File_admintoken_proto protoreflect.FileDescriptor

var file_admintoken_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84,
	0x01, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x4f, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4f,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x2a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x88, 0x02, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x54, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admintoken_proto_rawDescOnce sync.Once
	file_admintoken_proto_rawDescData = file_admintoken_proto_rawDesc
)

func file_admintoken_proto_rawDescGZIP() []byte {
	file_admintoken_proto_rawDescOnce.Do(func() {
		file_admintoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_admintoken_proto_rawDescData)
	})
	return file_admintoken_proto_rawDescData
}

var file_admintoken_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_admintoken_proto_goTypes = []interface{}{
	(*AdminToken)(nil),            // 0: admintoken.AdminToken
	(*GenerateTokenRequest)(nil),  // 1: admintoken.GenerateTokenRequest
	(*GenerateTokenResponse)(nil), // 2: admintoken.GenerateTokenResponse
	(*GetExistTokenRequest)(nil),  // 3: admintoken.GetExistTokenRequest
	(*GetExistTokenResponse)(nil), // 4: admintoken.GetExistTokenResponse
	(*DeleteTokenRequest)(nil),    // 5: admintoken.DeleteTokenRequest
	(*DeleteTokenResponse)(nil),   // 6: admintoken.DeleteTokenResponse
}
var file_admintoken_proto_depIdxs = []int32{
	0, // 0: admintoken.GenerateTokenResponse.adminToken:type_name -> admintoken.AdminToken
	0, // 1: admintoken.GetExistTokenResponse.adminToken:type_name -> admintoken.AdminToken
	1, // 2: admintoken.admintoken.generateToken:input_type -> admintoken.GenerateTokenRequest
	3, // 3: admintoken.admintoken.getExistToken:input_type -> admintoken.GetExistTokenRequest
	5, // 4: admintoken.admintoken.deleteToken:input_type -> admintoken.DeleteTokenRequest
	2, // 5: admintoken.admintoken.generateToken:output_type -> admintoken.GenerateTokenResponse
	4, // 6: admintoken.admintoken.getExistToken:output_type -> admintoken.GetExistTokenResponse
	6, // 7: admintoken.admintoken.deleteToken:output_type -> admintoken.DeleteTokenResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admintoken_proto_init() }
func file_admintoken_proto_init() {
	if File_admintoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admintoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_admintoken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequest); i {
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
		file_admintoken_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResponse); i {
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
		file_admintoken_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExistTokenRequest); i {
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
		file_admintoken_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExistTokenResponse); i {
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
		file_admintoken_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_admintoken_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_admintoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admintoken_proto_goTypes,
		DependencyIndexes: file_admintoken_proto_depIdxs,
		MessageInfos:      file_admintoken_proto_msgTypes,
	}.Build()
	File_admintoken_proto = out.File
	file_admintoken_proto_rawDesc = nil
	file_admintoken_proto_goTypes = nil
	file_admintoken_proto_depIdxs = nil
}
