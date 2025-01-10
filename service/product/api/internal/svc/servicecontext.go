package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"snai.micro.mall/service/product/api/internal/config"
	product "snai.micro.mall/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
