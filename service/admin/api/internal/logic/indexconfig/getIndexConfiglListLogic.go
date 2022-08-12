package indexconfig

import (
	"context"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndexConfiglListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndexConfiglListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndexConfiglListLogic {
	return &GetIndexConfiglListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndexConfiglListLogic) GetIndexConfiglList(req *types.GetIndexConfigListRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.IndexConfigRpc.GetConfigList(l.ctx, &indexconfig.GetConfigListRequest{
		ConfigType: req.ConfigType,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: &types.GetListResponse{
			List:       res.ConfigList,
			TotalCount: res.Total,
			TotalPage:  int64(math.Ceil(float64(res.Total) / float64(req.PageSize))),
			CurrPage:   req.PageNumber,
			PageSize:   req.PageSize,
		},
	}, nil
}
