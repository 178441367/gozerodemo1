package logic

import (
	"context"
	"github.com/178441367/gozerodemo1/apidemo/internal/svc"
	"github.com/178441367/gozerodemo1/apidemo/internal/types"
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

	return
}
