package admin

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/common/verify"
	"newbee-mall-gozero/service/admin/rpc/admin"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminPwdLogic {
	return &UpdateAdminPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminPwdLogic) UpdateAdminPwd(req *types.UpdateAdminPwdRequest) (resp *types.Response, err error) {
	// 校验输入格式
	if err := verify.Verify(*req, verify.GoodsAddParamVerify); err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        err.Error(),
		}, nil
	}
	// 更新信息
	_, err = l.svcCtx.AdminRpc.UpdateAdminPwd(l.ctx, &admin.UpdateAdminPwdRequest{
		Token:            req.Token,
		OriginalPassword: req.OriginalPassword,
		NewPassword:      req.NewPassword,
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
