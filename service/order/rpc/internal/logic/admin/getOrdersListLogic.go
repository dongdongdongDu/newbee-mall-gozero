package admin

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrdersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersListLogic {
	return &GetOrdersListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrdersListLogic) GetOrdersList(in *order.GetOrdersListRequest) (*order.GetOrdersListResponse, error) {
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	// 查询
	res, total, err := l.svcCtx.OrderModel.FindByNoAndStatus(l.ctx, in.OrderNo, in.OrderStatus, limit, offset)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	// 构造列表
	var orderList []*order.OrderInfo
	for _, od := range res {
		var orderInfo order.OrderInfo
		err := copier.Copy(&orderInfo, od)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return nil, errors.New("复制失败" + err.Error())
		}
		orderInfo.PayTime = od.PayTime.Time.Unix()
		orderInfo.CreateTime = od.CreateTime.Unix()
		orderInfo.UpdateTime = od.UpdateTime.Unix()
		orderList = append(orderList, &orderInfo)
	}

	return &order.GetOrdersListResponse{
		OrderList: orderList,
		Total:     total,
	}, nil
}
