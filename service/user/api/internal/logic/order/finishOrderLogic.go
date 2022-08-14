package order

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinishOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFinishOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinishOrderLogic {
	return &FinishOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FinishOrderLogic) FinishOrder(req *types.OrderRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "订单签收失败！" + err.Error(),
		}, nil
	}
	_, err = l.svcCtx.OrderRpc.FinishOrder(l.ctx, &order.OrderRequest{
		OrderNo: req.OrderNo,
		UserId:  token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "订单签收失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "订单签收成功",
		Data:       map[string]interface{}{},
	}, nil
}
