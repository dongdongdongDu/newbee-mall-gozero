package category

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"time"

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
	if level == 1 {
		levelTwoList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
			Ids:           []int64{req.Id},
			CategoryLevel: 2,
			Limit:         10,
		})
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取失败",
			}, nil
		}
		logx.Info(len(levelTwoList.CategoryList))
		logx.Info(levelTwoList.CategoryList)
		var secondLevelIds []int64
		for _, category := range levelTwoList.CategoryList {
			secondLevelIds = append(secondLevelIds, category.CategoryId)
		}

		levelThreeList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
			Ids:           secondLevelIds,
			CategoryLevel: 3,
			Limit:         10,
		})
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取失败",
			}, nil
		}
		logx.Info(levelThreeList.CategoryList)
		var secondLevelCategories []*types.GoodsCategory
		for _, category := range levelTwoList.CategoryList {
			secondCategory := types.GoodsCategory{
				CategoryId:    category.CategoryId,
				CategoryLevel: category.CategoryLevel,
				ParentId:      category.ParentId,
				CategoryName:  category.CategoryName,
				CategoryRank:  category.CategoryRank,
				IsDeleted:     category.IsDeleted,
				CreateTime:    time.Unix(category.CreateTime, 0).Format("2006-01-02 15:04:05"),
				UpdateTime:    time.Unix(category.UpdateTime, 0).Format("2006-01-02 15:04:05"),
			}
			secondLevelCategories = append(secondLevelCategories, &secondCategory)
		}
		var thirdLevelCategories []*types.GoodsCategory
		for _, category := range levelThreeList.CategoryList {
			thirdCategory := types.GoodsCategory{
				CategoryId:    category.CategoryId,
				CategoryLevel: category.CategoryLevel,
				ParentId:      category.ParentId,
				CategoryName:  category.CategoryName,
				CategoryRank:  category.CategoryRank,
				IsDeleted:     category.IsDeleted,
				CreateTime:    time.Unix(category.CreateTime, 0).Format("2006-01-02 15:04:05"),
				UpdateTime:    time.Unix(category.UpdateTime, 0).Format("2006-01-02 15:04:05"),
			}
			thirdLevelCategories = append(thirdLevelCategories, &thirdCategory)
		}
		logx.Info(secondLevelCategories)
		logx.Info(thirdLevelCategories)
		return &types.Response{
			ResultCode: response.SUCCESS,
			Msg:        "SUCCESS",
			Data: types.GetCategoryForSelectLevel1Response{
				SecondLevelCategories: secondLevelCategories,
				ThirdLevelCategories:  thirdLevelCategories,
			},
		}, nil
	} else if level == 2 {
		levelThreeList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
			Ids:           []int64{req.Id},
			CategoryLevel: 3,
			Limit:         10,
		})
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取失败",
			}, nil
		}
		var thirdLevelCategories []*types.GoodsCategory
		for _, category := range levelThreeList.CategoryList {
			thirdCategory := types.GoodsCategory{
				CategoryId:    category.CategoryId,
				CategoryLevel: category.CategoryLevel,
				ParentId:      category.ParentId,
				CategoryName:  category.CategoryName,
				CategoryRank:  category.CategoryRank,
				IsDeleted:     category.IsDeleted,
				CreateTime:    time.Unix(category.CreateTime, 0).Format("2006-01-02 15:04:05"),
				UpdateTime:    time.Unix(category.UpdateTime, 0).Format("2006-01-02 15:04:05"),
			}
			thirdLevelCategories = append(thirdLevelCategories, &thirdCategory)
		}
		return &types.Response{
			ResultCode: response.SUCCESS,
			Msg:        "SUCCESS",
			Data: types.GetCategoryForSelectLevel2Response{
				ThirdLevelCategories: thirdLevelCategories,
			},
		}, nil
	} else {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "参数异常",
		}, nil
	}
}
