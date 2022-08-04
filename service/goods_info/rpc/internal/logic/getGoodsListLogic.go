package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsListLogic) GetGoodsList(in *goodsinfo.GetGoodsListRequest) (*goodsinfo.GetGoodsListResponse, error) {
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	// 查找
	goodsRes, err := l.svcCtx.GoodsInfoModel.Find(l.ctx, in.GoodsName, in.GoodsStatus, limit, offset)
	if err != nil {
		logx.Error("商品列表获取失败" + err.Error())
		return nil, errors.New("商品列表获取失败" + err.Error())
	}
	// 构造列表
	goodsList := make([]*goodsinfo.GoodsInfo, 0)
	for _, goods := range goodsRes {
		var goodsInfo goodsinfo.GoodsInfo
		err = copier.Copy(&goodsInfo, goods)
		goodsList = append(goodsList, &goodsInfo)
	}
	total := len(goodsList)

	return &goodsinfo.GetGoodsListResponse{
		GoodsInfo: goodsList,
		Total:     int64(total),
	}, nil
}
