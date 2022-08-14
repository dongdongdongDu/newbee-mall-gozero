package order

import (
	"context"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"
	"time"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderByIdLogic) GetOrderById(req *types.GetOrderByIdRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.OrderRpc.GetOrderById(l.ctx, &order.GetOrderByIdRequest{
		Id: req.OrderId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！",
		}, nil
	}
	var orderItem []types.OrderItem
	err = copier.Copy(&orderItem, res.Order.OrderItems)
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetOrderByIdResponse{
			OrderId:                res.Order.OrderId,
			OrderNo:                res.Order.OrderNo,
			TotalPrice:             res.Order.TotalPrice,
			PayType:                res.Order.PayType,
			PayTypeString:          res.Order.PayTypeString,
			OrderStatus:            res.Order.OrderStatus,
			OrderStatusString:      res.Order.OrderStatusString,
			CreateTime:             time.Unix(res.Order.CreateTime, 0).Format("2006-01-02 15:04:05"),
			NewBeeMallOrderItemVOS: orderItem,
		},
	}, nil
}
