package user

import (
	"context"
	"errors"
	"time"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinishOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFinishOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinishOrderLogic {
	return &FinishOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FinishOrderLogic) FinishOrder(in *order.OrderRequest) (*order.EmptyResponse, error) {
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
	res.OrderStatus = 4
	res.UpdateTime = time.Now()
	// 更新
	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("订单更新失败" + err.Error())
		return nil, errors.New("订单更新失败" + err.Error())
	}

	return &order.EmptyResponse{}, nil
}
