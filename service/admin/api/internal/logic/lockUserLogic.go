package logic

import (
	"context"
	"encoding/json"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/rpc/admin"

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
	idsBytes, err := json.Marshal(req.Ids)
	if err != nil {
		logx.Error("ids序列化失败")
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败",
		}, nil
	}
	_, err = l.svcCtx.AdminRpc.LockUser(l.ctx, &admin.LockUserRequest{
		Ids:        string(idsBytes),
		LockStatus: req.LockStatus,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败",
		}, nil
	}
	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新成功",
		Data:       map[string]interface{}{},
	}, nil
}
