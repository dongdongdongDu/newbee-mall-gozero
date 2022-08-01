package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/user/model"
	"newbee-mall-gozero/service/user/rpc/internal/config"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
)

type ServiceContext struct {
	Config       config.Config
	UserModel    model.TbNewbeeMallUserModel
	UserTokenRpc usertoken.UsertokenClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		UserModel:    model.NewTbNewbeeMallUserModel(conn, c.CacheRedis),
		UserTokenRpc: usertoken.NewUsertoken(zrpc.MustNewClient(c.UserTokenRpc)),
	}
}
