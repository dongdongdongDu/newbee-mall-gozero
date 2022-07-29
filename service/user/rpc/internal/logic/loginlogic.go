package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"newbee-mall-gozero/common/md5"
	"newbee-mall-gozero/service/user/model"

	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByLoginName(l.ctx, in.LoginName)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(500, "用户不存在")
		}
		return nil, status.Error(500, "登录失败："+err.Error())
	}

	// 判断密码是否正确
	if md5.MD5V([]byte(in.PasswordMd5)) != res.PasswordMd5 {
		return nil, status.Error(100, "密码错误")
	}

	// 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&user.GenerateTokenRequest{
		UserId: res.UserId,
	})
	if err != nil {
		return nil, status.Error(500, "生成token错误")
	}

	return &user.LoginResponse{
		UserId:     res.UserId,
		Token:      token.AccessToken,
		UpdateTime: token.UpdateTime,
		ExpireTime: token.ExpireTime,
	}, nil
}
