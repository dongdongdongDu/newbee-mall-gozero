package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/user/api/internal/config"
	"newbee-mall-gozero/service/user/api/internal/middleware"
	"newbee-mall-gozero/service/user/rpc/user"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
)

type ServiceContext struct {
	Config       config.Config
	UserRpc      user.UserClient
	UserJwtAuth  rest.Middleware
	UserTokenRpc usertoken.UsertokenClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserTokenRpc := usertoken.NewUsertoken(zrpc.MustNewClient(c.UserTokenRpc))
	return &ServiceContext{
		Config:       c,
		UserRpc:      user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserJwtAuth:  middleware.NewUserJwtAuthMiddleware(UserTokenRpc).Handle,
		UserTokenRpc: UserTokenRpc,
	}
}
