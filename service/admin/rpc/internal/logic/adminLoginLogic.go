package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/admin/model"
	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginLogic) AdminLogin(in *admin.AdminLoginRequest) (*admin.AdminLoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.AdminModel.FindOneByLoginUserName(l.ctx, in.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("用户不存在")
			return nil, err
		}
		logx.Error("登录失败：" + err.Error())
		return nil, err
	}

	// 判断密码是否正确
	if in.PasswordMd5 != res.LoginPassword {
		logx.Error("密码错误")
		return nil, errors.New("密码错误")
	}

	// 生成token
	token, err := l.svcCtx.AdminTokenRpc.GenerateToken(l.ctx, &admintoken.GenerateTokenRequest{
		AdminUserId: res.AdminUserId,
	})
	if err != nil {
		logx.Error("生成token错误")
		return nil, errors.New("生成token错误")
	}
	// 复制
	var adminToken admin.AdminToken
	err = copier.Copy(&adminToken, token.AdminToken)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}

	return &admin.AdminLoginResponse{
		AdminToken: &adminToken,
	}, nil
}
