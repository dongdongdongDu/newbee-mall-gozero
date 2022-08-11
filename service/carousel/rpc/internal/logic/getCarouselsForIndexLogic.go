package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/carousel/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCarouselsForIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCarouselsForIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCarouselsForIndexLogic {
	return &GetCarouselsForIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCarouselsForIndexLogic) GetCarouselsForIndex(in *carousel.GetCarouselsForIndexRequest) (*carousel.GetCarouselsForIndexResponse, error) {
	res, err := l.svcCtx.CarouselModel.GetForIndex(l.ctx, in.Num)
	if err != nil {
		logx.Error("轮播图获取失败" + err.Error())
		return nil, errors.New("轮播图获取失败" + err.Error())
	}
	// 构造列表
	carouselList := make([]*carousel.CarouselsForIndex, 0)
	for _, item := range res {
		var c carousel.CarouselsForIndex
		err = copier.Copy(&c, item)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return nil, errors.New("复制失败" + err.Error())
		}
		carouselList = append(carouselList, &c)
	}

	return &carousel.GetCarouselsForIndexResponse{
		CarouselIndex: carouselList,
	}, nil
}
