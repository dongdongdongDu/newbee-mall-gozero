package order

import (
	"context"
	"github.com/jinzhu/copier"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.GetOrderListRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！" + err.Error(),
		}, nil
	}
	res, err := l.svcCtx.OrderRpc.GetOrderList(l.ctx, &order.GetOrderListRequest{
		Status:     req.Status,
		PageNumber: req.PageNumber,
		UserId:     token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！",
		}, nil
	}
	var list []types.ListItem
	for _, model := range res.OrderList {
		var listItem types.ListItem
		err = copier.Copy(&listItem, model)
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "查询失败！",
			}, nil
		}
		listItem.CreateTime = time.Unix(model.CreateTime, 0).Format("2006-01-02 15:04:05")
		list = append(list, listItem)
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetListResponse{
			List:       list,
			TotalCount: res.Total,
			TotalPage:  int64(math.Ceil(float64(res.Total) / float64(5))),
			CurrPage:   req.PageNumber,
			PageSize:   5,
		},
	}, nil
}
