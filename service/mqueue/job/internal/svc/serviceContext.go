package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/mqueue/job/internal/config"
	"newbee-mall-gozero/service/order/rpc/order"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server

	OrderRpc order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		OrderRpc:    order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
