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

type AddGoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsInfoLogic {
	return &AddGoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGoodsInfoLogic) AddGoodsInfo(in *goodsinfo.AddGoodsInfoRequest) (*goodsinfo.EmptyResponse, error) {
	// 查询goodsCategory
	goodsCategory, err := l.svcCtx.GoodsCategoryModel.FindOne(l.ctx, in.GoodsInfo.GoodsCategoryId)
	if err != nil {
		logx.Error("查找商品分类失败" + err.Error())
		return nil, errors.New("查找商品分类失败" + err.Error())
	}
	// 判断goodsCategory状态
	if goodsCategory.IsDeleted != 0 {
		logx.Error("分类数据异常")
		return nil, errors.New("查找商品分类失败")
	}
	if goodsCategory.CategoryLevel > 3 {
		logx.Error("分类数据异常")
		return nil, errors.New("分类数据异常")
	}
	goodsInfo := model.TbNewbeeMallGoodsInfo{
		GoodsName:          in.GoodsInfo.GoodsName,
		GoodsIntro:         in.GoodsInfo.GoodsIntro,
		GoodsCategoryId:    in.GoodsInfo.GoodsCategoryId,
		GoodsCoverImg:      in.GoodsInfo.GoodsCoverImg,
		GoodsCarousel:      in.GoodsInfo.GoodsCarousel,
		GoodsDetailContent: in.GoodsInfo.GoodsDetailContent,
		OriginalPrice:      in.GoodsInfo.OriginalPrice,
		SellingPrice:       in.GoodsInfo.SellingPrice,
		StockNum:           in.GoodsInfo.StockNum,
		Tag:                in.GoodsInfo.Tag,
		GoodsSellStatus:    in.GoodsInfo.GoodsSellStatus,
		CreateUser:         in.GoodsInfo.CreateUser,
		CreateTime:         time.Now(),
		UpdateUser:         in.GoodsInfo.UpdateUser,
		UpdateTime:         time.Now(),
	}
	_, err = l.svcCtx.GoodsInfoModel.Insert(l.ctx, &goodsInfo)
	if err != nil {
		logx.Error("添加失败：" + err.Error())
		return nil, err
	}

	return &goodsinfo.EmptyResponse{}, nil
}
