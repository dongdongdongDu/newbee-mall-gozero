package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/shopping_cart/model"

	"newbee-mall-gozero/service/shopping_cart/rpc/internal/svc"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartItemLogic {
	return &DeleteCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCartItemLogic) DeleteCartItem(in *shoppingcart.DeleteCartItemRequest) (*shoppingcart.EmptyResponse, error) {
	// 查找
	res, err := l.svcCtx.ShoppingCartItemModel.FindNotDelItemById(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("未查询到记录！")
			return nil, errors.New("未查询到记录！")
		}
		logx.Error("查找失败！" + err.Error())
		return nil, errors.New("查找失败！" + err.Error())
	}
	if res.UserId != in.UserId {
		logx.Error("userId不相等！")
		return nil, errors.New("禁止该操作！")
	}
	res.IsDeleted = 1
	err = l.svcCtx.ShoppingCartItemModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("删除失败！" + err.Error())
		return nil, errors.New("删除失败！")
	}

	return &shoppingcart.EmptyResponse{}, nil
}
