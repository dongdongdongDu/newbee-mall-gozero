package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/common/sub"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"newbee-mall-gozero/service/index_config/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigForIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigForIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigForIndexLogic {
	return &GetConfigForIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigForIndexLogic) GetConfigForIndex(in *indexconfig.GetConfigForIndexRequest) (*indexconfig.GetConfigForIndexResponse, error) {
	limit := in.Num
	configList, err := l.svcCtx.ConfigModel.GetConfigForIndex(l.ctx, in.ConfigType, limit)
	if err != nil {
		logx.Error("Config列表获取失败" + err.Error())
		return nil, errors.New("Config列表获取失败" + err.Error())
	}
	// 获取商品id
	var ids []int64
	for _, config := range configList {
		ids = append(ids, config.GoodsId)
	}
	// 获取商品信息
	goodsList, err := l.svcCtx.GoodsInfoModel.GetForIndex(l.ctx, ids)
	if err != nil {
		logx.Error("商品列表获取失败" + err.Error())
		return nil, errors.New("商品列表获取失败" + err.Error())
	}
	// 构造列表（超过30个字符显示...）
	indexGoodsList := make([]*indexconfig.IndexGoods, 0)
	for _, goods := range goodsList {
		indexGoods := indexconfig.IndexGoods{
			GoodsId:       goods.GoodsId,
			GoodsName:     sub.SubStrLen(goods.GoodsName, 30),
			GoodsIntro:    sub.SubStrLen(goods.GoodsIntro, 30),
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
			Tag:           goods.Tag,
		}
		indexGoodsList = append(indexGoodsList, &indexGoods)
	}

	return &indexconfig.GetConfigForIndexResponse{
		IndexGoodsList: indexGoodsList,
	}, nil
}
