package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminProfileLogic {
	return &GetAdminProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminProfileLogic) GetAdminProfile(in *admin.GetAdminProfileRequest) (*admin.GetAdminProfileResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.GetAdminProfileResponse{}, nil
}
