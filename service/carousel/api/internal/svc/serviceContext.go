package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"newbee-mall-gozero/service/carousel/api/internal/config"
	"newbee-mall-gozero/service/carousel/rpc/carousel"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
)

type ServiceContext struct {
	Config config.Config

	CarouselRpc    carousel.CarouselClient
	IndexConfigRpc indexconfig.IndexconfigClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CarouselRpc:    carousel.NewCarousel(zrpc.MustNewClient(c.CarouselRpc)),
		IndexConfigRpc: indexconfig.NewIndexconfig(zrpc.MustNewClient(c.IndexConfigRpc)),
	}
}
