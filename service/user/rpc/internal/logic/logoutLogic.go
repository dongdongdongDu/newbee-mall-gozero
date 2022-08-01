package logic

import (
	"context"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutRequest) (*user.LogoutResponse, error) {
	// 删除token
	_, err := l.svcCtx.UserTokenRpc.DeleteToken(l.ctx, &usertoken.DeleteTokenRequest{
		Token: in.Token,
	})
	if err != nil {
		logx.Error("token不存在", err.Error())
		return nil, err
	}
	return &user.LogoutResponse{}, nil
}
