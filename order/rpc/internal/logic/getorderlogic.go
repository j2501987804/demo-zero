package logic

import (
	"context"
	"demo/goods/rpc/goods"

	"demo/order/rpc/internal/svc"
	"demo/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *order.Request) (*order.Response, error) {
	info, err := l.svcCtx.GoodsRpc.Info(l.ctx, &goods.Request{Id: 2})
	if err != nil {
		return nil, err
	}
	return &order.Response{
		Id:        in.Id,
		GoodsId:   info.Id,
		GoodsName: info.Name,
	}, nil
}
