package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/sub"
	"newbee-mall-gozero/service/goods_info/model"

	"newbee-mall-gozero/service/shopping_cart/rpc/internal/svc"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartListLogic) GetCartList(in *shoppingcart.GetCartListRequest) (*shoppingcart.GetCartListResponse, error) {
	// 查找
	items, err := l.svcCtx.ShoppingCartItemModel.FindNotDelItemByUser(l.ctx, in.User)
	if err != nil {
		logx.Error("购物车列表获取失败" + err.Error())
		return nil, errors.New("购物车列表获取失败" + err.Error())
	}
	// 和GetCartItems中一样
	// 所有商品id
	var goodsIds []int64
	for _, item := range items {
		goodsIds = append(goodsIds, item.GoodsId)
	}
	// 查找商品
	goodsList, err := l.svcCtx.GoodsInfoModel.GetForIndex(l.ctx, goodsIds)
	if err != nil {
		logx.Error("商品列表获取失败" + err.Error())
		return nil, errors.New("商品列表获取失败" + err.Error())
	}
	// 构造map [商品id]商品
	goodsMap := make(map[int64]*model.TbNewbeeMallGoodsInfo)
	for _, goodsInfo := range goodsList {
		goodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	// 构造列表
	var cartList []*shoppingcart.CartItem
	for _, item := range items {
		var cartItem shoppingcart.CartItem
		// 只能复制 CartItemId GoodsId GoodsCount
		err = copier.Copy(&cartItem, item)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return nil, errors.New("复制失败" + err.Error())
		}
		// 从map中复制
		if info, ok := goodsMap[cartItem.GoodsId]; ok {
			cartItem.GoodsCoverImg = info.GoodsCoverImg
			cartItem.GoodsName = sub.SubStrLen(info.GoodsName, 28)
			cartItem.SellingPrice = info.SellingPrice
		}
		cartList = append(cartList, &cartItem)
	}

	return &shoppingcart.GetCartListResponse{
		CartList: cartList,
	}, nil
}
