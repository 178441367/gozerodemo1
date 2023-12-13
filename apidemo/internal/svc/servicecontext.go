package svc

import (
	"apidemo/internal/config"
	"github.com/178441367/gozerodemo1/rpcdemo/rpcdemoclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// 其他服务的客户端
	RpcClient rpcdemoclient.Rpcdemo
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "127.0.0.1:8080",
	})
	return &ServiceContext{
		Config:    c,
		RpcClient: rpcdemoclient.NewRpcdemo(client),
	}
}
