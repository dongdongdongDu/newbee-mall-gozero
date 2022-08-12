package carousel

import (
	"context"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/carousel/rpc/carousel"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCarouselListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCarouselListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCarouselListLogic {
	return &GetCarouselListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCarouselListLogic) GetCarouselList(req *types.GetCarouselListRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.CarouselRpc.GetCarouselList(l.ctx, &carousel.GetCarouselListRequest{
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: &types.GetListResponse{
			List:       res.CarouselList,
			TotalCount: res.Total,
			TotalPage:  int64(math.Ceil(float64(res.Total) / float64(req.PageSize))),
			CurrPage:   req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}
