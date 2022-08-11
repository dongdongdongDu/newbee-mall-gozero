package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"newbee-mall-gozero/service/carousel/model"
	"newbee-mall-gozero/service/carousel/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	CarouselModel model.TbNewbeeMallCarouselModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		CarouselModel: model.NewTbNewbeeMallCarouselModel(conn, c.CacheRedis),
	}
}
