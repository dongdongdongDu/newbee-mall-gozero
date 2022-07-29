package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/user/api/internal/config"
	"newbee-mall-gozero/service/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
