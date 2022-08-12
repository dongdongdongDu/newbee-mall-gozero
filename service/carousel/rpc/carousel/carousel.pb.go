// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: carousel.proto

package carousel

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

// 获取首页的轮播图
type GetCarouselsForIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetCarouselsForIndexRequest) Reset() {
	*x = GetCarouselsForIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarouselsForIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarouselsForIndexRequest) ProtoMessage() {}

func (x *GetCarouselsForIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarouselsForIndexRequest.ProtoReflect.Descriptor instead.
func (*GetCarouselsForIndexRequest) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{0}
}

func (x *GetCarouselsForIndexRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type CarouselsForIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselUrl string `protobuf:"bytes,1,opt,name=carouselUrl,proto3" json:"carouselUrl,omitempty"`
	RedirectUrl string `protobuf:"bytes,2,opt,name=redirectUrl,proto3" json:"redirectUrl,omitempty"`
}

func (x *CarouselsForIndex) Reset() {
	*x = CarouselsForIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarouselsForIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarouselsForIndex) ProtoMessage() {}

func (x *CarouselsForIndex) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarouselsForIndex.ProtoReflect.Descriptor instead.
func (*CarouselsForIndex) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{1}
}

func (x *CarouselsForIndex) GetCarouselUrl() string {
	if x != nil {
		return x.CarouselUrl
	}
	return ""
}

func (x *CarouselsForIndex) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type GetCarouselsForIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselIndex []*CarouselsForIndex `protobuf:"bytes,1,rep,name=carouselIndex,proto3" json:"carouselIndex,omitempty"`
}

func (x *GetCarouselsForIndexResponse) Reset() {
	*x = GetCarouselsForIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarouselsForIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarouselsForIndexResponse) ProtoMessage() {}

func (x *GetCarouselsForIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarouselsForIndexResponse.ProtoReflect.Descriptor instead.
func (*GetCarouselsForIndexResponse) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{2}
}

func (x *GetCarouselsForIndexResponse) GetCarouselIndex() []*CarouselsForIndex {
	if x != nil {
		return x.CarouselIndex
	}
	return nil
}

// 增加轮播图
type AddCarouselRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselUrl  string `protobuf:"bytes,1,opt,name=carouselUrl,proto3" json:"carouselUrl,omitempty"`
	RedirectUrl  string `protobuf:"bytes,2,opt,name=redirectUrl,proto3" json:"redirectUrl,omitempty"`
	CarouselRank int64  `protobuf:"varint,3,opt,name=carouselRank,proto3" json:"carouselRank,omitempty"`
	User         int64  `protobuf:"varint,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddCarouselRequest) Reset() {
	*x = AddCarouselRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCarouselRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCarouselRequest) ProtoMessage() {}

func (x *AddCarouselRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCarouselRequest.ProtoReflect.Descriptor instead.
func (*AddCarouselRequest) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{3}
}

func (x *AddCarouselRequest) GetCarouselUrl() string {
	if x != nil {
		return x.CarouselUrl
	}
	return ""
}

func (x *AddCarouselRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *AddCarouselRequest) GetCarouselRank() int64 {
	if x != nil {
		return x.CarouselRank
	}
	return 0
}

func (x *AddCarouselRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

// 更新轮播图
type UpdateCarouselRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselId   int64  `protobuf:"varint,1,opt,name=carouselId,proto3" json:"carouselId,omitempty"`
	CarouselUrl  string `protobuf:"bytes,2,opt,name=carouselUrl,proto3" json:"carouselUrl,omitempty"`
	RedirectUrl  string `protobuf:"bytes,3,opt,name=redirectUrl,proto3" json:"redirectUrl,omitempty"`
	CarouselRank int64  `protobuf:"varint,4,opt,name=carouselRank,proto3" json:"carouselRank,omitempty"`
	User         int64  `protobuf:"varint,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateCarouselRequest) Reset() {
	*x = UpdateCarouselRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarouselRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarouselRequest) ProtoMessage() {}

func (x *UpdateCarouselRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarouselRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarouselRequest) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCarouselRequest) GetCarouselId() int64 {
	if x != nil {
		return x.CarouselId
	}
	return 0
}

func (x *UpdateCarouselRequest) GetCarouselUrl() string {
	if x != nil {
		return x.CarouselUrl
	}
	return ""
}

func (x *UpdateCarouselRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *UpdateCarouselRequest) GetCarouselRank() int64 {
	if x != nil {
		return x.CarouselRank
	}
	return 0
}

func (x *UpdateCarouselRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

// 删除轮播图
type DeleteCarouselRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteCarouselRequest) Reset() {
	*x = DeleteCarouselRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarouselRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarouselRequest) ProtoMessage() {}

