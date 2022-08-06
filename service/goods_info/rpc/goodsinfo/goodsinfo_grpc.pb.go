// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goodsinfo.proto

package goodsinfo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoodsinfoClient is the client API for Goodsinfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsinfoClient interface {
	// 普通用户
	GetGoodsDetail(ctx context.Context, in *GetGoodsDetailRequest, opts ...grpc.CallOption) (*GetGoodsDetailResponse, error)
	SearchGoods(ctx context.Context, in *SearchGoodsRequest, opts ...grpc.CallOption) (*SearchGoodsResponse, error)
	// 管理员
	AddGoodsInfo(ctx context.Context, in *AddGoodsInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateGoodsInfo(ctx context.Context, in *UpdateGoodsInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	AlterGoodsStatus(ctx context.Context, in *AlterGoodsStatusRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteGoodsInfo(ctx context.Context, in *DeleteGoodsInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetGoodsInfo(ctx context.Context, in *GetGoodsInfoRequest, opts ...grpc.CallOption) (*GetGoodsInfoResponse, error)
	GetGoodsList(ctx context.Context, in *GetGoodsListRequest, opts ...grpc.CallOption) (*GetGoodsListResponse, error)
}

type goodsinfoClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsinfoClient(cc grpc.ClientConnInterface) GoodsinfoClient {
	return &goodsinfoClient{cc}
}

func (c *goodsinfoClient) GetGoodsDetail(ctx context.Context, in *GetGoodsDetailRequest, opts ...grpc.CallOption) (*GetGoodsDetailResponse, error) {
	out := new(GetGoodsDetailResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/getGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) SearchGoods(ctx context.Context, in *SearchGoodsRequest, opts ...grpc.CallOption) (*SearchGoodsResponse, error) {
	out := new(SearchGoodsResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/searchGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) AddGoodsInfo(ctx context.Context, in *AddGoodsInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/addGoodsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) UpdateGoodsInfo(ctx context.Context, in *UpdateGoodsInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/updateGoodsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) AlterGoodsStatus(ctx context.Context, in *AlterGoodsStatusRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/alterGoodsStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) DeleteGoodsInfo(ctx context.Context, in *DeleteGoodsInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/deleteGoodsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) GetGoodsInfo(ctx context.Context, in *GetGoodsInfoRequest, opts ...grpc.CallOption) (*GetGoodsInfoResponse, error) {
	out := new(GetGoodsInfoResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/getGoodsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsinfoClient) GetGoodsList(ctx context.Context, in *GetGoodsListRequest, opts ...grpc.CallOption) (*GetGoodsListResponse, error) {
	out := new(GetGoodsListResponse)
	err := c.cc.Invoke(ctx, "/goodsinfo.goodsinfo/getGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsinfoServer is the server API for Goodsinfo service.
// All implementations must embed UnimplementedGoodsinfoServer
// for forward compatibility
type GoodsinfoServer interface {
	// 普通用户
	GetGoodsDetail(context.Context, *GetGoodsDetailRequest) (*GetGoodsDetailResponse, error)
	SearchGoods(context.Context, *SearchGoodsRequest) (*SearchGoodsResponse, error)
	// 管理员
	AddGoodsInfo(context.Context, *AddGoodsInfoRequest) (*EmptyResponse, error)
	UpdateGoodsInfo(context.Context, *UpdateGoodsInfoRequest) (*EmptyResponse, error)
	AlterGoodsStatus(context.Context, *AlterGoodsStatusRequest) (*EmptyResponse, error)
	DeleteGoodsInfo(context.Context, *DeleteGoodsInfoRequest) (*EmptyResponse, error)
	GetGoodsInfo(context.Context, *GetGoodsInfoRequest) (*GetGoodsInfoResponse, error)
	GetGoodsList(context.Context, *GetGoodsListRequest) (*GetGoodsListResponse, error)
	mustEmbedUnimplementedGoodsinfoServer()
}

// UnimplementedGoodsinfoServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsinfoServer struct {
}

func (UnimplementedGoodsinfoServer) GetGoodsDetail(context.Context, *GetGoodsDetailRequest) (*GetGoodsDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (UnimplementedGoodsinfoServer) SearchGoods(context.Context, *SearchGoodsRequest) (*SearchGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGoods not implemented")
}
func (UnimplementedGoodsinfoServer) AddGoodsInfo(context.Context, *AddGoodsInfoRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGoodsInfo not implemented")
}
func (UnimplementedGoodsinfoServer) UpdateGoodsInfo(context.Context, *UpdateGoodsInfoRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoodsInfo not implemented")
}
func (UnimplementedGoodsinfoServer) AlterGoodsStatus(context.Context, *AlterGoodsStatusRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterGoodsStatus not implemented")
}
func (UnimplementedGoodsinfoServer) DeleteGoodsInfo(context.Context, *DeleteGoodsInfoRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoodsInfo not implemented")
}
func (UnimplementedGoodsinfoServer) GetGoodsInfo(context.Context, *GetGoodsInfoRequest) (*GetGoodsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsInfo not implemented")
}
func (UnimplementedGoodsinfoServer) GetGoodsList(context.Context, *GetGoodsListRequest) (*GetGoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsList not implemented")
}
func (UnimplementedGoodsinfoServer) mustEmbedUnimplementedGoodsinfoServer() {}

// UnsafeGoodsinfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsinfoServer will
// result in compilation errors.
type UnsafeGoodsinfoServer interface {
	mustEmbedUnimplementedGoodsinfoServer()
}

func RegisterGoodsinfoServer(s grpc.ServiceRegistrar, srv GoodsinfoServer) {
	s.RegisterService(&Goodsinfo_ServiceDesc, srv)
}

func _Goodsinfo_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/getGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).GetGoodsDetail(ctx, req.(*GetGoodsDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_SearchGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).SearchGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/searchGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).SearchGoods(ctx, req.(*SearchGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_AddGoodsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).AddGoodsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/addGoodsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).AddGoodsInfo(ctx, req.(*AddGoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_UpdateGoodsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).UpdateGoodsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/updateGoodsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).UpdateGoodsInfo(ctx, req.(*UpdateGoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_AlterGoodsStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterGoodsStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).AlterGoodsStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/alterGoodsStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).AlterGoodsStatus(ctx, req.(*AlterGoodsStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_DeleteGoodsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).DeleteGoodsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/deleteGoodsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).DeleteGoodsInfo(ctx, req.(*DeleteGoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_GetGoodsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).GetGoodsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/getGoodsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).GetGoodsInfo(ctx, req.(*GetGoodsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goodsinfo_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsinfoServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsinfo.goodsinfo/getGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsinfoServer).GetGoodsList(ctx, req.(*GetGoodsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goodsinfo_ServiceDesc is the grpc.ServiceDesc for Goodsinfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goodsinfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goodsinfo.goodsinfo",
	HandlerType: (*GoodsinfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getGoodsDetail",
			Handler:    _Goodsinfo_GetGoodsDetail_Handler,
		},
		{
			MethodName: "searchGoods",
			Handler:    _Goodsinfo_SearchGoods_Handler,
		},
		{
			MethodName: "addGoodsInfo",
			Handler:    _Goodsinfo_AddGoodsInfo_Handler,
		},
		{
			MethodName: "updateGoodsInfo",
			Handler:    _Goodsinfo_UpdateGoodsInfo_Handler,
		},
		{
			MethodName: "alterGoodsStatus",
			Handler:    _Goodsinfo_AlterGoodsStatus_Handler,
		},
		{
			MethodName: "deleteGoodsInfo",
			Handler:    _Goodsinfo_DeleteGoodsInfo_Handler,
		},
		{
			MethodName: "getGoodsInfo",
			Handler:    _Goodsinfo_GetGoodsInfo_Handler,
		},
		{
			MethodName: "getGoodsList",
			Handler:    _Goodsinfo_GetGoodsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goodsinfo.proto",
}