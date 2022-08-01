package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminProfileLogic {
	return &GetAdminProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminProfileLogic) GetAdminProfile() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
