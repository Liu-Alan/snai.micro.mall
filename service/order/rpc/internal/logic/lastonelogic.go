package logic

import (
	"context"

	"google.golang.org/grpc/status"
	"snai.micro.mall/service/order/rpc/internal/svc"
	"snai.micro.mall/service/order/rpc/types/order"
	"snai.micro.mall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LastOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLastOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LastOneLogic {
	return &LastOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LastOneLogic) LastOne(in *order.LastOneRequest) (*order.LastOneResponse, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &order.LastOneResponse{
		Id: res.Id,
	}, nil
}
