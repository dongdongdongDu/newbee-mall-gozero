package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/carousel/model"

	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/carousel/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCarouselLogic {
	return &GetCarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCarouselLogic) GetCarousel(in *carousel.GetCarouselRequest) (*carousel.GetCarouselResponse, error) {
	// 查找
	res, err := l.svcCtx.CarouselModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的轮播图")
			return nil, errors.New("不存在的轮播图")
		}
		logx.Error("轮播图获取失败" + err.Error())
		return nil, errors.New("轮播图获取失败" + err.Error())
	}
	// 复制
	//var c carousel.CarouselType
	//err = copier.Copy(&c, res)
	c := carousel.CarouselType{
		CarouselId:   res.CarouselId,
		CarouselUrl:  res.CarouselUrl,
		RedirectUrl:  res.RedirectUrl,
		CarouselRank: res.CarouselRank,
		IsDeleted:    res.IsDeleted,
		CreateTime:   res.CreateTime.Unix(),
		CreateUser:   res.CreateUser,
		UpdateTime:   res.UpdateTime.Unix(),
		UpdateUser:   res.UpdateUser,
	}

	//if err != nil {
	//	logx.Error("复制失败" + err.Error())
	//	return nil, errors.New("复制失败" + err.Error())
	//}

	return &carousel.GetCarouselResponse{
		Carousel: &c,
	}, nil
}
