package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"newbee-mall-gozero/service/user_address/model"
	"newbee-mall-gozero/service/user_address/rpc/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	UserAddressModel model.TbNewbeeMallUserAddressModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:           c,
		UserAddressModel: model.NewTbNewbeeMallUserAddressModel(conn, c.CacheRedis),
	}
}
