package shoppingCart

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartItemLogic) AddCartItem(req *types.AddCartItemRequest) (resp *types.Response, err error) {
	if req.GoodsCount < 1 {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "商品数量不能小于 1 ！",
		}, nil
	}
	if req.GoodsCount > 5 {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "超出单个商品的最大购买数量！",
		}, nil
	}
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "添加购物车失败！" + err.Error(),
		}, nil
	}

	// 创建
	_, err = l.svcCtx.ShoppingCartRpc.AddCartItem(l.ctx, &shoppingcart.AddCartItemRequest{
		GoodsId:    req.GoodsId,
		GoodsCount: req.GoodsCount,
		User:       token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "添加购物车失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "添加购物车成功",
		Data:       map[string]interface{}{},
	}, nil
}
