package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/admin_token/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExistTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExistTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExistTokenLogic {
	return &GetExistTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExistTokenLogic) GetExistToken(in *admintoken.GetExistTokenRequest) (*admintoken.GetExistTokenResponse, error) {
	// 查找adminToken
	res, err := l.svcCtx.AdminTokenModel.FindOneByToken(l.ctx, in.Token)
	if err != nil {
		logx.Error("查找adminToken失败" + err.Error())
		return nil, errors.New("查找adminToken失败" + err.Error())
	}

	return &admintoken.GetExistTokenResponse{
		AdminToken: &admintoken.AdminToken{
			AdminUserId: res.AdminUserId,
			Token:       res.Token,
			UpdateTime:  res.UpdateTime.Unix(),
			ExpireTime:  res.ExpireTime.Unix(),
		},
	}, nil
}
