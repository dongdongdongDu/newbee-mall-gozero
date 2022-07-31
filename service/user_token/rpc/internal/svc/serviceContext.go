package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"newbee-mall-gozero/service/user_token/model"
	"newbee-mall-gozero/service/user_token/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	TokenModel model.UserTokenModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		TokenModel: model.NewUserTokenModel(conn, c.CacheRedis),
	}
}
