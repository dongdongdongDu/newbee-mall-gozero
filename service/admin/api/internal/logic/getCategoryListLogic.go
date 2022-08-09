package logic

import (
	"context"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.GetCategoryListRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.GoodsCategoryRpc.GetCategoryList(l.ctx, &goodscategory.GetCategoryListRequest{
		CategoryLevel: req.CategoryLevel,
		ParentId:      req.ParentId,
		PageNumber:    req.PageNumber,
		PageSize:      req.PageSize,
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
		Data: &types.GetListResponse{
			List:       res.CategoryList,
			TotalCount: res.Total,
			TotalPage:  int64(math.Ceil(float64(res.Total) / float64(req.PageSize))),
			CurrPage:   req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}
