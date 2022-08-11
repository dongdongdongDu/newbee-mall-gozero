package logic

import (
	"context"
	"newbee-mall-gozero/service/carousel/model"
	"time"

	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/carousel/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCarouselLogic {
	return &AddCarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCarouselLogic) AddCarousel(in *carousel.AddCarouselRequest) (*carousel.EmptyResponse, error) {
	// 构造新对象
	newCarousel := model.TbNewbeeMallCarousel{
		CarouselUrl:  in.CarouselUrl,
		RedirectUrl:  in.RedirectUrl,
		CarouselRank: in.CarouselRank,
		IsDeleted:    0,
		CreateTime:   time.Now(),
		CreateUser:   in.User,
		UpdateTime:   time.Now(),
		UpdateUser:   in.User,
	}
	_, err := l.svcCtx.CarouselModel.Insert(l.ctx, &newCarousel)
	if err != nil {
		logx.Error("添加失败：" + err.Error())
		return nil, err
	}

	return &carousel.EmptyResponse{}, nil
}
