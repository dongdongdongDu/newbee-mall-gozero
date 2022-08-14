package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	goodsInfo "newbee-mall-gozero/service/goods_info/model"
	"newbee-mall-gozero/service/order/model/order"
	"newbee-mall-gozero/service/order/model/order_address"
	"newbee-mall-gozero/service/order/model/order_item"
	"newbee-mall-gozero/service/order/rpc/internal/config"
	shoppingCart "newbee-mall-gozero/service/shopping_cart/model"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"
	userAddress "newbee-mall-gozero/service/user_address/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	OrderModel            order.TbNewbeeMallOrderModel
	OrderItemModel        order_item.TbNewbeeMallOrderItemModel
	OrderAddressModel     order_address.TbNewbeeMallOrderAddressModel
	ShoppingCartItemModel shoppingCart.TbNewbeeMallShoppingCartItemModel
	GoodsInfoModel        goodsInfo.TbNewbeeMallGoodsInfoModel
	ShoppingCartRpc       shoppingcart.ShoppingcartClient
	UserAddressModel      userAddress.TbNewbeeMallUserAddressModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
		OrderModel:            order.NewTbNewbeeMallOrderModel(conn, c.CacheRedis),
		OrderItemModel:        order_item.NewTbNewbeeMallOrderItemModel(conn, c.CacheRedis),
		OrderAddressModel:     order_address.NewTbNewbeeMallOrderAddressModel(conn, c.CacheRedis),
		ShoppingCartItemModel: shoppingCart.NewTbNewbeeMallShoppingCartItemModel(conn, c.CacheRedis),
		GoodsInfoModel:        goodsInfo.NewTbNewbeeMallGoodsInfoModel(conn, c.CacheRedis),
		ShoppingCartRpc:       shoppingcart.NewShoppingcart(zrpc.MustNewClient(c.ShoppingCartRpc)),
		UserAddressModel:      userAddress.NewTbNewbeeMallUserAddressModel(conn, c.CacheRedis),
	}
}
