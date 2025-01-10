package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"snai.micro.mall/service/user/api/internal/config"
	user "snai.micro.mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
