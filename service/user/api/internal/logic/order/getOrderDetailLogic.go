package order

import (
	"context"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
	"time"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.OrderRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询订单详情失败！" + err.Error(),
		}, nil
	}
	res, err := l.svcCtx.OrderRpc.GetOrderDetail(l.ctx, &order.OrderRequest{
		OrderNo: req.OrderNo,
		UserId:  token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询订单详情失败！",
		}, nil
	}
	var orderItems []types.OrderItem
	err = copier.Copy(&orderItems, res.OrderItems)
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询订单详情失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetOrderDetailResponse{
			OrderNo:           res.OrderNo,
			TotalPrice:        res.TotalPrice,
			PayStatus:         res.PayStatus,
			PayType:           res.PayType,
			PayTypeString:     res.PayTypeString,
			PayTime:           time.Unix(res.PayTime, 0).Format("2006-01-02 15:04:05"),
			OrderStatus:       res.OrderStatus,
			OrderStatusString: res.OrderStatusString,
			CreateTime:        time.Unix(res.CreateTime, 0).Format("2006-01-02 15:04:05"),
			OrderItems:        orderItems,
		},
	}, nil
}
