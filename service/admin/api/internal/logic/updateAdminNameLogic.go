package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminNameLogic {
	return &UpdateAdminNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminNameLogic) UpdateAdminName(req *types.UpdateAdminNameRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
