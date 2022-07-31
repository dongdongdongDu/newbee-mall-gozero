package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"newbee-mall-gozero/service/user/model"
	"newbee-mall-gozero/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.TbNewbeeMallUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewTbNewbeeMallUserModel(conn, c.CacheRedis),
	}
}
