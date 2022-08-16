package logic

import (
	"context"
	"encoding/json"
	"errors"

	"newbee-mall-gozero/service/mqueue/job/internal/svc"
	"newbee-mall-gozero/service/mqueue/job/jobtype"
	"newbee-mall-gozero/service/order/rpc/order"

	"github.com/hibiken/asynq"
)

// CloseOrderHandler 关闭没有支付的订单
type CloseOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseOrderHandler(svcCtx *svc.ServiceContext) *CloseOrderHandler {
	return &CloseOrderHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 超时关闭没有支付的订单  : if return err != nil , asynq will retry
func (l *CloseOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.DeferCloseOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.New("超时关闭失败:" + err.Error() + " payLoad:" + string(t.Payload()))
	}

	_, err := l.svcCtx.OrderRpc.DeferCloseOrder(ctx, &order.DeferCloseRequest{
		OrderNo: p.OrderNo,
	})
	if err != nil {
		return errors.New("订单" + p.OrderNo + "超时关闭失败:" + err.Error())
	}

	return nil
}
