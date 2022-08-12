package admin

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"
	"newbee-mall-gozero/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminProfileLogic {
	return &GetAdminProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminProfileLogic) GetAdminProfile(req *types.GetAdminProfileRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.AdminRpc.GetAdminProfile(l.ctx, &admin.GetAdminProfileRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "未查询到记录",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetAdminProfileResponse{
			Admin: types.Admin{
				AdminUserId:   res.Admin.AdminUserId,
				LoginUserName: res.Admin.LoginUserName,
				LoginPassword: "******",
				NickName:      res.Admin.NickName,
				Locked:        res.Admin.Locked,
			},
		},
	}, nil
}
