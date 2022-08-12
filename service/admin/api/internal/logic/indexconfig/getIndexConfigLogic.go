package indexconfig

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"time"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndexConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndexConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndexConfigLogic {
	return &GetIndexConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndexConfigLogic) GetIndexConfig(req *types.GetIndexConfigRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.IndexConfigRpc.GetConfig(l.ctx, &indexconfig.GetConfigRequest{
		Id: req.Id,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！" + err.Error(),
		}, nil
	}

	config := types.IndexConfig{
		ConfigId:    res.Config.ConfigId,
		ConfigName:  res.Config.ConfigName,
		ConfigType:  res.Config.ConfigType,
		GoodsId:     res.Config.GoodsId,
		RedirectUrl: res.Config.RedirectUrl,
		ConfigRank:  res.Config.ConfigRank,
		IsDeleted:   res.Config.IsDeleted,
		CreateTime:  time.Unix(res.Config.CreateTime, 0).Format("2006-01-02 15:04:05"),
		CreateUser:  res.Config.CreateUser,
		UpdateTime:  time.Unix(res.Config.UpdateTime, 0).Format("2006-01-02 15:04:05"),
		UpdateUser:  res.Config.UpdateUser,
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetIndexConfigResponse{
			IndexConfig: config,
		},
	}, nil
}
