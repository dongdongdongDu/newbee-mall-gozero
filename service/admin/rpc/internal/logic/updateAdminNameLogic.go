package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminNameLogic {
	return &UpdateAdminNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminNameLogic) UpdateAdminName(in *admin.UpdateAdminNameRequest) (*admin.UpdateAdminNameResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.UpdateAdminNameResponse{}, nil
}
