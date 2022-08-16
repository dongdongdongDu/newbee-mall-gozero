package svc

import (
	"github.com/hibiken/asynq"
	"newbee-mall-gozero/service/order/rpc/internal/config"
)

//create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
