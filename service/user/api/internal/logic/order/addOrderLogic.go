package order

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/common/verify"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrderLogic) AddOrder(req *types.AddOrderRequest) (resp *types.Response, err error) {
	// 校验输入格式
	if err := verify.Verify(*req, verify.SaveOrderParamVerify); err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        err.Error(),
		}, nil
	}
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "生成订单失败！" + err.Error(),
		}, nil
	}
	res, err := l.svcCtx.OrderRpc.AddOrder(l.ctx, &order.AddOrderRequest{
		CartItemIds: req.CartItemIds,
		AddressId:   req.AddressId,
		UserId:      token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "生成订单失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data:       res.OrderNo,
	}, nil
}
