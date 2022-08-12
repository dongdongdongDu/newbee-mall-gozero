package admin

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"
	"newbee-mall-gozero/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogoutLogic {
	return &AdminLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogoutLogic) AdminLogout(req *types.LogoutRequest) (resp *types.Response, err error) {
	// 退出登录
	_, err = l.svcCtx.AdminRpc.AdminLogout(l.ctx, &admin.AdminLogoutRequest{
		AdminToken: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "登录失败",
		}, nil
	}
	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "登出成功",
		Data:       map[string]interface{}{},
	}, nil
}
