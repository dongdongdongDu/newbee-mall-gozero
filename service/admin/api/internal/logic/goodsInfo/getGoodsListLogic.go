package goodsInfo

import (
	"context"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsListLogic) GetGoodsList(req *types.GetGoodsListRequest) (resp *types.Response, err error) {
	// 查找
	res, err := l.svcCtx.GoodsInfoRpc.GetGoodsList(l.ctx, &goodsinfo.GetGoodsListRequest{
		PageNumber:  req.PageNumber,
		PageSize:    req.PageSize,
		GoodsName:   req.GoodsName,
		GoodsStatus: req.GoodsSellStatus,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "获取成功",
		Data: types.GetListResponse{
			List:       res.GoodsInfo,
			TotalCount: res.Total,
			TotalPage:  int64(math.Ceil(float64(res.Total) / float64(req.PageSize))),
			CurrPage:   req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}
