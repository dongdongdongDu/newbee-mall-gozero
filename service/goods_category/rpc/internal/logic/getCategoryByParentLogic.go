package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_category/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByParentLogic {
	return &GetCategoryByParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByParentLogic) GetCategoryByParent(in *goodscategory.GetCategoryByParentRequest) (*goodscategory.GetCategoryByParentResponse, error) {
	// 查找
	categories, err := l.svcCtx.GoodsCategoryModel.GetCategoriesByParentIdsAndLevel(l.ctx, in.Ids, in.CategoryLevel, in.Limit)
	if err != nil {
		logx.Error("分类列表获取失败" + err.Error())
		return nil, errors.New("分类列表获取失败" + err.Error())
	}
	var categoryList []*goodscategory.Category
	for _, item := range categories {
		//var category goodscategory.Category
		//err = copier.Copy(&category, item)
		//if err != nil {
		//	logx.Error("分类复制失败" + err.Error())
		//	return nil, errors.New("分类复制失败" + err.Error())
		//}
		category := goodscategory.Category{
			CategoryId:    item.CategoryId,
			CategoryLevel: item.CategoryLevel,
			ParentId:      item.ParentId,
			CategoryName:  item.CategoryName,
			CategoryRank:  item.CategoryRank,
			IsDeleted:     item.IsDeleted,
			CreateTime:    item.CreateTime.Unix(),
			UpdateTime:    item.UpdateTime.Unix(),
		}
		categoryList = append(categoryList, &category)
	}

	return &goodscategory.GetCategoryByParentResponse{
		CategoryList: categoryList,
	}, nil
}
