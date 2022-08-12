package carousel

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/carousel/rpc/carousel"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCarouselLogic {
	return &DeleteCarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCarouselLogic) DeleteCarousel(req *types.DeleteCarouselRequest) (resp *types.Response, err error) {
	// 删除
	_, err = l.svcCtx.CarouselRpc.DeleteCarousel(l.ctx, &carousel.DeleteCarouselRequest{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "删除失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "删除成功",
		Data:       map[string]interface{}{},
	}, nil
}
