package user

import (
	"context"
	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 登录
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		LoginName:   req.LoginName,
		PasswordMd5: req.PasswordMd5,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		UserId:     res.UserId,
		Token:      res.Token,
		UpdateTime: res.UpdateTime,
		ExpireTime: res.ExpireTime,
	}, nil
}
