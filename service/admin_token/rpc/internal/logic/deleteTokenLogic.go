package logic

import (
	"context"
	"errors"

	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/admin_token/rpc/internal/svc"

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

func (l *DeleteTokenLogic) DeleteToken(in *admintoken.DeleteTokenRequest) (*admintoken.DeleteTokenResponse, error) {
	// 查找token
	res, err := l.svcCtx.AdminTokenModel.FindOneByToken(l.ctx, in.Token)
	if err != nil {
		logx.Error("查找adminToken失败" + err.Error())
		return nil, errors.New("查找adminToken失败" + err.Error())
	}
	// 删除
	err = l.svcCtx.AdminTokenModel.Delete(l.ctx, res.AdminUserId)
	if err != nil {
		logx.Error("删除adminToken失败" + err.Error())
		return nil, errors.New("删除adminToken失败" + err.Error())
	}

	return &admintoken.DeleteTokenResponse{}, nil
}
