package indexconfig

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteIndexConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIndexConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIndexConfigLogic {
	return &DeleteIndexConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteIndexConfigLogic) DeleteIndexConfig(req *types.DeleteIndexConfigRequest) (resp *types.Response, err error) {
	// 删除
	_, err = l.svcCtx.IndexConfigRpc.DeleteConfig(l.ctx, &indexconfig.DeleteConfigRequest{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "删除失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "删除成功",
		Data:       map[string]interface{}{},
	}, nil
}
