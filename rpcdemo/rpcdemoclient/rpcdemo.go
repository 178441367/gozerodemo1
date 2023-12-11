// Code generated by goctl. DO NOT EDIT.
// Source: rpcdemo.proto

package rpcdemoclient

import (
	"context"

	"rpcdemo/rpcdemo"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = rpcdemo.Request
	Response = rpcdemo.Response

	Rpcdemo interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultRpcdemo struct {
		cli zrpc.Client
	}
)

func NewRpcdemo(cli zrpc.Client) Rpcdemo {
	return &defaultRpcdemo{
		cli: cli,
	}
}

func (m *defaultRpcdemo) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := rpcdemo.NewRpcdemoClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
