package logic

import (
	"context"

	"newbee-mall-gozero/service/user/rpc/internal/svc"
	"newbee-mall-gozero/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *user.GenerateTokenRequest) (*user.GenerateTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GenerateTokenResponse{}, nil
}
