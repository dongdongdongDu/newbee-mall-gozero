package logic

import (
	"context"
	"errors"
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

	return &usertoken.GetExistTokenResponse{
		Token: &usertoken.Token{
			UserId:     res.UserId,
			Token:      res.Token,
			UpdateTime: res.UpdateTime.Unix(),
			ExpireTime: res.ExpireTime.Unix(),
		},
	}, nil
}
