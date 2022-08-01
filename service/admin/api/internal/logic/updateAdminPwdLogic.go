package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminPwdLogic {
	return &UpdateAdminPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminPwdLogic) UpdateAdminPwd(req *types.UpdateAdminPwdRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
