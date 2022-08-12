package admin

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"
	"newbee-mall-gozero/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminNameLogic {
	return &UpdateAdminNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminNameLogic) UpdateAdminName(req *types.UpdateAdminNameRequest) (resp *types.Response, err error) {
	// 更新信息
	_, err = l.svcCtx.AdminRpc.UpdateAdminName(l.ctx, &admin.UpdateAdminNameRequest{
		Token:         req.Token,
		LoginUserName: req.LoginUserName,
		NickName:      req.NickName,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新成功",
		Data:       map[string]interface{}{},
	}, nil
}
