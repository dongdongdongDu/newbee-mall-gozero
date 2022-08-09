package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"newbee-mall-gozero/service/goods_category/model"
	"newbee-mall-gozero/service/goods_category/rpc/internal/config"
)

type ServiceContext struct {
	Config             config.Config
	GoodsCategoryModel model.TbNewbeeMallGoodsCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		GoodsCategoryModel: model.NewTbNewbeeMallGoodsCategoryModel(conn, c.CacheRedis),
	}
}
