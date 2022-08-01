package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LockUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLockUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LockUserLogic {
	return &LockUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LockUserLogic) LockUser(req *types.LockUserRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
