package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"
	"newbee-mall-gozero/service/user/api/internal/config"
	"newbee-mall-gozero/service/user/api/internal/middleware"
	"newbee-mall-gozero/service/user/rpc/user"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
)

type ServiceContext struct {
	Config          config.Config
	UserRpc         user.UserClient
	UserJwtAuth     rest.Middleware
	UserTokenRpc    usertoken.UsertokenClient
	GoodsInfoRpc    goodsinfo.GoodsinfoClient
	UserAddressRpc  useraddress.UseraddressClient
	ShoppingCartRpc shoppingcart.ShoppingcartClient
	OrderRpc        order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserTokenRpc := usertoken.NewUsertoken(zrpc.MustNewClient(c.UserTokenRpc))
	return &ServiceContext{
		Config:          c,
		UserRpc:         user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserJwtAuth:     middleware.NewUserJwtAuthMiddleware(UserTokenRpc).Handle,
		UserTokenRpc:    UserTokenRpc,
		GoodsInfoRpc:    goodsinfo.NewGoodsinfo(zrpc.MustNewClient(c.GoodsInfoRpc)),
		UserAddressRpc:  useraddress.NewUseraddress(zrpc.MustNewClient(c.UserAddressRpc)),
		ShoppingCartRpc: shoppingcart.NewShoppingcart(zrpc.MustNewClient(c.ShoppingCartRpc)),
		OrderRpc:        order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
