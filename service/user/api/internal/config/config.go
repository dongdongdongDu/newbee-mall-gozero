package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpc         zrpc.RpcClientConf
	UserTokenRpc    zrpc.RpcClientConf
	GoodsInfoRpc    zrpc.RpcClientConf
	UserAddressRpc  zrpc.RpcClientConf
	ShoppingCartRpc zrpc.RpcClientConf
}
