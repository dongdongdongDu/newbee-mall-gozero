package carousel

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/carousel/rpc/carousel"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCarouselLogic {
	return &UpdateCarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCarouselLogic) UpdateCarousel(req *types.UpdateCarouselRequest) (resp *types.Response, err error) {
	// 获取当前用户
	adminToken, err := l.svcCtx.AdminTokenRpc.GetExistToken(l.ctx, &admintoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！" + err.Error(),
		}, nil
	}
	_, err = l.svcCtx.CarouselRpc.UpdateCarousel(l.ctx, &carousel.UpdateCarouselRequest{
		CarouselId:   req.CarouselId,
		CarouselUrl:  req.CarouselUrl,
		RedirectUrl:  req.RedirectUrl,
		CarouselRank: req.CarouselRank,
		User:         adminToken.AdminToken.AdminUserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新成功",
		Data:       map[string]interface{}{},
	}, nil
}
