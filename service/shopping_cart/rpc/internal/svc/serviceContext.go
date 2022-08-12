package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	goodsInfo "newbee-mall-gozero/service/goods_info/model"
	"newbee-mall-gozero/service/shopping_cart/model"
	"newbee-mall-gozero/service/shopping_cart/rpc/internal/config"
)

type ServiceContext struct {
	Config                config.Config
	ShoppingCartItemModel model.TbNewbeeMallShoppingCartItemModel
	GoodsInfoModel        goodsInfo.TbNewbeeMallGoodsInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
		ShoppingCartItemModel: model.NewTbNewbeeMallShoppingCartItemModel(conn, c.CacheRedis),
		GoodsInfoModel:        goodsInfo.NewTbNewbeeMallGoodsInfoModel(conn, c.CacheRedis),
	}
}
