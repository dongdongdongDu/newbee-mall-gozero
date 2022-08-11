package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/goods_category/model"

	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_category/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryLogic) GetCategory(in *goodscategory.GetCategoryRequest) (*goodscategory.GetCategoryResponse, error) {
	// 查找
	res, err := l.svcCtx.GoodsCategoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的分类")
			return nil, errors.New("不存在的分类")
		}
		logx.Error("分类获取失败" + err.Error())
		return nil, errors.New("分类获取失败" + err.Error())
	}
	// 复制
	var category goodscategory.Category
	err = copier.Copy(&category, res)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}

	return &goodscategory.GetCategoryResponse{
		Category: &category,
	}, nil
}
