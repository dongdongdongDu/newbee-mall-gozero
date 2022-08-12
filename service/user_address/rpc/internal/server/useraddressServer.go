// Code generated by goctl. DO NOT EDIT!
// Source: useraddress.proto

package server

import (
	"context"

	"newbee-mall-gozero/service/user_address/rpc/internal/logic"
	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"
)

type UseraddressServer struct {
	svcCtx *svc.ServiceContext
	useraddress.UnimplementedUseraddressServer
}

func NewUseraddressServer(svcCtx *svc.ServiceContext) *UseraddressServer {
	return &UseraddressServer{
		svcCtx: svcCtx,
	}
}

func (s *UseraddressServer) SaveAddress(ctx context.Context, in *useraddress.SaveAddressRequest) (*useraddress.EmptyResponse, error) {
	l := logic.NewSaveAddressLogic(ctx, s.svcCtx)
	return l.SaveAddress(in)
}

func (s *UseraddressServer) UpdateAddress(ctx context.Context, in *useraddress.UpdateAddressRequest) (*useraddress.EmptyResponse, error) {
	l := logic.NewUpdateAddressLogic(ctx, s.svcCtx)
	return l.UpdateAddress(in)
}

func (s *UseraddressServer) DeleteAddress(ctx context.Context, in *useraddress.DeleteAddressRequest) (*useraddress.EmptyResponse, error) {
	l := logic.NewDeleteAddressLogic(ctx, s.svcCtx)
	return l.DeleteAddress(in)
}

func (s *UseraddressServer) GetMyAddress(ctx context.Context, in *useraddress.GetMyAddressRequest) (*useraddress.GetMyAddressResponse, error) {
	l := logic.NewGetMyAddressLogic(ctx, s.svcCtx)
	return l.GetMyAddress(in)
}

func (s *UseraddressServer) GetAddressById(ctx context.Context, in *useraddress.GetAddressByIdRequest) (*useraddress.GetAddressByIdResponse, error) {
	l := logic.NewGetAddressByIdLogic(ctx, s.svcCtx)
	return l.GetAddressById(in)
}

func (s *UseraddressServer) GetDefaultAddress(ctx context.Context, in *useraddress.GetDefaultAddressRequest) (*useraddress.GetDefaultAddressResponse, error) {
	l := logic.NewGetDefaultAddressLogic(ctx, s.svcCtx)
	return l.GetDefaultAddress(in)
}
