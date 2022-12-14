// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: useraddress.proto

package useraddress

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

// 空响应
type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[0]
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
	return file_useraddress_proto_rawDescGZIP(), []int{0}
}

// 保存地址（新增）
type SaveAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName      string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPhone     string `protobuf:"bytes,2,opt,name=userPhone,proto3" json:"userPhone,omitempty"`
	DefaultFlag   int64  `protobuf:"varint,3,opt,name=defaultFlag,proto3" json:"defaultFlag,omitempty"`
	ProvinceName  string `protobuf:"bytes,4,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
	CityName      string `protobuf:"bytes,5,opt,name=cityName,proto3" json:"cityName,omitempty"`
	RegionName    string `protobuf:"bytes,6,opt,name=regionName,proto3" json:"regionName,omitempty"`
	DetailAddress string `protobuf:"bytes,7,opt,name=detailAddress,proto3" json:"detailAddress,omitempty"`
	UserId        int64  `protobuf:"varint,8,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *SaveAddressRequest) Reset() {
	*x = SaveAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveAddressRequest) ProtoMessage() {}

func (x *SaveAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveAddressRequest.ProtoReflect.Descriptor instead.
func (*SaveAddressRequest) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{1}
}

func (x *SaveAddressRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SaveAddressRequest) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *SaveAddressRequest) GetDefaultFlag() int64 {
	if x != nil {
		return x.DefaultFlag
	}
	return 0
}

func (x *SaveAddressRequest) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *SaveAddressRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *SaveAddressRequest) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *SaveAddressRequest) GetDetailAddress() string {
	if x != nil {
		return x.DetailAddress
	}
	return ""
}

func (x *SaveAddressRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 更新地址
type UpdateAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressId     int64  `protobuf:"varint,1,opt,name=addressId,proto3" json:"addressId,omitempty"`
	UserName      string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPhone     string `protobuf:"bytes,3,opt,name=userPhone,proto3" json:"userPhone,omitempty"`
	DefaultFlag   int64  `protobuf:"varint,4,opt,name=defaultFlag,proto3" json:"defaultFlag,omitempty"`
	ProvinceName  string `protobuf:"bytes,5,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
	CityName      string `protobuf:"bytes,6,opt,name=cityName,proto3" json:"cityName,omitempty"`
	RegionName    string `protobuf:"bytes,7,opt,name=regionName,proto3" json:"regionName,omitempty"`
	DetailAddress string `protobuf:"bytes,8,opt,name=detailAddress,proto3" json:"detailAddress,omitempty"`
	UserId        int64  `protobuf:"varint,9,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateAddressRequest) Reset() {
	*x = UpdateAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddressRequest) ProtoMessage() {}

func (x *UpdateAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddressRequest.ProtoReflect.Descriptor instead.
func (*UpdateAddressRequest) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAddressRequest) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *UpdateAddressRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UpdateAddressRequest) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *UpdateAddressRequest) GetDefaultFlag() int64 {
	if x != nil {
		return x.DefaultFlag
	}
	return 0
}

func (x *UpdateAddressRequest) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *UpdateAddressRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *UpdateAddressRequest) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *UpdateAddressRequest) GetDetailAddress() string {
	if x != nil {
		return x.DetailAddress
	}
	return ""
}

func (x *UpdateAddressRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 删除地址
type DeleteAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteAddressRequest) Reset() {
	*x = DeleteAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAddressRequest) ProtoMessage() {}

