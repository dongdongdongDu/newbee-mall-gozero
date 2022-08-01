package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"
	"newbee-mall-gozero/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminLoginRequest) (resp *types.Response, err error) {
	// 登录
	res, err := l.svcCtx.AdminRpc.AdminLogin(l.ctx, &admin.AdminLoginRequest{
		UserName:    req.UserName,
		PasswordMd5: req.PasswordMd5,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "登录失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.AdminLoginResponse{
			Token: res.AdminToken.Token,
		},
	}, nil

}
