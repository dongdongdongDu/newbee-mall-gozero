package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/admin/api/internal/config"
	"newbee-mall-gozero/service/admin/api/internal/middleware"
	"newbee-mall-gozero/service/admin/rpc/admin"
	"newbee-mall-gozero/service/admin_token/rpc/admintoken"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
)

type ServiceContext struct {
	Config           config.Config
	AdminRpc         admin.AdminClient
	AdminJwtAuth     rest.Middleware
	AdminTokenRpc    admintoken.AdmintokenClient
	GoodsInfoRpc     goodsinfo.GoodsinfoClient
	GoodsCategoryRpc goodscategory.GoodscategoryClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	AdminTokenRpc := admintoken.NewAdmintoken(zrpc.MustNewClient(c.AdminTokenRpc))
	return &ServiceContext{
		Config:           c,
		AdminRpc:         admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		AdminJwtAuth:     middleware.NewAdminJwtAuthMiddleware(AdminTokenRpc).Handle,
		AdminTokenRpc:    AdminTokenRpc,
		GoodsInfoRpc:     goodsinfo.NewGoodsinfo(zrpc.MustNewClient(c.GoodsInfoRpc)),
		GoodsCategoryRpc: goodscategory.NewGoodscategory(zrpc.MustNewClient(c.GoodsCategoryRpc)),
	}
}