func (x *DeleteCarouselRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarouselRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarouselRequest) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCarouselRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Carousel Model
type CarouselType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselId   int64  `protobuf:"varint,1,opt,name=carouselId,proto3" json:"carouselId,omitempty"`
	CarouselUrl  string `protobuf:"bytes,2,opt,name=carouselUrl,proto3" json:"carouselUrl,omitempty"`
	RedirectUrl  string `protobuf:"bytes,3,opt,name=redirectUrl,proto3" json:"redirectUrl,omitempty"`
	CarouselRank int64  `protobuf:"varint,4,opt,name=carouselRank,proto3" json:"carouselRank,omitempty"`
	IsDeleted    int64  `protobuf:"varint,5,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	CreateTime   int64  `protobuf:"varint,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CreateUser   int64  `protobuf:"varint,7,opt,name=createUser,proto3" json:"createUser,omitempty"`
	UpdateTime   int64  `protobuf:"varint,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateUser   int64  `protobuf:"varint,9,opt,name=updateUser,proto3" json:"updateUser,omitempty"`
}

func (x *CarouselType) Reset() {
	*x = CarouselType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarouselType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarouselType) ProtoMessage() {}

func (x *CarouselType) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarouselType.ProtoReflect.Descriptor instead.
func (*CarouselType) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{6}
}

func (x *CarouselType) GetCarouselId() int64 {
	if x != nil {
		return x.CarouselId
	}
	return 0
}

func (x *CarouselType) GetCarouselUrl() string {
	if x != nil {
		return x.CarouselUrl
	}
	return ""
}

func (x *CarouselType) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *CarouselType) GetCarouselRank() int64 {
	if x != nil {
		return x.CarouselRank
	}
	return 0
}

func (x *CarouselType) GetIsDeleted() int64 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

func (x *CarouselType) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarouselType) GetCreateUser() int64 {
	if x != nil {
		return x.CreateUser
	}
	return 0
}

func (x *CarouselType) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *CarouselType) GetUpdateUser() int64 {
	if x != nil {
		return x.UpdateUser
	}
	return 0
}

// 查找
type GetCarouselRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCarouselRequest) Reset() {
	*x = GetCarouselRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarouselRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarouselRequest) ProtoMessage() {}

func (x *GetCarouselRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarouselRequest.ProtoReflect.Descriptor instead.
func (*GetCarouselRequest) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{7}
}

func (x *GetCarouselRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCarouselResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carousel *CarouselType `protobuf:"bytes,1,opt,name=carousel,proto3" json:"carousel,omitempty"`
}

func (x *GetCarouselResponse) Reset() {
	*x = GetCarouselResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarouselResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarouselResponse) ProtoMessage() {}

func (x *GetCarouselResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarouselResponse.ProtoReflect.Descriptor instead.
func (*GetCarouselResponse) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{8}
}

func (x *GetCarouselResponse) GetCarousel() *CarouselType {
	if x != nil {
		return x.Carousel
	}
	return nil
}

// 查找列表
type GetCarouselListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int64 `protobuf:"varint,1,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	PageSize   int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetCarouselListRequest) Reset() {
	*x = GetCarouselListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarouselListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarouselListRequest) ProtoMessage() {}

func (x *GetCarouselListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarouselListRequest.ProtoReflect.Descriptor instead.
func (*GetCarouselListRequest) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{9}
}

func (x *GetCarouselListRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetCarouselListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetCarouselListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselList []*CarouselType `protobuf:"bytes,1,rep,name=carouselList,proto3" json:"carouselList,omitempty"`
	Total        int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetCarouselListResponse) Reset() {
	*x = GetCarouselListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carousel_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarouselListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarouselListResponse) ProtoMessage() {}

func (x *GetCarouselListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarouselListResponse.ProtoReflect.Descriptor instead.
func (*GetCarouselListResponse) Descriptor() ([]byte, []int) {
	return file_carousel_proto_rawDescGZIP(), []int{10}
}

func (x *GetCarouselListResponse) GetCarouselList() []*CarouselType {
	if x != nil {
		return x.CarouselList
	}
	return nil
}

func (x *GetCarouselListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
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
		mi := &file_carousel_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_carousel_proto_msgTypes[11]
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
	return file_carousel_proto_rawDescGZIP(), []int{11}
}

var File_carousel_proto protoreflect.FileDescriptor

var file_carousel_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x22, 0x2f, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x57, 0x0a, 0x11, 0x43,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x55,
	0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x61, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x61,
	0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73,
	0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x55, 0x72, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x61,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xb3, 0x01, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c,
	0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x6f,
	0x75, 0x73, 0x65, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x0c,
	0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x61, 0x6e, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x72,
	0x6f, 0x75, 0x73, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73,
	0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6b, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x61, 0x72,
	0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf3, 0x03, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x6f,
	0x75, 0x73, 0x65, 0x6c, 0x12, 0x65, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x2e, 0x63,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x61,
	0x64, 0x64, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72,
	0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a,
	0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x12,
	0x1f, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x67, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f,
	0x75, 0x73, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x72,
	0x6f, 0x75, 0x73, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x63, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_carousel_proto_rawDescOnce sync.Once
	file_carousel_proto_rawDescData = file_carousel_proto_rawDesc
)

