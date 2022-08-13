// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: shoppingcart.proto

package shoppingcart

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
		mi := &file_shoppingcart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[0]
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
	return file_shoppingcart_proto_rawDescGZIP(), []int{0}
}

// 添加购物车项
type AddCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId    int64 `protobuf:"varint,1,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	GoodsCount int64 `protobuf:"varint,2,opt,name=goodsCount,proto3" json:"goodsCount,omitempty"`
	UserId     int64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddCartItemRequest) Reset() {
	*x = AddCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartItemRequest) ProtoMessage() {}

func (x *AddCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartItemRequest.ProtoReflect.Descriptor instead.
func (*AddCartItemRequest) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{1}
}

func (x *AddCartItemRequest) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *AddCartItemRequest) GetGoodsCount() int64 {
	if x != nil {
		return x.GoodsCount
	}
	return 0
}

func (x *AddCartItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 修改购物车项
type UpdateCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId int64 `protobuf:"varint,1,opt,name=cartItemId,proto3" json:"cartItemId,omitempty"`
	GoodsCount int64 `protobuf:"varint,2,opt,name=goodsCount,proto3" json:"goodsCount,omitempty"`
	UserId     int64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateCartItemRequest) Reset() {
	*x = UpdateCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartItemRequest) ProtoMessage() {}

func (x *UpdateCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateCartItemRequest) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCartItemRequest) GetCartItemId() int64 {
	if x != nil {
		return x.CartItemId
	}
	return 0
}

func (x *UpdateCartItemRequest) GetGoodsCount() int64 {
	if x != nil {
		return x.GoodsCount
	}
	return 0
}

func (x *UpdateCartItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 删除购物车项
type DeleteCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteCartItemRequest) Reset() {
	*x = DeleteCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartItemRequest) ProtoMessage() {}

func (x *DeleteCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteCartItemRequest) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCartItemRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteCartItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 获取购物车内容
type GetCartListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetCartListRequest) Reset() {
	*x = GetCartListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartListRequest) ProtoMessage() {}

func (x *GetCartListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartListRequest.ProtoReflect.Descriptor instead.
func (*GetCartListRequest) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId    int64  `protobuf:"varint,1,opt,name=cartItemId,proto3" json:"cartItemId,omitempty"`
	GoodsId       int64  `protobuf:"varint,2,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	GoodsCount    int64  `protobuf:"varint,3,opt,name=goodsCount,proto3" json:"goodsCount,omitempty"`
	GoodsName     string `protobuf:"bytes,4,opt,name=goodsName,proto3" json:"goodsName,omitempty"`
	GoodsCoverImg string `protobuf:"bytes,5,opt,name=goodsCoverImg,proto3" json:"goodsCoverImg,omitempty"`
	SellingPrice  int64  `protobuf:"varint,6,opt,name=sellingPrice,proto3" json:"sellingPrice,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{5}
}

func (x *CartItem) GetCartItemId() int64 {
	if x != nil {
		return x.CartItemId
	}
	return 0
}

func (x *CartItem) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *CartItem) GetGoodsCount() int64 {
	if x != nil {
		return x.GoodsCount
	}
	return 0
}

func (x *CartItem) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *CartItem) GetGoodsCoverImg() string {
	if x != nil {
		return x.GoodsCoverImg
	}
	return ""
}

func (x *CartItem) GetSellingPrice() int64 {
	if x != nil {
		return x.SellingPrice
	}
	return 0
}

type GetCartListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartList []*CartItem `protobuf:"bytes,1,rep,name=cartList,proto3" json:"cartList,omitempty"`
}

func (x *GetCartListResponse) Reset() {
	*x = GetCartListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartListResponse) ProtoMessage() {}

func (x *GetCartListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartListResponse.ProtoReflect.Descriptor instead.
func (*GetCartListResponse) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{6}
}

func (x *GetCartListResponse) GetCartList() []*CartItem {
	if x != nil {
		return x.CartList
	}
	return nil
}

// 查询明细
type GetCartItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemIds []int64 `protobuf:"varint,1,rep,packed,name=cartItemIds,proto3" json:"cartItemIds,omitempty"`
	UserId      int64   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetCartItemsRequest) Reset() {
	*x = GetCartItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartItemsRequest) ProtoMessage() {}

func (x *GetCartItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartItemsRequest.ProtoReflect.Descriptor instead.
func (*GetCartItemsRequest) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{7}
}

func (x *GetCartItemsRequest) GetCartItemIds() []int64 {
	if x != nil {
		return x.CartItemIds
	}
	return nil
}

func (x *GetCartItemsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartList []*CartItem `protobuf:"bytes,1,rep,name=cartList,proto3" json:"cartList,omitempty"`
}

func (x *GetCartItemsResponse) Reset() {
	*x = GetCartItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingcart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartItemsResponse) ProtoMessage() {}

func (x *GetCartItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingcart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartItemsResponse.ProtoReflect.Descriptor instead.
func (*GetCartItemsResponse) Descriptor() ([]byte, []int) {
	return file_shoppingcart_proto_rawDescGZIP(), []int{8}
}

func (x *GetCartItemsResponse) GetCartList() []*CartItem {
	if x != nil {
		return x.CartList
	}
	return nil
}

var File_shoppingcart_proto protoreflect.FileDescriptor

var file_shoppingcart_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xcc, 0x01, 0x0a, 0x08,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x6d, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x08, 0x63, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x72, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0xaf, 0x03, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63,
	0x61, 0x72, 0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x67, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x0c, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shoppingcart_proto_rawDescOnce sync.Once
	file_shoppingcart_proto_rawDescData = file_shoppingcart_proto_rawDesc
)

func file_shoppingcart_proto_rawDescGZIP() []byte {
	file_shoppingcart_proto_rawDescOnce.Do(func() {
		file_shoppingcart_proto_rawDescData = protoimpl.X.CompressGZIP(file_shoppingcart_proto_rawDescData)
	})
	return file_shoppingcart_proto_rawDescData
}

var file_shoppingcart_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shoppingcart_proto_goTypes = []interface{}{
	(*EmptyResponse)(nil),         // 0: shoppingcart.EmptyResponse
	(*AddCartItemRequest)(nil),    // 1: shoppingcart.AddCartItemRequest
	(*UpdateCartItemRequest)(nil), // 2: shoppingcart.UpdateCartItemRequest
	(*DeleteCartItemRequest)(nil), // 3: shoppingcart.DeleteCartItemRequest
	(*GetCartListRequest)(nil),    // 4: shoppingcart.GetCartListRequest
	(*CartItem)(nil),              // 5: shoppingcart.CartItem
	(*GetCartListResponse)(nil),   // 6: shoppingcart.GetCartListResponse
	(*GetCartItemsRequest)(nil),   // 7: shoppingcart.GetCartItemsRequest
	(*GetCartItemsResponse)(nil),  // 8: shoppingcart.GetCartItemsResponse
}
var file_shoppingcart_proto_depIdxs = []int32{
	5, // 0: shoppingcart.GetCartListResponse.cartList:type_name -> shoppingcart.CartItem
	5, // 1: shoppingcart.GetCartItemsResponse.cartList:type_name -> shoppingcart.CartItem
	1, // 2: shoppingcart.shoppingcart.addCartItem:input_type -> shoppingcart.AddCartItemRequest
	2, // 3: shoppingcart.shoppingcart.updateCartItem:input_type -> shoppingcart.UpdateCartItemRequest
	3, // 4: shoppingcart.shoppingcart.deleteCartItem:input_type -> shoppingcart.DeleteCartItemRequest
	4, // 5: shoppingcart.shoppingcart.getCartList:input_type -> shoppingcart.GetCartListRequest
	7, // 6: shoppingcart.shoppingcart.getCartItems:input_type -> shoppingcart.GetCartItemsRequest
	0, // 7: shoppingcart.shoppingcart.addCartItem:output_type -> shoppingcart.EmptyResponse
	0, // 8: shoppingcart.shoppingcart.updateCartItem:output_type -> shoppingcart.EmptyResponse
	0, // 9: shoppingcart.shoppingcart.deleteCartItem:output_type -> shoppingcart.EmptyResponse
	6, // 10: shoppingcart.shoppingcart.getCartList:output_type -> shoppingcart.GetCartListResponse
	8, // 11: shoppingcart.shoppingcart.getCartItems:output_type -> shoppingcart.GetCartItemsResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shoppingcart_proto_init() }
func file_shoppingcart_proto_init() {
	if File_shoppingcart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shoppingcart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_shoppingcart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartItemRequest); i {
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
		file_shoppingcart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartItemRequest); i {
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
		file_shoppingcart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartItemRequest); i {
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
		file_shoppingcart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartListRequest); i {
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
		file_shoppingcart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_shoppingcart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartListResponse); i {
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
		file_shoppingcart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartItemsRequest); i {
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
		file_shoppingcart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartItemsResponse); i {
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
			RawDescriptor: file_shoppingcart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shoppingcart_proto_goTypes,
		DependencyIndexes: file_shoppingcart_proto_depIdxs,
		MessageInfos:      file_shoppingcart_proto_msgTypes,
	}.Build()
	File_shoppingcart_proto = out.File
	file_shoppingcart_proto_rawDesc = nil
	file_shoppingcart_proto_goTypes = nil
	file_shoppingcart_proto_depIdxs = nil
}
