package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminPwdLogic {
	return &UpdateAdminPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminPwdLogic) UpdateAdminPwd(in *admin.UpdateAdminPwdRequest) (*admin.UpdateAdminPwdResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.UpdateAdminPwdResponse{}, nil
}
