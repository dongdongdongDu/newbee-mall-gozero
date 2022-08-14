package order

import (
	"context"
	"github.com/jinzhu/copier"
	"math"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/order/rpc/order"
	"time"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrdersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersListLogic {
	return &GetOrdersListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrdersListLogic) GetOrdersList(req *types.GetOrdersListRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.OrderRpc.GetOrdersList(l.ctx, &order.GetOrdersListRequest{
		OrderNo:     req.OrderNo,
		OrderStatus: req.OrderStatus,
		PageNumber:  req.PageNumber,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "查询失败！",
		}, nil
	}
	var list []types.OrdersListItem
	for _, model := range res.OrderList {
		var listItem types.OrdersListItem
		err = copier.Copy(&listItem, model)
		if err != nil {
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "查询失败！",
			}, nil
		}
		listItem.PayTime = time.Unix(model.PayTime, 0).Format("2006-01-02 15:04:05")
		listItem.CreateTime = time.Unix(model.CreateTime, 0).Format("2006-01-02 15:04:05")
		listItem.UpdateTime = time.Unix(model.UpdateTime, 0).Format("2006-01-02 15:04:05")
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
			PageSize:   req.PageSize,
		},
	}, nil
}
