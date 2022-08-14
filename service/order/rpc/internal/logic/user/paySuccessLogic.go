package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaySuccessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaySuccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaySuccessLogic {
	return &PaySuccessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaySuccessLogic) PaySuccess(in *order.PaySuccessRequest) (*order.EmptyResponse, error) {
	// 查询
	res, err := l.svcCtx.OrderModel.FindOneByNo(l.ctx, in.OrderNo)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	if res.OrderStatus != 0 {
		logx.Error("订单状态异常！")
		return nil, errors.New("订单状态异常！")
	}
	// 修改订单状态模拟支付
	res.OrderStatus = 1
	res.PayType = in.PayType
	res.PayStatus = 1
	res.PayTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	res.UpdateTime = time.Now()
	// 更新
	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("订单更新失败" + err.Error())
		return nil, errors.New("订单更新失败" + err.Error())
	}

	return &order.EmptyResponse{}, nil
}
