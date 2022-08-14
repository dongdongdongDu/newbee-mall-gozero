package user

import (
	"context"
	"errors"

	orderModel "newbee-mall-gozero/service/order/model/order"
	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(in *order.OrderRequest) (*order.GetOrderDetailResponse, error) {
	// 查找订单
	res, err := l.svcCtx.OrderModel.FindOneByNo(l.ctx, in.OrderNo)
	if err != nil {
		logx.Error("查询订单失败" + err.Error())
		return nil, errors.New("查询订单失败" + err.Error())
	}
	if res.UserId != in.UserId {
		logx.Error("userId不相等")
		return nil, errors.New("禁止该操作")
	}
	// 查找订单项
	items, err := l.svcCtx.OrderItemModel.FindItemsByOrderId(l.ctx, res.OrderId)
	if err != nil {
		logx.Error("查询订单项失败" + err.Error())
		return nil, errors.New("查询订单项失败" + err.Error())
	}
	if len(items) == 0 {
		logx.Error("订单项不存在")
		return nil, errors.New("订单项不存在")
	}
	var orderItems []*order.OrderItem
	err = copier.Copy(&orderItems, items)
	if err != nil {
		logx.Error("复制失败")
		return nil, errors.New("复制失败")
	}

	return &order.GetOrderDetailResponse{
		OrderNo:           res.OrderNo,
		TotalPrice:        res.TotalPrice,
		PayStatus:         res.PayStatus,
		PayType:           res.PayType,
		PayTypeString:     orderModel.GetPayTypeStr(res.PayType),
		PayTime:           res.PayTime.Time.Unix(),
		OrderStatus:       res.OrderStatus,
		OrderStatusString: orderModel.GetOrderStatusStr(res.OrderStatus),
		CreateTime:        res.CreateTime.Unix(),
		OrderItems:        orderItems,
	}, nil
}
