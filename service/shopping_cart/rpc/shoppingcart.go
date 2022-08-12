package main

import (
	"flag"
	"fmt"

	"newbee-mall-gozero/service/shopping_cart/rpc/internal/config"
	"newbee-mall-gozero/service/shopping_cart/rpc/internal/server"
	"newbee-mall-gozero/service/shopping_cart/rpc/internal/svc"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/shoppingcart.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewShoppingcartServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		shoppingcart.RegisterShoppingcartServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
