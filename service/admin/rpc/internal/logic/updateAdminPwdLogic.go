package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/admin/model"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"

	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminPwdLogic {
	return &UpdateAdminPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminPwdLogic) UpdateAdminPwd(in *admin.UpdateAdminPwdRequest) (*admin.EmptyResponse, error) {
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
	res, err := l.svcCtx.AdminModel.FindOne(l.ctx, adminUserId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的用户")
			return nil, errors.New("不存在的用户")
		}
		logx.Error("用户信息获取失败" + err.Error())
		return nil, errors.New("用户信息获取失败" + err.Error())
	}
	// 判断原密码
	if res.LoginPassword != in.OriginalPassword {
		logx.Error("原密码不正确")
		return nil, errors.New("原密码不正确")
	}
	res.LoginPassword = in.NewPassword

	// 写入更新
	err = l.svcCtx.AdminModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("用户信息更新失败" + err.Error())
		return nil, errors.New("用户信息更新失败" + err.Error())
	}

	return &admin.EmptyResponse{}, nil
}
