// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package server

import (
	"context"

	"demo/goods/rpc/goods"
	"demo/goods/rpc/internal/logic"
	"demo/goods/rpc/internal/svc"
)

type GoodsServer struct {
	svcCtx *svc.ServiceContext
	goods.UnimplementedGoodsServer
}

func NewGoodsServer(svcCtx *svc.ServiceContext) *GoodsServer {
	return &GoodsServer{
		svcCtx: svcCtx,
	}
}

func (s *GoodsServer) Info(ctx context.Context, in *goods.Request) (*goods.Response, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}