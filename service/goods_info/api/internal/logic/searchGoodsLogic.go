package logic

import (
	"context"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/api/internal/svc"
	"newbee-mall-gozero/service/goods_info/api/internal/types"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchGoodsLogic) SearchGoods(req *types.SearchGoodsRequest) (resp *types.Response, err error) {
	// 搜索
	res, err := l.svcCtx.GoodsInfoRpc.SearchGoods(l.ctx, &goodsinfo.SearchGoodsRequest{
		PageNumber:      req.PageNumber,
		GoodsCategoryId: req.GoodsCategoryId,
		Keyword:         req.Keyword,
		OrderBy:         req.OrderBy,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "获取成功",
		Data: types.SearchGoodsResponse{
			List:       res.SearchGoodsList,
			TotalCount: res.Total,
			TotalPage:  int64(math.Ceil(float64(res.Total) / 10)),
			CurrPage:   req.PageNumber,
			PageSize:   10,
		},
	}, nil
}
