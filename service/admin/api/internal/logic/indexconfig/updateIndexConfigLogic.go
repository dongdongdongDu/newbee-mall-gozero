package indexconfig

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIndexConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIndexConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIndexConfigLogic {
	return &UpdateIndexConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIndexConfigLogic) UpdateIndexConfig(req *types.UpdateIndexConfigRequest) (resp *types.Response, err error) {
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
	_, err = l.svcCtx.IndexConfigRpc.UpdateConfig(l.ctx, &indexconfig.UpdateConfigRequest{
		ConfigId:    req.ConfigId,
		ConfigName:  req.ConfigName,
		ConfigType:  req.ConfigType,
		GoodsId:     req.GoodsId,
		RedirectUrl: req.RedirectUrl,
		ConfigRank:  req.ConfigRank,
		User:        adminToken.AdminToken.AdminUserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新成功",
		Data:       map[string]interface{}{},
	}, nil
}
