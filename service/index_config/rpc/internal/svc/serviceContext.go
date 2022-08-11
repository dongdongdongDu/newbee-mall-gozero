package svc

import (
	goodsInfo "newbee-mall-gozero/service/goods_info/model"
	"newbee-mall-gozero/service/index_config/model"
	"newbee-mall-gozero/service/index_config/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ConfigModel    model.TbNewbeeMallIndexConfigModel
	GoodsInfoModel goodsInfo.TbNewbeeMallGoodsInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		ConfigModel:    model.NewTbNewbeeMallIndexConfigModel(conn, c.CacheRedis),
		GoodsInfoModel: goodsInfo.NewTbNewbeeMallGoodsInfoModel(conn, c.CacheRedis),
	}
}
