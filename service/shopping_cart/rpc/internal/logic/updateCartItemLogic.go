package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/shopping_cart/model"
	"time"

	"newbee-mall-gozero/service/shopping_cart/rpc/internal/svc"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartItemLogic) UpdateCartItem(in *shoppingcart.UpdateCartItemRequest) (*shoppingcart.EmptyResponse, error) {
	res, err := l.svcCtx.ShoppingCartItemModel.FindNotDelItemById(l.ctx, in.CartItemId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("未查询到记录！")
			return nil, errors.New("未查询到记录！")
		}
		logx.Error("查找失败！" + err.Error())
		return nil, errors.New("查找失败！" + err.Error())
	}
	if res.UserId != in.UserId {
		logx.Error("userId不相等")
		return nil, errors.New("禁止该操作")
	}
	res.GoodsCount = in.GoodsCount
	res.UpdateTime = time.Now()

	err = l.svcCtx.ShoppingCartItemModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("更新失败" + err.Error())
		return nil, errors.New("更新失败")
	}

	return &shoppingcart.EmptyResponse{}, nil
}
