package logic

import (
	"context"
	"errors"

	"newbee-mall-gozero/service/admin/model"
	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminProfileLogic {
	return &GetAdminProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminProfileLogic) GetAdminProfile(in *admin.GetAdminProfileRequest) (*admin.GetAdminProfileResponse, error) {
	// 查询token
	tokenRes, err := l.svcCtx.AdminTokenRpc.GetExistToken(l.ctx, &admintoken.GetExistTokenRequest{
		Token: in.Token,
	})
	if err != nil {
		logx.Error("token不存在")
		return nil, errors.New("token不存在")
	}
	// 查询用户
	adminUserId := tokenRes.AdminToken.AdminUserId
	adminRes, err := l.svcCtx.AdminModel.FindOne(l.ctx, adminUserId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的用户")
			return nil, errors.New("不存在的用户")
		}
		logx.Error("用户信息获取失败" + err.Error())
		return nil, errors.New("用户信息获取失败" + err.Error())
	}

	return &admin.GetAdminProfileResponse{
		Admin: &admin.AdminModel{
			AdminUserId:   adminUserId,
			LoginUserName: adminRes.LoginUserName,
			LoginPassword: adminRes.LoginPassword,
			NickName:      adminRes.NickName,
			Locked:        adminRes.Locked,
		},
	}, nil
}
