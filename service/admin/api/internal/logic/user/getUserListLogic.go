package user

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/rpc/admin"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.AdminRpc.GetUserList(l.ctx, &admin.GetUserListRequest{
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取失败",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "获取成功",
		Data: types.GetListResponse{
			List:       res.List,
			TotalCount: res.TotalCount,
			TotalPage:  res.TotalPage,
			CurrPage:   res.CurrPage,
			PageSize:   res.PageSize,
		},
	}, nil
}
