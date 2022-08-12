// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: useraddress.proto

package useraddress

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

// UseraddressClient is the client API for Useraddress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UseraddressClient interface {
	SaveAddress(ctx context.Context, in *SaveAddressRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetMyAddress(ctx context.Context, in *GetMyAddressRequest, opts ...grpc.CallOption) (*GetMyAddressResponse, error)
	GetAddressById(ctx context.Context, in *GetAddressByIdRequest, opts ...grpc.CallOption) (*GetAddressByIdResponse, error)
	GetDefaultAddress(ctx context.Context, in *GetDefaultAddressRequest, opts ...grpc.CallOption) (*GetDefaultAddressResponse, error)
}

type useraddressClient struct {
	cc grpc.ClientConnInterface
}

func NewUseraddressClient(cc grpc.ClientConnInterface) UseraddressClient {
	return &useraddressClient{cc}
}

func (c *useraddressClient) SaveAddress(ctx context.Context, in *SaveAddressRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/useraddress.useraddress/saveAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useraddressClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/useraddress.useraddress/updateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useraddressClient) DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/useraddress.useraddress/deleteAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useraddressClient) GetMyAddress(ctx context.Context, in *GetMyAddressRequest, opts ...grpc.CallOption) (*GetMyAddressResponse, error) {
	out := new(GetMyAddressResponse)
	err := c.cc.Invoke(ctx, "/useraddress.useraddress/getMyAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useraddressClient) GetAddressById(ctx context.Context, in *GetAddressByIdRequest, opts ...grpc.CallOption) (*GetAddressByIdResponse, error) {
	out := new(GetAddressByIdResponse)
	err := c.cc.Invoke(ctx, "/useraddress.useraddress/getAddressById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useraddressClient) GetDefaultAddress(ctx context.Context, in *GetDefaultAddressRequest, opts ...grpc.CallOption) (*GetDefaultAddressResponse, error) {
	out := new(GetDefaultAddressResponse)
	err := c.cc.Invoke(ctx, "/useraddress.useraddress/getDefaultAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UseraddressServer is the server API for Useraddress service.
// All implementations must embed UnimplementedUseraddressServer
// for forward compatibility
type UseraddressServer interface {
	SaveAddress(context.Context, *SaveAddressRequest) (*EmptyResponse, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*EmptyResponse, error)
	DeleteAddress(context.Context, *DeleteAddressRequest) (*EmptyResponse, error)
	GetMyAddress(context.Context, *GetMyAddressRequest) (*GetMyAddressResponse, error)
	GetAddressById(context.Context, *GetAddressByIdRequest) (*GetAddressByIdResponse, error)
	GetDefaultAddress(context.Context, *GetDefaultAddressRequest) (*GetDefaultAddressResponse, error)
	mustEmbedUnimplementedUseraddressServer()
}

// UnimplementedUseraddressServer must be embedded to have forward compatible implementations.
type UnimplementedUseraddressServer struct {
}

func (UnimplementedUseraddressServer) SaveAddress(context.Context, *SaveAddressRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAddress not implemented")
}
func (UnimplementedUseraddressServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedUseraddressServer) DeleteAddress(context.Context, *DeleteAddressRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedUseraddressServer) GetMyAddress(context.Context, *GetMyAddressRequest) (*GetMyAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyAddress not implemented")
}
func (UnimplementedUseraddressServer) GetAddressById(context.Context, *GetAddressByIdRequest) (*GetAddressByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressById not implemented")
}
func (UnimplementedUseraddressServer) GetDefaultAddress(context.Context, *GetDefaultAddressRequest) (*GetDefaultAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultAddress not implemented")
}
func (UnimplementedUseraddressServer) mustEmbedUnimplementedUseraddressServer() {}

// UnsafeUseraddressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UseraddressServer will
// result in compilation errors.
type UnsafeUseraddressServer interface {
	mustEmbedUnimplementedUseraddressServer()
}

func RegisterUseraddressServer(s grpc.ServiceRegistrar, srv UseraddressServer) {
	s.RegisterService(&Useraddress_ServiceDesc, srv)
}

func _Useraddress_SaveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseraddressServer).SaveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useraddress.useraddress/saveAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseraddressServer).SaveAddress(ctx, req.(*SaveAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Useraddress_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseraddressServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useraddress.useraddress/updateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseraddressServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Useraddress_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseraddressServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useraddress.useraddress/deleteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseraddressServer).DeleteAddress(ctx, req.(*DeleteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Useraddress_GetMyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseraddressServer).GetMyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useraddress.useraddress/getMyAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseraddressServer).GetMyAddress(ctx, req.(*GetMyAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Useraddress_GetAddressById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseraddressServer).GetAddressById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useraddress.useraddress/getAddressById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseraddressServer).GetAddressById(ctx, req.(*GetAddressByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Useraddress_GetDefaultAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseraddressServer).GetDefaultAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/useraddress.useraddress/getDefaultAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseraddressServer).GetDefaultAddress(ctx, req.(*GetDefaultAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Useraddress_ServiceDesc is the grpc.ServiceDesc for Useraddress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Useraddress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "useraddress.useraddress",
	HandlerType: (*UseraddressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveAddress",
			Handler:    _Useraddress_SaveAddress_Handler,
		},
		{
			MethodName: "updateAddress",
			Handler:    _Useraddress_UpdateAddress_Handler,
		},
		{
			MethodName: "deleteAddress",
			Handler:    _Useraddress_DeleteAddress_Handler,
		},
		{
			MethodName: "getMyAddress",
			Handler:    _Useraddress_GetMyAddress_Handler,
		},
		{
			MethodName: "getAddressById",
			Handler:    _Useraddress_GetAddressById_Handler,
		},
		{
			MethodName: "getDefaultAddress",
			Handler:    _Useraddress_GetDefaultAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "useraddress.proto",
}
