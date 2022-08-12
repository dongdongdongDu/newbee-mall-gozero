package carousel

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"time"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCarouselLogic {
	return &GetCarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCarouselLogic) GetCarousel(req *types.GetCarouselRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.CarouselRpc.GetCarousel(l.ctx, &carousel.GetCarouselRequest{
		Id: req.Id,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！" + err.Error(),
		}, nil
	}
	c := types.Carousel{
		CarouselId:   res.Carousel.CarouselId,
		CarouselUrl:  res.Carousel.CarouselUrl,
		RedirectUrl:  res.Carousel.RedirectUrl,
		CarouselRank: res.Carousel.CarouselRank,
		IsDeleted:    res.Carousel.IsDeleted,
		CreateTime:   time.Unix(res.Carousel.CreateTime, 0).Format("2006-01-02 15:04:05"),
		CreateUser:   res.Carousel.CreateUser,
		UpdateTime:   time.Unix(res.Carousel.UpdateTime, 0).Format("2006-01-02 15:04:05"),
		UpdateUser:   res.Carousel.UpdateUser,
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetCarouselResponse{
			Carousel: c,
		},
	}, nil
}
