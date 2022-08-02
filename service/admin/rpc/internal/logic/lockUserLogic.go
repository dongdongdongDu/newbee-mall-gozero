package logic

import (
	"context"
	"encoding/json"
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
	// todo: add your logic here and delete this line
	if in.LockStatus != 0 && in.LockStatus != 1 {
		logx.Error("非法操作")
		return nil, errors.New("非法操作")
	}

	var ids []int64
	err := json.Unmarshal([]byte(in.Ids), &ids)
	if err != nil {
		logx.Error("ids序列化失败")
		return nil, errors.New("ids序列化失败")
	}
	for _, id := range ids {
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
