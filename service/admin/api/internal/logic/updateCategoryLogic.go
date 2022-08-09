package logic

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"strconv"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryRequest) (resp *types.Response, err error) {
	// 获取当前用户
	adminToken, err := l.svcCtx.AdminTokenRpc.GetExistToken(l.ctx, &admintoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！" + err.Error(),
		}, nil
	}
	rank, err := strconv.Atoi(req.CategoryRank)
	if err != nil {
		logx.Error("string 转 int 失败")
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！" + err.Error(),
		}, nil
	}
	// 更新
	_, err = l.svcCtx.GoodsCategoryRpc.UpdateCategory(l.ctx, &goodscategory.UpdateCategoryRequest{
		Category: &goodscategory.Category{
			CategoryLevel: req.CategoryLevel,
			ParentId:      req.ParentId,
			CategoryName:  req.CategoryName,
			CategoryRank:  int64(rank),
			IsDeleted:     req.IsDeleted,
		},
		UpdateUser: adminToken.AdminToken.AdminUserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新成功",
		Data:       map[string]interface{}{},
	}, nil
}
