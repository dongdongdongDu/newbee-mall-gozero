package logic

import (
	"context"
	"errors"
	"time"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeferCloseOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeferCloseOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeferCloseOrderLogic {
	return &DeferCloseOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeferCloseOrderLogic) DeferCloseOrder(in *order.DeferCloseRequest) (*order.EmptyResponse, error) {
	res, err := l.svcCtx.OrderModel.FindOneByNo(l.ctx, in.OrderNo)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	// 订单没被删除且是待支付状态
	if res.IsDeleted == 0 && res.OrderStatus == 0 {
		res.OrderStatus = -2
		res.UpdateTime = time.Now()
		err = l.svcCtx.OrderModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("订单" + res.OrderNo + "关闭失败")
			return nil, errors.New("订单" + res.OrderNo + "关闭失败")
		}
	}

	return &order.EmptyResponse{}, nil
}
