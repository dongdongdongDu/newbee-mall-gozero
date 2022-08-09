package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/goods_category/model"

	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_category/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *goodscategory.DeleteCategoryRequest) (*goodscategory.EmptyResponse, error) {
	// 遍历Ids修改IsDeleted值
	for _, id := range in.Ids {
		// 查找
		res, err := l.svcCtx.GoodsCategoryModel.FindOne(l.ctx, id)
		if err != nil {
			if err == model.ErrNotFound {
				logx.Error("不存在的分类")
				return nil, errors.New("不存在的分类")
			}
			logx.Error("分类获取失败" + err.Error())
			return nil, errors.New("分类获取失败" + err.Error())
		}
		res.IsDeleted = 1
		err = l.svcCtx.GoodsCategoryModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("分类删除失败" + err.Error())
			return nil, errors.New("分类删除失败")
		}
	}
	return &goodscategory.EmptyResponse{}, nil
}
