package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/goods_info/model"
	"time"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsInfoLogic {
	return &UpdateGoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsInfoLogic) UpdateGoodsInfo(in *goodsinfo.UpdateGoodsInfoRequest) (*goodsinfo.EmptyResponse, error) {
	// 查询商品
	res, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, in.GoodsInfo.GoodsId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的商品")
			return nil, errors.New("不存在的商品")
		}
		logx.Error("商品信息获取失败" + err.Error())
		return nil, errors.New("商品信息获取失败" + err.Error())
	}

	res.GoodsName = in.GoodsInfo.GoodsName
	res.GoodsIntro = in.GoodsInfo.GoodsIntro
	res.GoodsCategoryId = in.GoodsInfo.GoodsCategoryId
	res.GoodsCoverImg = in.GoodsInfo.GoodsCoverImg
	res.GoodsCarousel = in.GoodsInfo.GoodsCarousel
	res.GoodsDetailContent = in.GoodsInfo.GoodsDetailContent
	res.OriginalPrice = in.GoodsInfo.OriginalPrice
	res.SellingPrice = in.GoodsInfo.SellingPrice
	res.StockNum = in.GoodsInfo.StockNum
	res.Tag = in.GoodsInfo.Tag
	res.GoodsSellStatus = in.GoodsInfo.GoodsSellStatus
	res.UpdateUser = in.GoodsInfo.UpdateUser
	res.UpdateTime = time.Now()

	err = l.svcCtx.GoodsInfoModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("商品信息更新失败" + err.Error())
		return nil, errors.New("商品信息更新失败")
	}

	return &goodsinfo.EmptyResponse{}, nil
}
