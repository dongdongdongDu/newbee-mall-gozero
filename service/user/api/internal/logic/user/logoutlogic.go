package user

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutRequest) (resp *types.Response, err error) {
	// 退出登录
	_, err = l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "登出失败",
		}, nil
	}
	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "登出成功",
		Data:       map[string]interface{}{},
	}, nil
}
