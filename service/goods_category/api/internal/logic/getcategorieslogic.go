package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"

	"newbee-mall-gozero/service/goods_category/api/internal/svc"
	"newbee-mall-gozero/service/goods_category/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoriesLogic {
	return &GetCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoriesLogic) GetCategories() (resp *types.Response, err error) {
	//获取一级分类的固定数量的数据
	firstLevelList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
		Ids:           []int64{0},
		CategoryLevel: 1,
		Limit:         10,
	})
	if err != nil {
		logx.Error("一级分类列表获取失败" + err.Error())
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "一级分类列表获取失败" + err.Error(),
		}, nil
	}
	var firstLevelIds []int64
	for _, category := range firstLevelList.CategoryList {
		firstLevelIds = append(firstLevelIds, category.CategoryId)
	}
	//获取二级分类的数据
	secondLevelList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
		Ids:           firstLevelIds,
		CategoryLevel: 2,
		Limit:         10,
	})
	if err != nil {
		logx.Error("二级分类列表获取失败" + err.Error())
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "二级分类列表获取失败" + err.Error(),
		}, nil
	}
	var secondLevelIds []int64
	for _, category := range secondLevelList.CategoryList {
		secondLevelIds = append(secondLevelIds, category.CategoryId)
	}
	//获取三级分类的数据
	thirdLevelList, err := l.svcCtx.GoodsCategoryRpc.GetCategoryByParent(l.ctx, &goodscategory.GetCategoryByParentRequest{
		Ids:           secondLevelIds,
		CategoryLevel: 3,
		Limit:         10,
	})
	if err != nil {
		logx.Error("三级分类列表获取失败" + err.Error())
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "三级分类列表获取失败" + err.Error(),
		}, nil
	}
	thirdLevelCategoryMap := make(map[int64][]*goodscategory.Category)
	for _, thirdLevelCategory := range thirdLevelList.CategoryList {
		thirdLevelCategoryMap[thirdLevelCategory.ParentId] = []*goodscategory.Category{}
	}
	for k, v := range thirdLevelCategoryMap {
		for _, third := range thirdLevelList.CategoryList {
			if k == third.ParentId {
				v = append(v, third)
			}
			thirdLevelCategoryMap[k] = v
		}
	}

	var secondLevelCategoryVOS []types.SecondLevelCategoryVO
	//处理二级分类
	for _, secondLevelCategory := range secondLevelList.CategoryList {
		var secondLevelCategoryVO types.SecondLevelCategoryVO
		err = copier.Copy(&secondLevelCategoryVO, secondLevelCategory)
		if err != nil {
			logx.Error("复制出错")
			return &types.Response{
				ResultCode: response.SUCCESS,
				Msg:        "查询失败",
				Data:       map[string]interface{}{},
			}, nil
		}
		//如果该二级分类下有数据则放入 secondLevelCategoryVOS 对象中
		if _, ok := thirdLevelCategoryMap[secondLevelCategory.CategoryId]; ok {
			//根据二级分类的id取出thirdLevelCategoryMap分组中的三级分类list
			tempGoodsCategories := thirdLevelCategoryMap[secondLevelCategory.CategoryId]
			var thirdLevelCategoryRes []types.ThirdLevelCategoryVO
			err = copier.Copy(&thirdLevelCategoryRes, tempGoodsCategories)
			if err != nil {
				logx.Error("复制出错")
				return &types.Response{
					ResultCode: response.SUCCESS,
					Msg:        "查询失败",
					Data:       map[string]interface{}{},
				}, nil
			}
			secondLevelCategoryVO.ThirdLevelCategoryVOS = thirdLevelCategoryRes
			secondLevelCategoryVOS = append(secondLevelCategoryVOS, secondLevelCategoryVO)
		}

	}

	//处理一级分类
	//根据 parentId 将 thirdLevelCategories 分组
	secondLevelCategoryVOMap := make(map[int64][]types.SecondLevelCategoryVO)
	for _, secondLevelCategory := range secondLevelCategoryVOS {
		secondLevelCategoryVOMap[secondLevelCategory.ParentId] = []types.SecondLevelCategoryVO{}
	}
	for k, v := range secondLevelCategoryVOMap {
		for _, second := range secondLevelCategoryVOS {
			if k == second.ParentId {
				var secondLevelCategory types.SecondLevelCategoryVO
				err = copier.Copy(&secondLevelCategory, &second)
				if err != nil {
					logx.Error("复制出错")
					return &types.Response{
						ResultCode: response.SUCCESS,
						Msg:        "查询失败",
						Data:       map[string]interface{}{},
					}, nil
				}
				v = append(v, secondLevelCategory)
			}
			secondLevelCategoryVOMap[k] = v
		}
	}
	var newBeeMallIndexCategoryVOS []types.GetCategoriesForIndexResponse
	for _, firstCategory := range firstLevelList.CategoryList {
		var newBeeMallIndexCategoryVO types.GetCategoriesForIndexResponse
		err = copier.Copy(&newBeeMallIndexCategoryVO, &firstCategory)
		if err != nil {
			logx.Error("复制出错")
			return &types.Response{
				ResultCode: response.SUCCESS,
				Msg:        "查询失败",
				Data:       map[string]interface{}{},
			}, nil
		}
		//如果该一级分类下有数据则放入 newBeeMallIndexCategoryVOS 对象中
		if _, ok := secondLevelCategoryVOMap[firstCategory.CategoryId]; ok {
			//根据一级分类的id取出secondLevelCategoryVOMap分组中的二级级分类list
			tempGoodsCategories := secondLevelCategoryVOMap[firstCategory.CategoryId]
			newBeeMallIndexCategoryVO.SecondLevelCategoryVOS = tempGoodsCategories
			newBeeMallIndexCategoryVOS = append(newBeeMallIndexCategoryVOS, newBeeMallIndexCategoryVO)
		}
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data:       newBeeMallIndexCategoryVOS,
	}, nil
}
