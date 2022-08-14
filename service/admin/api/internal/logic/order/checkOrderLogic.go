package order

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOrderLogic {
	return &CheckOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckOrderLogic) CheckOrder(req *types.OrderRequest) (resp *types.Response, err error) {
	_, err = l.svcCtx.OrderRpc.CheckOrder(l.ctx, &order.CheckOrderRequest{
		Ids: req.Ids,
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
