package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryForSelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryForSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryForSelectLogic {
	return &GetCategoryForSelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryForSelectLogic) GetCategoryForSelect(req *types.GetCategoryForSelectRequest) (resp *types.Response, err error) {
	// 查找
	res, err := l.svcCtx.GoodsCategoryRpc.GetCategory(l.ctx, &goodscategory.GetCategoryRequest{
		Id: req.Id,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取失败！",
		}, nil
	}
	level := res.Category.CategoryLevel
	if level == 3 || level == 0 {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "参数异常",
		}, nil
	}
	categoryResult := make(map[string]interface{})
	if level == 1 {
		levelTwoList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
			Ids:           []int64{req.Id},
			CategoryLevel: 2,
		})
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取失败",
			}, nil
		}
		var secondLevelIds []int64
		for _, category := range levelTwoList.CategoryList {
			secondLevelIds = append(secondLevelIds, category.CategoryId)
		}

		levelThreeList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
			Ids:           secondLevelIds,
			CategoryLevel: 3,
		})
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取失败",
			}, nil
		}
		categoryResult["secondLevelCategories"] = levelTwoList.CategoryList
		categoryResult["thirdLevelCategories"] = levelThreeList.CategoryList

	}
	if level == 2 {
		levelThreeList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
			Ids:           []int64{req.Id},
			CategoryLevel: 3,
		})
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取失败",
			}, nil
		}
		categoryResult["thirdLevelCategories"] = levelThreeList.CategoryList
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data:       categoryResult,
	}, nil
}
