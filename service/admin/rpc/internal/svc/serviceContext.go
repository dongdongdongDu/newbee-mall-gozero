package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/admin/model"
	"newbee-mall-gozero/service/admin/rpc/internal/config"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
)

type ServiceContext struct {
	Config        config.Config
	AdminModel    model.TbNewbeeMallAdminUserModel
	AdminTokenRpc admintoken.AdmintokenClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		AdminModel:    model.NewTbNewbeeMallAdminUserModel(conn, c.CacheRedis),
		AdminTokenRpc: admintoken.NewAdmintoken(zrpc.MustNewClient(c.AdminTokenRpc)),
	}
}
