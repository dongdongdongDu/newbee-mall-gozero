package logic

import (
	"context"
	"errors"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlterGoodsStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlterGoodsStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlterGoodsStatusLogic {
	return &AlterGoodsStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlterGoodsStatusLogic) AlterGoodsStatus(in *goodsinfo.AlterGoodsStatusRequest) (*goodsinfo.EmptyResponse, error) {
	if in.Status != 0 && in.Status != 1 {
		logx.Error("非法操作")
		return nil, errors.New("非法操作")
	}

	for _, id := range in.Ids {
		// 查找
		goodsRes, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, id)
		if err != nil {
			logx.Error("不存在的商品id")
			return nil, errors.New("不存在的商品id")
		}
		// 更新状态
		goodsRes.GoodsSellStatus = in.Status
		err = l.svcCtx.GoodsInfoModel.Update(l.ctx, goodsRes)
		if err != nil {
			logx.Error("商品状态更新失败" + err.Error())
			return nil, errors.New("商品状态更新失败" + err.Error())
		}
	}

	return &goodsinfo.EmptyResponse{}, nil
}
