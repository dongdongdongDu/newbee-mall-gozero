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

type UpdateCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartItemLogic) UpdateCartItem(req *types.UpdateCartItemRequest) (resp *types.Response, err error) {
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
			Msg:        "修改购物车失败！" + err.Error(),
		}, nil
	}
	// 更新
	_, err = l.svcCtx.ShoppingCartRpc.UpdateCartItem(l.ctx, &shoppingcart.UpdateCartItemRequest{
		CartItemId: req.CartItemId,
		GoodsCount: req.GoodsCount,
		UserId:     token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "修改购物车失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "修改购物车成功",
		Data:       map[string]interface{}{},
	}, nil
}
