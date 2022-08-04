package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/goods_info/api/internal/config"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
)

type ServiceContext struct {
	Config       config.Config
	GoodsInfoRpc goodsinfo.GoodsinfoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		GoodsInfoRpc: goodsinfo.NewGoodsinfo(zrpc.MustNewClient(c.GoodsInfoRpc)),
	}
}
