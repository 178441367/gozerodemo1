package main

import (
	"flag"
	"fmt"
	"github.com/178441367/gozerodemo1/rpcdemo/rpcdemo"

	"github.com/178441367/gozerodemo1/rpcdemo/internal/config"
	"github.com/178441367/gozerodemo1/rpcdemo/internal/server"
	"github.com/178441367/gozerodemo1/rpcdemo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rpcdemo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpcdemo.RegisterRpcdemoServer(grpcServer, server.NewRpcdemoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
