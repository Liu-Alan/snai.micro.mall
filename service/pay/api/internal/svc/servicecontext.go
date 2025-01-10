package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"snai.micro.mall/service/pay/api/internal/config"
	pay "snai.micro.mall/service/pay/rpc/payclient"
)

type ServiceContext struct {
	Config config.Config

	PayRpc pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
