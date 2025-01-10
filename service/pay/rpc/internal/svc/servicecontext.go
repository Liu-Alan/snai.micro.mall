package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	order "snai.micro.mall/service/order/rpc/orderclient"
	"snai.micro.mall/service/pay/model"
	"snai.micro.mall/service/pay/rpc/internal/config"
	user "snai.micro.mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel

	UserRpc  user.User
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
