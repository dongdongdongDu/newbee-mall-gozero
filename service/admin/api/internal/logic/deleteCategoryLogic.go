package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategoryRequest) (resp *types.Response, err error) {
	_, err = l.svcCtx.GoodsCategoryRpc.DeleteCategory(l.ctx, &goodscategory.DeleteCategoryRequest{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "删除失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "删除失败",
		Data:       map[string]interface{}{},
	}, nil
}
