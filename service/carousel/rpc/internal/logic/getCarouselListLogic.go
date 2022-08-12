package logic

import (
	"context"
	"errors"

	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/carousel/rpc/internal/svc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCarouselListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCarouselListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCarouselListLogic {
	return &GetCarouselListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCarouselListLogic) GetCarouselList(in *carousel.GetCarouselListRequest) (*carousel.GetCarouselListResponse, error) {
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	// 查找
	list, err := l.svcCtx.CarouselModel.TakePart(l.ctx, limit, offset)
	if err != nil {
		logx.Error("轮播图列表获取失败" + err.Error())
		return nil, errors.New("轮播图列表获取失败" + err.Error())
	}
	// 构造列表
	carouselList := make([]*carousel.CarouselType, 0)
	for _, item := range list {
		var c carousel.CarouselType
		err = copier.Copy(&c, item)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return nil, errors.New("复制失败" + err.Error())
		}
		c.CreateTime = item.CreateTime.Unix()
		c.UpdateTime = item.UpdateTime.Unix()
		carouselList = append(carouselList, &c)
	}
	// 查询总数
	total, err := l.svcCtx.CarouselModel.CountAll(l.ctx)
	if err != nil {
		logx.Error("轮播图总数获取失败" + err.Error())
		return nil, errors.New("轮播图总数获取失败" + err.Error())
	}

	return &carousel.GetCarouselListResponse{
		CarouselList: carouselList,
		Total:        total,
	}, nil
}
