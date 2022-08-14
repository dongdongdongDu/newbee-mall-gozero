package order

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaySuccessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaySuccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaySuccessLogic {
	return &PaySuccessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaySuccessLogic) PaySuccess(req *types.PaySuccessRequest) (resp *types.Response, err error) {
	_, err = l.svcCtx.OrderRpc.PaySuccess(l.ctx, &order.PaySuccessRequest{
		OrderNo: req.OrderNo,
		PayType: req.PayType,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "订单支付失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "订单支付成功",
		Data:       map[string]interface{}{},
	}, nil
}
