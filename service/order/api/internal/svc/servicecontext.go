package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"snai.micro.mall/service/order/api/internal/config"
	order "snai.micro.mall/service/order/rpc/orderclient"
	product "snai.micro.mall/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   order.Order
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
