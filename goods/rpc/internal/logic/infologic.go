package logic

import (
	"context"

	"demo/goods/rpc/goods"
	"demo/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *goods.Request) (*goods.Response, error) {
	return &goods.Response{
		Id:   in.Id,
		Name: "苹果",
	}, nil
}
