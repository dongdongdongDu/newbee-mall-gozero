// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: usertoken.proto

package usertoken

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

// UsertokenClient is the client API for Usertoken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsertokenClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
	GetExistToken(ctx context.Context, in *GetExistTokenRequest, opts ...grpc.CallOption) (*GetExistTokenResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
}

type usertokenClient struct {
	cc grpc.ClientConnInterface
}

func NewUsertokenClient(cc grpc.ClientConnInterface) UsertokenClient {
	return &usertokenClient{cc}
}

func (c *usertokenClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	out := new(GenerateTokenResponse)
	err := c.cc.Invoke(ctx, "/usertoken.usertoken/generateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usertokenClient) GetExistToken(ctx context.Context, in *GetExistTokenRequest, opts ...grpc.CallOption) (*GetExistTokenResponse, error) {
	out := new(GetExistTokenResponse)
	err := c.cc.Invoke(ctx, "/usertoken.usertoken/getExistToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usertokenClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	out := new(DeleteTokenResponse)
	err := c.cc.Invoke(ctx, "/usertoken.usertoken/deleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsertokenServer is the server API for Usertoken service.
// All implementations must embed UnimplementedUsertokenServer
// for forward compatibility
type UsertokenServer interface {
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error)
	GetExistToken(context.Context, *GetExistTokenRequest) (*GetExistTokenResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error)
	mustEmbedUnimplementedUsertokenServer()
}

// UnimplementedUsertokenServer must be embedded to have forward compatible implementations.
type UnimplementedUsertokenServer struct {
}

func (UnimplementedUsertokenServer) GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUsertokenServer) GetExistToken(context.Context, *GetExistTokenRequest) (*GetExistTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExistToken not implemented")
}
func (UnimplementedUsertokenServer) DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedUsertokenServer) mustEmbedUnimplementedUsertokenServer() {}

// UnsafeUsertokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface.md is not recommended, as added methods to UsertokenServer will
// result in compilation errors.
type UnsafeUsertokenServer interface {
	mustEmbedUnimplementedUsertokenServer()
}

func RegisterUsertokenServer(s grpc.ServiceRegistrar, srv UsertokenServer) {
	s.RegisterService(&Usertoken_ServiceDesc, srv)
}

func _Usertoken_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsertokenServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usertoken.usertoken/generateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsertokenServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usertoken_GetExistToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExistTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsertokenServer).GetExistToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usertoken.usertoken/getExistToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsertokenServer).GetExistToken(ctx, req.(*GetExistTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usertoken_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsertokenServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usertoken.usertoken/deleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsertokenServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Usertoken_ServiceDesc is the grpc.ServiceDesc for Usertoken service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usertoken_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usertoken.usertoken",
	HandlerType: (*UsertokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateToken",
			Handler:    _Usertoken_GenerateToken_Handler,
		},
		{
			MethodName: "getExistToken",
			Handler:    _Usertoken_GetExistToken_Handler,
		},
		{
			MethodName: "deleteToken",
			Handler:    _Usertoken_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usertoken.proto",
}
