package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	category "newbee-mall-gozero/service/goods_category/model"
	"newbee-mall-gozero/service/goods_info/model"
	"newbee-mall-gozero/service/goods_info/rpc/internal/config"
)

type ServiceContext struct {
	Config             config.Config
	GoodsInfoModel     model.TbNewbeeMallGoodsInfoModel
	GoodsCategoryModel category.TbNewbeeMallGoodsCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		GoodsInfoModel:     model.NewTbNewbeeMallGoodsInfoModel(conn, c.CacheRedis),
		GoodsCategoryModel: category.NewTbNewbeeMallGoodsCategoryModel(conn, c.CacheRedis),
	}
}
