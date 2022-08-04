package logic

import (
	"context"
	"errors"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsInfoLogic {
	return &DeleteGoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGoodsInfoLogic) DeleteGoodsInfo(in *goodsinfo.DeleteGoodsInfoRequest) (*goodsinfo.DeleteGoodsInfoResponse, error) {
	err := l.svcCtx.GoodsInfoModel.Delete(l.ctx, in.GoodsId)
	if err != nil {
		logx.Error("删除商品失败" + err.Error())
		return nil, errors.New("删除商品失败" + err.Error())
	}

	return &goodsinfo.DeleteGoodsInfoResponse{}, nil
}
