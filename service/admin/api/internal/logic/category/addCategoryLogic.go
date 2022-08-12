package category

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/common/verify"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.AddCategoryRequest) (resp *types.Response, err error) {
	// 校验输入格式
	if err := verify.Verify(*req, verify.AddCategoryVerify); err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        err.Error(),
		}, nil
	}
	// 获取当前用户
	adminToken, err := l.svcCtx.AdminTokenRpc.GetExistToken(l.ctx, &admintoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "创建失败！" + err.Error(),
		}, nil
	}
	// 创建
	_, err = l.svcCtx.GoodsCategoryRpc.AddCategory(l.ctx, &goodscategory.AddCategoryRequest{
		Category: &goodscategory.Category{
			CategoryLevel: req.CategoryLevel,
			ParentId:      req.ParentId,
			CategoryName:  req.CategoryName,
			CategoryRank:  req.CategoryRank,
			IsDeleted:     req.IsDeleted,
		},
		CreateUser: adminToken.AdminToken.AdminUserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "创建失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "创建成功",
		Data:       map[string]interface{}{},
	}, nil
}
