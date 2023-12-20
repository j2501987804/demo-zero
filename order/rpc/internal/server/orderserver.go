// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"demo/order/rpc/internal/logic"
	"demo/order/rpc/internal/svc"
	"demo/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) GetOrder(ctx context.Context, in *order.Request) (*order.Response, error) {
	l := logic.NewGetOrderLogic(ctx, s.svcCtx)
	return l.GetOrder(in)
}