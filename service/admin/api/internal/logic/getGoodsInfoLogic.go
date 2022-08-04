package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"time"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsInfoLogic {
	return &GetGoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsInfoLogic) GetGoodsInfo(req *types.GetGoodsInfoRequest) (resp *types.Response, err error) {
	// 查找
	res, err := l.svcCtx.GoodsInfoRpc.GetGoodsInfo(l.ctx, &goodsinfo.GetGoodsInfoRequest{
		Id: req.GoodsId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！" + err.Error(),
		}, nil
	}
	goodsInfo := types.GoodsInfo{
		GoodsId:            res.GoodsInfo.GoodsId,
		GoodsName:          res.GoodsInfo.GoodsName,
		GoodsIntro:         res.GoodsInfo.GoodsIntro,
		GoodsCategoryId:    res.GoodsInfo.GoodsCategoryId,
		GoodsCoverImg:      res.GoodsInfo.GoodsCoverImg,
		GoodsCarousel:      res.GoodsInfo.GoodsCarousel,
		GoodsDetailContent: res.GoodsInfo.GoodsDetailContent,
		OriginalPrice:      res.GoodsInfo.OriginalPrice,
		SellingPrice:       res.GoodsInfo.SellingPrice,
		StockNum:           res.GoodsInfo.StockNum,
		Tag:                res.GoodsInfo.Tag,
		GoodsSellStatus:    res.GoodsInfo.GoodsSellStatus,
		CreateUser:         res.GoodsInfo.CreateUser,
		CreateTime:         time.Unix(res.GoodsInfo.CreateTime, 0).Format("2006-01-02 15:04:05"),
		UpdateUser:         res.GoodsInfo.UpdateUser,
		UpdateTime:         time.Unix(res.GoodsInfo.UpdateTime, 0).Format("2006-01-02 15:04:05"),
	}
	// todo:根据GoodsCategoryId查询商品分类

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: &types.GetGoodsInfoResponse{
			Goods:          goodsInfo,
			ThirdCategory:  types.GoodsCategory{},
			SecondCategory: types.GoodsCategory{},
			FirstCategory:  types.GoodsCategory{},
		},
	}, nil
}
