package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_category/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryListLogic) GetCategoryList(in *goodscategory.GetCategoryListRequest) (*goodscategory.GetCategoryListResponse, error) {
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	// 查找
	list, total, err := l.svcCtx.GoodsCategoryModel.GetCategoryListByParentIdAndLevel(l.ctx, in.ParentId, in.CategoryLevel, limit, offset)
	if err != nil {
		logx.Error("分类列表获取失败" + err.Error())
		return nil, errors.New("分类列表获取失败" + err.Error())
	}
	// 构造列表
	categoryList := make([]*goodscategory.Category, 0)
	for _, item := range list {
		var category goodscategory.Category
		err = copier.Copy(&category, item)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return nil, errors.New("复制失败" + err.Error())
		}
		category.CreateTime = item.CreateTime.Unix()
		category.UpdateTime = item.UpdateTime.Unix()
		categoryList = append(categoryList, &category)
	}

	return &goodscategory.GetCategoryListResponse{
		CategoryList: categoryList,
		Total:        total,
	}, nil
}
