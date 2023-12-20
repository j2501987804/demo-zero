package svc

import (
	"demo/goods/rpc/goodsclient"
	"demo/order/rpc/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	GoodsRpc goodsclient.Goods
}

func NewServiceContext(c config.Config) *ServiceContext {
	println(1)
	return &ServiceContext{
		Config:   c,
		GoodsRpc: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
	}
}
