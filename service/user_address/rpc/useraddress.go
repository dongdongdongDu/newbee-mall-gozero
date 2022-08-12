package main

import (
	"flag"
	"fmt"

	"newbee-mall-gozero/service/user_address/rpc/internal/config"
	"newbee-mall-gozero/service/user_address/rpc/internal/server"
	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/useraddress.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUseraddressServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		useraddress.RegisterUseraddressServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
