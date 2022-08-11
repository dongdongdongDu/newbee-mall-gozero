package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/carousel/model"

	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/carousel/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCarouselLogic {
	return &DeleteCarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCarouselLogic) DeleteCarousel(in *carousel.DeleteCarouselRequest) (*carousel.EmptyResponse, error) {
	// 遍历Ids修改IsDeleted值
	for _, id := range in.Ids {
		// 查找
		res, err := l.svcCtx.CarouselModel.FindOne(l.ctx, id)
		if err != nil {
			if err == model.ErrNotFound {
				logx.Error("不存在的轮播图")
				return nil, errors.New("不存在的轮播图")
			}
			logx.Error("轮播图获取失败" + err.Error())
			return nil, errors.New("轮播图获取失败" + err.Error())
		}
		res.IsDeleted = 1
		err = l.svcCtx.CarouselModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("轮播图删除失败" + err.Error())
			return nil, errors.New("轮播图删除失败")
		}
	}

	return &carousel.EmptyResponse{}, nil
}
