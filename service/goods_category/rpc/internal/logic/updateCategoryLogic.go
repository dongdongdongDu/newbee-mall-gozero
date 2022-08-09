package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/goods_category/model"
	"time"

	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_category/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *goodscategory.UpdateCategoryRequest) (*goodscategory.EmptyResponse, error) {
	// 查找
	res, err := l.svcCtx.GoodsCategoryModel.FindOne(l.ctx, in.Category.CategoryId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的分类")
			return nil, errors.New("不存在的分类")
		}
		logx.Error("分类获取失败" + err.Error())
		return nil, errors.New("分类获取失败" + err.Error())
	}
	res.CategoryName = in.Category.CategoryName
	res.CategoryRank = in.Category.CategoryRank
	res.UpdateUser = in.UpdateUser
	res.UpdateTime = time.Now()

	err = l.svcCtx.GoodsCategoryModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("分类更新失败" + err.Error())
		return nil, errors.New("分类更新失败")
	}

	return &goodscategory.EmptyResponse{}, nil
}