func file_carousel_proto_rawDescGZIP() []byte {
	file_carousel_proto_rawDescOnce.Do(func() {
		file_carousel_proto_rawDescData = protoimpl.X.CompressGZIP(file_carousel_proto_rawDescData)
	})
	return file_carousel_proto_rawDescData
}

var file_carousel_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_carousel_proto_goTypes = []interface{}{
	(*GetCarouselsForIndexRequest)(nil),  // 0: carousel.GetCarouselsForIndexRequest
	(*CarouselsForIndex)(nil),            // 1: carousel.CarouselsForIndex
	(*GetCarouselsForIndexResponse)(nil), // 2: carousel.GetCarouselsForIndexResponse
	(*AddCarouselRequest)(nil),           // 3: carousel.AddCarouselRequest
	(*UpdateCarouselRequest)(nil),        // 4: carousel.UpdateCarouselRequest
	(*DeleteCarouselRequest)(nil),        // 5: carousel.DeleteCarouselRequest
	(*CarouselType)(nil),                 // 6: carousel.CarouselType
	(*GetCarouselRequest)(nil),           // 7: carousel.GetCarouselRequest
	(*GetCarouselResponse)(nil),          // 8: carousel.GetCarouselResponse
	(*GetCarouselListRequest)(nil),       // 9: carousel.GetCarouselListRequest
	(*GetCarouselListResponse)(nil),      // 10: carousel.GetCarouselListResponse
	(*EmptyResponse)(nil),                // 11: carousel.EmptyResponse
}
var file_carousel_proto_depIdxs = []int32{
	1,  // 0: carousel.GetCarouselsForIndexResponse.carouselIndex:type_name -> carousel.CarouselsForIndex
	6,  // 1: carousel.GetCarouselResponse.carousel:type_name -> carousel.CarouselType
	6,  // 2: carousel.GetCarouselListResponse.carouselList:type_name -> carousel.CarouselType
	0,  // 3: carousel.carousel.getCarouselsForIndex:input_type -> carousel.GetCarouselsForIndexRequest
	3,  // 4: carousel.carousel.addCarousel:input_type -> carousel.AddCarouselRequest
	4,  // 5: carousel.carousel.updateCarousel:input_type -> carousel.UpdateCarouselRequest
	5,  // 6: carousel.carousel.deleteCarousel:input_type -> carousel.DeleteCarouselRequest
	7,  // 7: carousel.carousel.getCarousel:input_type -> carousel.GetCarouselRequest
	9,  // 8: carousel.carousel.getCarouselList:input_type -> carousel.GetCarouselListRequest
	2,  // 9: carousel.carousel.getCarouselsForIndex:output_type -> carousel.GetCarouselsForIndexResponse
	11, // 10: carousel.carousel.addCarousel:output_type -> carousel.EmptyResponse
	11, // 11: carousel.carousel.updateCarousel:output_type -> carousel.EmptyResponse
	11, // 12: carousel.carousel.deleteCarousel:output_type -> carousel.EmptyResponse
	8,  // 13: carousel.carousel.getCarousel:output_type -> carousel.GetCarouselResponse
	10, // 14: carousel.carousel.getCarouselList:output_type -> carousel.GetCarouselListResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_carousel_proto_init() }
func file_carousel_proto_init() {
	if File_carousel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carousel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarouselsForIndexRequest); i {
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
		file_carousel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarouselsForIndex); i {
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
		file_carousel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarouselsForIndexResponse); i {
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
		file_carousel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCarouselRequest); i {
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
		file_carousel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarouselRequest); i {
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
		file_carousel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarouselRequest); i {
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
		file_carousel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarouselType); i {
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
		file_carousel_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarouselRequest); i {
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
		file_carousel_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarouselResponse); i {
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
		file_carousel_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarouselListRequest); i {
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
		file_carousel_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarouselListResponse); i {
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
		file_carousel_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_carousel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carousel_proto_goTypes,
		DependencyIndexes: file_carousel_proto_depIdxs,
		MessageInfos:      file_carousel_proto_msgTypes,
	}.Build()
	File_carousel_proto = out.File
	file_carousel_proto_rawDesc = nil
	file_carousel_proto_goTypes = nil
	file_carousel_proto_depIdxs = nil
}