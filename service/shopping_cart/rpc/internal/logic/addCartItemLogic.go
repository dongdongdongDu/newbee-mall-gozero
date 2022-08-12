package logic

import (
	"context"
	"errors"
	"time"

	"newbee-mall-gozero/service/shopping_cart/model"
	"newbee-mall-gozero/service/shopping_cart/rpc/internal/svc"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartItemLogic) AddCartItem(in *shoppingcart.AddCartItemRequest) (*shoppingcart.EmptyResponse, error) {
	// 查找商品
	_, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, in.GoodsId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的商品")
			return nil, errors.New("不存在的商品")
		}
		logx.Error("商品信息获取失败" + err.Error())
		return nil, errors.New("商品信息获取失败" + err.Error())
	}
	// 查找是否已在购物车里
	res, err := l.svcCtx.ShoppingCartItemModel.FindItem(l.ctx, in.GoodsId, in.User)
	if err == nil {
		logx.Error("查找失败" + err.Error())
		return nil, errors.New("查找失败" + err.Error())
	}
	if len(res) != 0 {
		logx.Error("已存在！" + err.Error())
		return nil, errors.New("已存在！无需重复添加！" + err.Error())
	}
	total, err := l.svcCtx.ShoppingCartItemModel.CountUserItemAll(l.ctx, in.User)
	if err != nil {
		logx.Error("获取购物车总数失败" + err.Error())
		return nil, errors.New("获取购物车总数失败" + err.Error())
	}
	if total > 20 {
		logx.Error("超出购物车最大容量！")
		return nil, errors.New("超出购物车最大容量！")
	}
	item := model.TbNewbeeMallShoppingCartItem{
		UserId:     in.User,
		GoodsId:    in.GoodsId,
		GoodsCount: in.GoodsCount,
		IsDeleted:  0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	_, err = l.svcCtx.ShoppingCartItemModel.Insert(l.ctx, &item)
	if err != nil {
		logx.Error("添加失败：" + err.Error())
		return nil, err
	}

	return &shoppingcart.EmptyResponse{}, nil
}
