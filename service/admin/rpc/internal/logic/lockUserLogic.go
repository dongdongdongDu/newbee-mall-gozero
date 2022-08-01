package logic

import (
	"context"

	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LockUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLockUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LockUserLogic {
	return &LockUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LockUserLogic) LockUser(in *admin.LockUserRequest) (*admin.LockUserResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.LockUserResponse{}, nil
}
