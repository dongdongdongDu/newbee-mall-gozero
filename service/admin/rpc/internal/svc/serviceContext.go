package svc

import (
	"newbee-mall-gozero/service/admin/model"
	"newbee-mall-gozero/service/admin/rpc/internal/config"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	user "newbee-mall-gozero/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AdminModel    model.TbNewbeeMallAdminUserModel
	AdminTokenRpc admintoken.AdmintokenClient
	UserModel     user.TbNewbeeMallUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		AdminModel:    model.NewTbNewbeeMallAdminUserModel(conn, c.CacheRedis),
		AdminTokenRpc: admintoken.NewAdmintoken(zrpc.MustNewClient(c.AdminTokenRpc)),
		UserModel:     user.NewTbNewbeeMallUserModel(conn, c.CacheRedis),
	}
}
