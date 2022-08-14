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

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListLogic) GetOrderList(in *order.GetOrderListRequest) (*order.GetOrderListResponse, error) {
	// 查询
	orders, total, err := l.svcCtx.OrderModel.FindByUserAndStatus(l.ctx, in.PageNumber, in.UserId, in.Status)
	if err != nil {
		logx.Error("查询订单失败" + err.Error())
		return nil, errors.New("查询订单失败" + err.Error())
	}
	if total == 0 {
		logx.Error("查询订单失败 订单为空")
		return nil, errors.New("查询订单失败 订单为空")
	}
	// 返回订单列表
	var orderList []*order.OrderModel
	for _, od := range orders {
		var m order.OrderModel
		err = copier.Copy(&m, od)
		if err != nil {
			logx.Error("复制失败")
			return nil, errors.New("复制失败")
		}
		// 设置状态中文显示值
		m.CreateTime = od.CreateTime.Unix()
		m.PayTypeString = orderModel.GetPayTypeStr(od.PayType)
		m.OrderStatusString = orderModel.GetOrderStatusStr(od.OrderStatus)
		orderList = append(orderList, &m)
	}
	// 封装每个订单列表对象的订单项数据
	for _, m := range orderList {
		// 查找订单项
		items, err := l.svcCtx.OrderItemModel.FindItemsByOrderId(l.ctx, m.OrderId)
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
		m.OrderItems = orderItems
	}

	return &order.GetOrderListResponse{
		OrderList: orderList,
		Total:     total,
	}, nil
}
