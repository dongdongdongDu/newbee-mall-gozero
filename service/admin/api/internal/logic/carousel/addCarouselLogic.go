package carousel

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/common/verify"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/carousel/rpc/carousel"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCarouselLogic {
	return &AddCarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCarouselLogic) AddCarousel(req *types.AddCarouselRequest) (resp *types.Response, err error) {
	// 校验输入格式
	if err := verify.Verify(*req, verify.CarouselAddParamVerify); err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        err.Error(),
		}, nil
	}
	// 获取当前用户
	adminToken, err := l.svcCtx.AdminTokenRpc.GetExistToken(l.ctx, &admintoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "创建失败！" + err.Error(),
		}, nil
	}
	_, err = l.svcCtx.CarouselRpc.AddCarousel(l.ctx, &carousel.AddCarouselRequest{
		CarouselUrl:  req.CarouselUrl,
		RedirectUrl:  req.RedirectUrl,
		CarouselRank: req.CarouselRank,
		User:         adminToken.AdminToken.AdminUserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "创建失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "创建成功",
		Data:       map[string]interface{}{},
	}, nil
}
