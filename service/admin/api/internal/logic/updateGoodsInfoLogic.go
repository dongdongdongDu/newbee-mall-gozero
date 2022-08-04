package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsInfoLogic {
	return &UpdateGoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodsInfoLogic) UpdateGoodsInfo(req *types.UpdateGoodsInfoRequest) (resp *types.Response, err error) {
	// 更新
	_, err = l.svcCtx.GoodsInfoRpc.UpdateGoodsInfo(l.ctx, &goodsinfo.UpdateGoodsInfoRequest{GoodsInfo: &goodsinfo.GoodsInfo{
		GoodsId:            req.GoodsId,
		GoodsName:          req.GoodsName,
		GoodsIntro:         req.GoodsIntro,
		GoodsCategoryId:    req.GoodsId,
		GoodsCoverImg:      req.GoodsCoverImg,
		GoodsCarousel:      req.GoodsCarousel,
		GoodsDetailContent: req.GoodsDetailContent,
		OriginalPrice:      req.OriginalPrice,
		SellingPrice:       req.SellingPrice,
		StockNum:           req.StockNum,
		Tag:                req.Tag,
		GoodsSellStatus:    req.GoodsSellStatus,
	}})

	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新成功",
		Data:       map[string]interface{}{},
	}, nil
}
