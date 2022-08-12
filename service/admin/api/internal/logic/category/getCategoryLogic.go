package category

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.GetCategoryRequest) (resp *types.Response, err error) {
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

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: map[string]interface{}{
			"mallGoodsCategory": res.Category,
		},
	}, nil
}
