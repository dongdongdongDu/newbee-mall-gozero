package logic

import (
	"context"
	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogoutLogic {
	return &AdminLogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLogoutLogic) AdminLogout(in *admin.AdminLogoutRequest) (*admin.AdminLogoutResponse, error) {
	// 删除token
	_, err := l.svcCtx.AdminTokenRpc.DeleteToken(l.ctx, &admintoken.DeleteTokenRequest{
		Token: in.AdminToken,
	})
	if err != nil {
		logx.Error("token不存在", err.Error())
		return nil, err
	}

	return &admin.AdminLogoutResponse{}, nil
}
