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

type AddIndexConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIndexConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIndexConfigLogic {
	return &AddIndexConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIndexConfigLogic) AddIndexConfig(req *types.AddIndexConfigRequest) (resp *types.Response, err error) {
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
	_, err = l.svcCtx.IndexConfigRpc.AddConfig(l.ctx, &indexconfig.AddConfigRequest{
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
			Msg:        "创建失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "创建成功",
		Data:       map[string]interface{}{},
	}, nil
}