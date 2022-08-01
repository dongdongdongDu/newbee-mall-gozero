package logic

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
	_ = req
	// 获取token
	token := l.ctx.Value("token").(string)
	// 退出登录
	_, err = l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutRequest{
		Token: token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "登录失败！",
		}, nil
	}
	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "登出成功",
		Data:       types.LogoutResponse{},
	}, nil
}
