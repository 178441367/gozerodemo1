package logic

import (
	"context"
	"github.com/178441367/gozerodemo1/rpcdemo/rpcdemo"

	"apidemo/internal/svc"
	"apidemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApidemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApidemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApidemoLogic {
	return &ApidemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApidemoLogic) Apidemo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.RpcClient.Ping(l.ctx, &rpcdemo.Request{Ping: req.Name})
	if err != nil {
		return nil, err
	}
	resp = &types.Response{
		Message: res.Pong,
	}
	return
}
