package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"

	"newbee-mall-gozero/service/carousel/api/internal/svc"
	"newbee-mall-gozero/service/carousel/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndexInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndexInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndexInfoLogic {
	return &GetIndexInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndexInfoLogic) GetIndexInfo() (resp *types.Response, err error) {
	carouselRes, err := l.svcCtx.CarouselRpc.GetCarouselsForIndex(l.ctx, &carousel.GetCarouselsForIndexRequest{
		Num: 5,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "轮播图获取失败" + err.Error(),
		}, nil
	}
	hotGoodsRes, err := l.svcCtx.IndexConfigRpc.GetConfigForIndex(l.ctx, &indexconfig.GetConfigForIndexRequest{
		ConfigType: 3,
		Num:        4,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "热销商品获取失败" + err.Error(),
		}, nil
	}
	newGoodsRes, err := l.svcCtx.IndexConfigRpc.GetConfigForIndex(l.ctx, &indexconfig.GetConfigForIndexRequest{
		ConfigType: 4,
		Num:        5,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "新品获取失败" + err.Error(),
		}, nil
	}
	recommendGoodsRes, err := l.svcCtx.IndexConfigRpc.GetConfigForIndex(l.ctx, &indexconfig.GetConfigForIndexRequest{
		ConfigType: 5,
		Num:        10,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "推荐商品获取失败" + err.Error(),
		}, nil
	}
	// 构造列表
	var carouselList []*types.Carousel
	for _, item := range carouselRes.CarouselIndex {
		c := &types.Carousel{
			CarouselUrl: item.CarouselUrl,
			RedirectUrl: item.RedirectUrl,
		}
		carouselList = append(carouselList, c)
	}
	var hotGoodsList []*types.IndexGoods
	for _, goods := range hotGoodsRes.IndexGoodsList {
		hotGoods := &types.IndexGoods{
			GoodsId:       goods.GoodsId,
			GoodsName:     goods.GoodsName,
			GoodsIntro:    goods.GoodsIntro,
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
			Tag:           goods.Tag,
		}
		hotGoodsList = append(hotGoodsList, hotGoods)
	}
	var newGoodsList []*types.IndexGoods
	for _, goods := range newGoodsRes.IndexGoodsList {
		newGoods := &types.IndexGoods{
			GoodsId:       goods.GoodsId,
			GoodsName:     goods.GoodsName,
			GoodsIntro:    goods.GoodsIntro,
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
			Tag:           goods.Tag,
		}
		newGoodsList = append(newGoodsList, newGoods)
	}
	var recommendGoodsList []*types.IndexGoods
	for _, goods := range recommendGoodsRes.IndexGoodsList {
		recommendGoods := &types.IndexGoods{
			GoodsId:       goods.GoodsId,
			GoodsName:     goods.GoodsName,
			GoodsIntro:    goods.GoodsIntro,
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
			Tag:           goods.Tag,
		}
		recommendGoodsList = append(recommendGoodsList, recommendGoods)
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetIndexInfoResponse{
			Carousels:      carouselList,
			HotGoods:       hotGoodsList,
			NewGoods:       newGoodsList,
			RecommendGoods: recommendGoodsList,
		},
	}, nil
}
