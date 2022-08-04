package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGoodsLogic) SearchGoods(in *goodsinfo.SearchGoodsRequest) (*goodsinfo.SearchGoodsResponse, error) {
	// 查找
	goodsInfo, total, err := l.svcCtx.GoodsInfoModel.Search(l.ctx, in.PageNumber, in.GoodsCategoryId, in.Keyword, in.OrderBy)
	if err != nil {
		logx.Error("搜索商品失败" + err.Error())
		return nil, errors.New("搜索商品失败" + err.Error())
	}
	goodsList := make([]*goodsinfo.SearchGoods, 0)
	for _, goods := range goodsInfo {
		goodsList = append(goodsList, &goodsinfo.SearchGoods{
			GoodsId:       goods.GoodsId,
			GoodsName:     goods.GoodsName,
			GoodsIntro:    goods.GoodsIntro,
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
		})
	}

	return &goodsinfo.SearchGoodsResponse{
		SearchGoodsList: goodsList,
		Total:           total,
	}, nil
}
