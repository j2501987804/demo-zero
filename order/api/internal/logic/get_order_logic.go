package logic

import (
	"context"
	"demo/order/rpc/order"

	"demo/order/api/internal/svc"
	"demo/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderResp, err error) {
	info, err := l.svcCtx.OrderRpc.GetOrder(l.ctx, &order.Request{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.OrderResp{
		Id:        info.Id,
		GoodsName: info.GoodsName,
		GoodsId:   info.GoodsId,
	}, nil
}
