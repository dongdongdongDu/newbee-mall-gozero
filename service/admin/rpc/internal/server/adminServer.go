// Code generated by goctl. DO NOT EDIT!
// Source: admin.proto

package server

import (
	"context"

	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/logic"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"
)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminServer) AdminLogin(ctx context.Context, in *admin.AdminLoginRequest) (*admin.AdminLoginResponse, error) {
	l := logic.NewAdminLoginLogic(ctx, s.svcCtx)
	return l.AdminLogin(in)
}

func (s *AdminServer) GetAdminProfile(ctx context.Context, in *admin.GetAdminProfileRequest) (*admin.GetAdminProfileResponse, error) {
	l := logic.NewGetAdminProfileLogic(ctx, s.svcCtx)
	return l.GetAdminProfile(in)
}

func (s *AdminServer) UpdateAdminName(ctx context.Context, in *admin.UpdateAdminNameRequest) (*admin.EmptyResponse, error) {
	l := logic.NewUpdateAdminNameLogic(ctx, s.svcCtx)
	return l.UpdateAdminName(in)
}

func (s *AdminServer) UpdateAdminPwd(ctx context.Context, in *admin.UpdateAdminPwdRequest) (*admin.EmptyResponse, error) {
	l := logic.NewUpdateAdminPwdLogic(ctx, s.svcCtx)
	return l.UpdateAdminPwd(in)
}

func (s *AdminServer) AdminLogout(ctx context.Context, in *admin.AdminLogoutRequest) (*admin.EmptyResponse, error) {
	l := logic.NewAdminLogoutLogic(ctx, s.svcCtx)
	return l.AdminLogout(in)
}
