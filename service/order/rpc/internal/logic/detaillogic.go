package logic

import (
	"context"

	"google.golang.org/grpc/status"
	"snai.micro.mall/service/order/rpc/internal/svc"
	"snai.micro.mall/service/order/rpc/types/order"
	"snai.micro.mall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *order.DetailRequest) (*order.DetailResponse, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &order.DetailResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Pid:    res.Pid,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
