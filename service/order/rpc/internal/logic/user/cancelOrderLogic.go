package user

import (
	"context"
	"errors"
	"newbee-mall-gozero/common/nums"
	"time"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelOrderLogic) CancelOrder(in *order.OrderRequest) (*order.EmptyResponse, error) {
	// 查找
	res, err := l.svcCtx.OrderModel.FindOneByNo(l.ctx, in.OrderNo)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	if res.UserId != in.UserId {
		logx.Error("userId不相等")
		return nil, errors.New("禁止该操作")
	}
	if nums.NumInList(res.OrderStatus, []int64{4, -1, -2, -3}) {
		logx.Error("订单状态异常！")
		return nil, errors.New("订单状态异常！")
	}
	res.OrderStatus = -1
	res.UpdateTime = time.Now()
	// 更新
	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("订单更新失败" + err.Error())
		return nil, errors.New("订单更新失败" + err.Error())
	}

	return &order.EmptyResponse{}, nil
}
