package goodsInfo

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsDetailLogic {
	return &GetGoodsDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsDetailLogic) GetGoodsDetail(req *types.GetGoodsDetailRequest) (resp *types.Response, err error) {
	// 获取商品详情
	res, err := l.svcCtx.GoodsInfoRpc.GetGoodsDetail(l.ctx, &goodsinfo.GetGoodsDetailRequest{
		GoodsId: req.Id,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetGoodsDetailResponse{
			GoodsId:            res.GoodsId,
			GoodsName:          res.GoodsName,
			GoodsIntro:         res.GoodsIntro,
			GoodsCoverImg:      res.GoodsCoverImg,
			SellingPrice:       res.SellingPrice,
			GoodsDetailContent: res.GoodsDetailContent,
			OriginalPrice:      res.OriginalPrice,
			Tag:                res.Tag,
			GoodsCarouselList:  res.GoodsCarouselList,
		},
	}, nil
}
