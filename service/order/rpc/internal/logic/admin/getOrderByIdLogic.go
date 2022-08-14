package admin

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	orderModel "newbee-mall-gozero/service/order/model/order"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *order.GetOrderByIdRequest) (*order.GetOrderByIdResponse, error) {
	// 查询order
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	// 查询orderItem
	items, err := l.svcCtx.OrderItemModel.FindItemsByOrderId(l.ctx, in.Id)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	// 构造返回对象
	var orderItems []*order.OrderItem
	err = copier.Copy(&orderItems, items)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}
	var od order.OrderModel
	err = copier.Copy(&od, res)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}
	od.PayTypeString = orderModel.GetPayTypeStr(res.PayType)
	od.OrderStatusString = orderModel.GetOrderStatusStr(res.OrderStatus)
	od.CreateTime = res.CreateTime.Unix()
	od.OrderItems = orderItems

	return &order.GetOrderByIdResponse{
		Order: &od,
	}, nil
}
