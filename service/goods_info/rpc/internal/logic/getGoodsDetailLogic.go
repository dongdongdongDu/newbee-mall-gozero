package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsDetailLogic {
	return &GetGoodsDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsDetailLogic) GetGoodsDetail(in *goodsinfo.GetGoodsDetailRequest) (*goodsinfo.GetGoodsDetailResponse, error) {
	// 查找
	goodsInfo, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, in.GoodsId)
	if err != nil {
		logx.Error("查找商品信息失败" + err.Error())
		return nil, errors.New("查找商品信息失败" + err.Error())
	}
	// 判断状态
	if goodsInfo.GoodsSellStatus != 0 {
		logx.Error("商品已下架")
		return nil, errors.New("商品已下架")
	}
	var list []string
	list = append(list, goodsInfo.GoodsCarousel)
	// 复制
	var res goodsinfo.GetGoodsDetailResponse
	err = copier.Copy(&res, goodsInfo)
	if err != nil {
		logx.Error("copy失败" + err.Error())
		return nil, errors.New("copy失败" + err.Error())
	}

	res.GoodsCarouselList = list
	return &res, nil
}
