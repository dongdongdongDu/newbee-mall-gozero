package logic

import (
	"context"
	"newbee-mall-gozero/service/goods_category/model"
	"time"

	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_category/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCategoryLogic) AddCategory(in *goodscategory.AddCategoryRequest) (*goodscategory.EmptyResponse, error) {
	category := model.TbNewbeeMallGoodsCategory{
		CategoryLevel: in.Category.CategoryLevel,
		ParentId:      in.Category.ParentId,
		CategoryName:  in.Category.CategoryName,
		CategoryRank:  in.Category.CategoryRank,
		IsDeleted:     in.Category.IsDeleted,
		CreateTime:    time.Now(),
		CreateUser:    in.CreateUser,
		UpdateTime:    time.Now(),
	}
	_, err := l.svcCtx.GoodsCategoryModel.Insert(l.ctx, &category)
	if err != nil {
		logx.Error("添加失败：" + err.Error())
		return nil, err
	}
	return &goodscategory.EmptyResponse{}, nil
}
