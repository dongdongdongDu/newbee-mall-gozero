package admin

import (
	"context"
	"errors"
	"time"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseOrderLogic {
	return &CloseOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CloseOrderLogic) CloseOrder(in *order.CheckOrderRequest) (*order.EmptyResponse, error) {
	for _, id := range in.Ids {
		res, err := l.svcCtx.OrderModel.FindOne(l.ctx, id)
		if err != nil {
			logx.Error("查询失败" + err.Error())
			return nil, errors.New("查询失败" + err.Error())
		}
		if res.IsDeleted != 0 {
			logx.Error("订单" + res.OrderNo + "已删除")
			return nil, errors.New("订单" + res.OrderNo + "已删除")
		}
		if res.OrderStatus == 4 || res.OrderStatus < 0 {
			logx.Error("订单" + res.OrderNo + "不能执行关闭操作")
			return nil, errors.New("订单" + res.OrderNo + "不能执行关闭操作")
		}
		res.OrderStatus = -3
		res.UpdateTime = time.Now()
		err = l.svcCtx.OrderModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("订单" + res.OrderNo + "关闭失败")
			return nil, errors.New("订单" + res.OrderNo + "关闭失败")
		}
	}
	return &order.EmptyResponse{}, nil
}