func (x *DeleteAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAddressRequest.ProtoReflect.Descriptor instead.
func (*DeleteAddressRequest) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAddressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteAddressRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 地址Model
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressId     int64  `protobuf:"varint,1,opt,name=addressId,proto3" json:"addressId,omitempty"`
	UserId        int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName      string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPhone     string `protobuf:"bytes,4,opt,name=userPhone,proto3" json:"userPhone,omitempty"`
	DefaultFlag   int64  `protobuf:"varint,5,opt,name=defaultFlag,proto3" json:"defaultFlag,omitempty"`
	ProvinceName  string `protobuf:"bytes,6,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
	CityName      string `protobuf:"bytes,7,opt,name=cityName,proto3" json:"cityName,omitempty"`
	RegionName    string `protobuf:"bytes,8,opt,name=regionName,proto3" json:"regionName,omitempty"`
	DetailAddress string `protobuf:"bytes,9,opt,name=detailAddress,proto3" json:"detailAddress,omitempty"`
	IsDeleted     int64  `protobuf:"varint,10,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	CreateTime    int64  `protobuf:"varint,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime    int64  `protobuf:"varint,12,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{4}
}

func (x *Address) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Address) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Address) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Address) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *Address) GetDefaultFlag() int64 {
	if x != nil {
		return x.DefaultFlag
	}
	return 0
}

func (x *Address) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *Address) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Address) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *Address) GetDetailAddress() string {
	if x != nil {
		return x.DetailAddress
	}
	return ""
}

func (x *Address) GetIsDeleted() int64 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

func (x *Address) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Address) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

// 获取用户地址
type GetMyAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetMyAddressRequest) Reset() {
	*x = GetMyAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyAddressRequest) ProtoMessage() {}

func (x *GetMyAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyAddressRequest.ProtoReflect.Descriptor instead.
func (*GetMyAddressRequest) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{5}
}

func (x *GetMyAddressRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetMyAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressList []*Address `protobuf:"bytes,1,rep,name=addressList,proto3" json:"addressList,omitempty"`
}

func (x *GetMyAddressResponse) Reset() {
	*x = GetMyAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyAddressResponse) ProtoMessage() {}

func (x *GetMyAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyAddressResponse.ProtoReflect.Descriptor instead.
func (*GetMyAddressResponse) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{6}
}

func (x *GetMyAddressResponse) GetAddressList() []*Address {
	if x != nil {
		return x.AddressList
	}
	return nil
}

// 根据id获取地址
type GetAddressByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetAddressByIdRequest) Reset() {
	*x = GetAddressByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddressByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressByIdRequest) ProtoMessage() {}

func (x *GetAddressByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressByIdRequest.ProtoReflect.Descriptor instead.
func (*GetAddressByIdRequest) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{7}
}

func (x *GetAddressByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAddressByIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAddressByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetAddressByIdResponse) Reset() {
	*x = GetAddressByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddressByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressByIdResponse) ProtoMessage() {}

func (x *GetAddressByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressByIdResponse.ProtoReflect.Descriptor instead.
func (*GetAddressByIdResponse) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{8}
}

func (x *GetAddressByIdResponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

// 获取默认地址
type GetDefaultAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetDefaultAddressRequest) Reset() {
	*x = GetDefaultAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultAddressRequest) ProtoMessage() {}

func (x *GetDefaultAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultAddressRequest.ProtoReflect.Descriptor instead.
func (*GetDefaultAddressRequest) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{9}
}

func (x *GetDefaultAddressRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetDefaultAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetDefaultAddressResponse) Reset() {
	*x = GetDefaultAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useraddress_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultAddressResponse) ProtoMessage() {}

func (x *GetDefaultAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_useraddress_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultAddressResponse.ProtoReflect.Descriptor instead.
func (*GetDefaultAddressResponse) Descriptor() ([]byte, []int) {
	return file_useraddress_proto_rawDescGZIP(), []int{10}
}

func (x *GetDefaultAddressResponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_useraddress_proto protoreflect.FileDescriptor

var file_useraddress_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x8e, 0x02, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xae, 0x02, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6c,
	0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xff, 0x02, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6c, 0x61,
	0x67, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x32, 0x8d, 0x04, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x4a, 0x0a, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c,
	0x67, 0x65, 0x74, 0x4d, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11,
	0x67, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_useraddress_proto_rawDescOnce sync.Once
	file_useraddress_proto_rawDescData = file_useraddress_proto_rawDesc
)

func file_useraddress_proto_rawDescGZIP() []byte {
	file_useraddress_proto_rawDescOnce.Do(func() {
		file_useraddress_proto_rawDescData = protoimpl.X.CompressGZIP(file_useraddress_proto_rawDescData)
	})
	return file_useraddress_proto_rawDescData
}

var file_useraddress_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_useraddress_proto_goTypes = []interface{}{
	(*EmptyResponse)(nil),             // 0: useraddress.EmptyResponse
	(*SaveAddressRequest)(nil),        // 1: useraddress.SaveAddressRequest
	(*UpdateAddressRequest)(nil),      // 2: useraddress.UpdateAddressRequest
	(*DeleteAddressRequest)(nil),      // 3: useraddress.DeleteAddressRequest
	(*Address)(nil),                   // 4: useraddress.Address
	(*GetMyAddressRequest)(nil),       // 5: useraddress.GetMyAddressRequest
	(*GetMyAddressResponse)(nil),      // 6: useraddress.GetMyAddressResponse
	(*GetAddressByIdRequest)(nil),     // 7: useraddress.GetAddressByIdRequest
	(*GetAddressByIdResponse)(nil),    // 8: useraddress.GetAddressByIdResponse
	(*GetDefaultAddressRequest)(nil),  // 9: useraddress.GetDefaultAddressRequest
	(*GetDefaultAddressResponse)(nil), // 10: useraddress.GetDefaultAddressResponse
}
var file_useraddress_proto_depIdxs = []int32{
	4,  // 0: useraddress.GetMyAddressResponse.addressList:type_name -> useraddress.Address
	4,  // 1: useraddress.GetAddressByIdResponse.address:type_name -> useraddress.Address
	4,  // 2: useraddress.GetDefaultAddressResponse.address:type_name -> useraddress.Address
	1,  // 3: useraddress.useraddress.saveAddress:input_type -> useraddress.SaveAddressRequest
	2,  // 4: useraddress.useraddress.updateAddress:input_type -> useraddress.UpdateAddressRequest
	3,  // 5: useraddress.useraddress.deleteAddress:input_type -> useraddress.DeleteAddressRequest
	5,  // 6: useraddress.useraddress.getMyAddress:input_type -> useraddress.GetMyAddressRequest
	7,  // 7: useraddress.useraddress.getAddressById:input_type -> useraddress.GetAddressByIdRequest
	9,  // 8: useraddress.useraddress.getDefaultAddress:input_type -> useraddress.GetDefaultAddressRequest
	0,  // 9: useraddress.useraddress.saveAddress:output_type -> useraddress.EmptyResponse
	0,  // 10: useraddress.useraddress.updateAddress:output_type -> useraddress.EmptyResponse
	0,  // 11: useraddress.useraddress.deleteAddress:output_type -> useraddress.EmptyResponse
	6,  // 12: useraddress.useraddress.getMyAddress:output_type -> useraddress.GetMyAddressResponse
	8,  // 13: useraddress.useraddress.getAddressById:output_type -> useraddress.GetAddressByIdResponse
	10, // 14: useraddress.useraddress.getDefaultAddress:output_type -> useraddress.GetDefaultAddressResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_useraddress_proto_init() }
func file_useraddress_proto_init() {
	if File_useraddress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_useraddress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_useraddress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveAddressRequest); i {
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
		file_useraddress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAddressRequest); i {
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
		file_useraddress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAddressRequest); i {
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
		file_useraddress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_useraddress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyAddressRequest); i {
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
		file_useraddress_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyAddressResponse); i {
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
		file_useraddress_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddressByIdRequest); i {
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
		file_useraddress_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddressByIdResponse); i {
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
		file_useraddress_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultAddressRequest); i {
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
		file_useraddress_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultAddressResponse); i {
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
			RawDescriptor: file_useraddress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_useraddress_proto_goTypes,
		DependencyIndexes: file_useraddress_proto_depIdxs,
		MessageInfos:      file_useraddress_proto_msgTypes,
	}.Build()
	File_useraddress_proto = out.File
	file_useraddress_proto_rawDesc = nil
	file_useraddress_proto_goTypes = nil
	file_useraddress_proto_depIdxs = nil
}
