package logic

import (
	"context"
	"errors"

	"newbee-mall-gozero/service/user_token/rpc/internal/svc"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *usertoken.DeleteTokenRequest) (*usertoken.DeleteTokenResponse, error) {
	// 查找token
	res, err := l.svcCtx.TokenModel.FindOneByToken(l.ctx, in.Token)
	if err != nil {
		logx.Error("查找token失败" + err.Error())
		return nil, errors.New("查找token失败" + err.Error())
	}
	// 删除
	err = l.svcCtx.TokenModel.Delete(l.ctx, res.UserId)
	if err != nil {
		logx.Error("删除token失败" + err.Error())
		return nil, errors.New("删除token失败" + err.Error())
	}

	return &usertoken.DeleteTokenResponse{}, nil
}
