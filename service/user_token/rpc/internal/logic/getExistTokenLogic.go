package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"newbee-mall-gozero/service/user_token/rpc/internal/svc"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

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

func (l *GetExistTokenLogic) GetExistToken(in *usertoken.GetExistTokenRequest) (*usertoken.GetExistTokenResponse, error) {
	// 查找token
	res, err := l.svcCtx.TokenModel.FindOneByToken(l.ctx, in.Token)
	if err != nil {
		logx.Error("查找token失败" + err.Error())
		return nil, errors.New("查找token失败" + err.Error())
	}
	// 复制
	var tokenRes usertoken.Token
	_ = copier.Copy(&tokenRes, res)

	return &usertoken.GetExistTokenResponse{
		Token: &tokenRes,
	}, nil
}
