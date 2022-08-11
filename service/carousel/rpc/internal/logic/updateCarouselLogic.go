package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/carousel/model"
	"time"

	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/carousel/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCarouselLogic {
	return &UpdateCarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCarouselLogic) UpdateCarousel(in *carousel.UpdateCarouselRequest) (*carousel.EmptyResponse, error) {
	// 查找
	res, err := l.svcCtx.CarouselModel.FindOne(l.ctx, in.CarouselId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的轮播图")
			return nil, errors.New("不存在的轮播图")
		}
		logx.Error("轮播图获取失败" + err.Error())
		return nil, errors.New("轮播图获取失败" + err.Error())
	}
	res.CarouselUrl = in.CarouselUrl
	res.RedirectUrl = in.RedirectUrl
	res.CarouselRank = in.CarouselRank
	res.UpdateTime = time.Now()
	res.UpdateUser = in.User

	err = l.svcCtx.CarouselModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("轮播图更新失败" + err.Error())
		return nil, errors.New("轮播图更新失败")
	}

	return &carousel.EmptyResponse{}, nil
}
