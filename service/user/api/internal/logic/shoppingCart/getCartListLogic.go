package shoppingCart

import (
	"context"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartListLogic) GetCartList(req *types.GetCartListRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取购物车失败！" + err.Error(),
		}, nil
	}
	// 查询
	res, err := l.svcCtx.ShoppingCartRpc.GetCartList(l.ctx, &shoppingcart.GetCartListRequest{
		User: token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取购物车失败！" + err.Error(),
		}, nil
	}
	// 构造列表
	var cartList []types.CartItem
	for _, i := range res.CartList {
		var item types.CartItem
		err = copier.Copy(&item, i)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取购物明细异常" + err.Error(),
			}, nil
		}
		cartList = append(cartList, item)
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetCartListResponse{
			CartItemList: cartList,
		},
	}, nil
}
