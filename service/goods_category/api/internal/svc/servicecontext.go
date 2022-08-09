package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/goods_category/api/internal/config"
	"newbee-mall-gozero/service/goods_category/rpc/goodscategory"
)

type ServiceContext struct {
	Config config.Config

	GoodsCategoryRpc goodscategory.GoodscategoryClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		GoodsCategoryRpc: goodscategory.NewGoodscategory(zrpc.MustNewClient(c.GoodsCategoryRpc)),
	}
}
