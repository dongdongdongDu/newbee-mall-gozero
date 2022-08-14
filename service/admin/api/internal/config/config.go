package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AdminRpc         zrpc.RpcClientConf
	AdminTokenRpc    zrpc.RpcClientConf
	GoodsInfoRpc     zrpc.RpcClientConf
	GoodsCategoryRpc zrpc.RpcClientConf
	CarouselRpc      zrpc.RpcClientConf
	IndexConfigRpc   zrpc.RpcClientConf
	OrderRpc         zrpc.RpcClientConf
}
