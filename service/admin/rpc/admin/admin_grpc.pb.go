// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: admin.proto

package admin

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

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	GetAdminProfile(ctx context.Context, in *GetAdminProfileRequest, opts ...grpc.CallOption) (*GetAdminProfileResponse, error)
	UpdateAdminName(ctx context.Context, in *UpdateAdminNameRequest, opts ...grpc.CallOption) (*UpdateAdminNameResponse, error)
	UpdateAdminPwd(ctx context.Context, in *UpdateAdminPwdRequest, opts ...grpc.CallOption) (*UpdateAdminPwdResponse, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
	LockUser(ctx context.Context, in *LockUserRequest, opts ...grpc.CallOption) (*LockUserResponse, error)
	AdminLogout(ctx context.Context, in *AdminLogoutRequest, opts ...grpc.CallOption) (*AdminLogoutResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/adminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetAdminProfile(ctx context.Context, in *GetAdminProfileRequest, opts ...grpc.CallOption) (*GetAdminProfileResponse, error) {
	out := new(GetAdminProfileResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/getAdminProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateAdminName(ctx context.Context, in *UpdateAdminNameRequest, opts ...grpc.CallOption) (*UpdateAdminNameResponse, error) {
	out := new(UpdateAdminNameResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/updateAdminName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateAdminPwd(ctx context.Context, in *UpdateAdminPwdRequest, opts ...grpc.CallOption) (*UpdateAdminPwdResponse, error) {
	out := new(UpdateAdminPwdResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/updateAdminPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/getUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) LockUser(ctx context.Context, in *LockUserRequest, opts ...grpc.CallOption) (*LockUserResponse, error) {
	out := new(LockUserResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/lockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminLogout(ctx context.Context, in *AdminLogoutRequest, opts ...grpc.CallOption) (*AdminLogoutResponse, error) {
	out := new(AdminLogoutResponse)
	err := c.cc.Invoke(ctx, "/admin.admin/adminLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	GetAdminProfile(context.Context, *GetAdminProfileRequest) (*GetAdminProfileResponse, error)
	UpdateAdminName(context.Context, *UpdateAdminNameRequest) (*UpdateAdminNameResponse, error)
	UpdateAdminPwd(context.Context, *UpdateAdminPwdRequest) (*UpdateAdminPwdResponse, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	LockUser(context.Context, *LockUserRequest) (*LockUserResponse, error)
	AdminLogout(context.Context, *AdminLogoutRequest) (*AdminLogoutResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServer) GetAdminProfile(context.Context, *GetAdminProfileRequest) (*GetAdminProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminProfile not implemented")
}
func (UnimplementedAdminServer) UpdateAdminName(context.Context, *UpdateAdminNameRequest) (*UpdateAdminNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdminName not implemented")
}
func (UnimplementedAdminServer) UpdateAdminPwd(context.Context, *UpdateAdminPwdRequest) (*UpdateAdminPwdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdminPwd not implemented")
}
func (UnimplementedAdminServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedAdminServer) LockUser(context.Context, *LockUserRequest) (*LockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockUser not implemented")
}
func (UnimplementedAdminServer) AdminLogout(context.Context, *AdminLogoutRequest) (*AdminLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogout not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/adminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetAdminProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetAdminProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/getAdminProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetAdminProfile(ctx, req.(*GetAdminProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateAdminName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdminNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateAdminName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/updateAdminName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateAdminName(ctx, req.(*UpdateAdminNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateAdminPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdminPwdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateAdminPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/updateAdminPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateAdminPwd(ctx, req.(*UpdateAdminPwdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/getUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_LockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).LockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/lockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).LockUser(ctx, req.(*LockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.admin/adminLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLogout(ctx, req.(*AdminLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "adminLogin",
			Handler:    _Admin_AdminLogin_Handler,
		},
		{
			MethodName: "getAdminProfile",
			Handler:    _Admin_GetAdminProfile_Handler,
		},
		{
			MethodName: "updateAdminName",
			Handler:    _Admin_UpdateAdminName_Handler,
		},
		{
			MethodName: "updateAdminPwd",
			Handler:    _Admin_UpdateAdminPwd_Handler,
		},
		{
			MethodName: "getUserList",
			Handler:    _Admin_GetUserList_Handler,
		},
		{
			MethodName: "lockUser",
			Handler:    _Admin_LockUser_Handler,
		},
		{
			MethodName: "adminLogout",
			Handler:    _Admin_AdminLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
