package logic

import (
	"context"
	"errors"
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
	if in.LockStatus != 0 && in.LockStatus != 1 {
		logx.Error("非法操作")
		return nil, errors.New("非法操作")
	}

	for _, id := range in.Ids {
		// 查找
		userRes, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
		if err != nil {
			logx.Error("不存在的用户id")
			return nil, errors.New("不存在的用户id")
		}
		// 更新锁定
		userRes.LockedFlag = in.LockStatus
		err = l.svcCtx.UserModel.Update(l.ctx, userRes)
		if err != nil {
			logx.Error("用户锁定更新失败" + err.Error())
			return nil, errors.New("用户锁定更新失败" + err.Error())
		}
	}
	return &admin.LockUserResponse{}, nil
}
