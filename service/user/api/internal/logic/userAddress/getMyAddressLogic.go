package userAddress

import (
	"context"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"
	"time"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyAddressLogic {
	return &GetMyAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyAddressLogic) GetMyAddress(req *types.GetMyAddressRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取用户地址失败！" + err.Error(),
		}, nil
	}
	// 查找
	res, err := l.svcCtx.UserAddressRpc.GetMyAddress(l.ctx, &useraddress.GetMyAddressRequest{
		UserId: token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取用户地址失败！",
		}, nil
	}
	// 复制
	var addressList []types.Address
	for _, item := range res.AddressList {
		var address types.Address
		err := copier.Copy(&address, item)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return &types.Response{
				ResultCode: response.ERROR,
				Msg:        "获取用户地址失败" + err.Error(),
			}, nil
		}
		address.CreateTime = time.Unix(item.CreateTime, 0).Format("2006-01-02 15:04:05")
		address.UpdateTime = time.Unix(item.UpdateTime, 0).Format("2006-01-02 15:04:05")
		addressList = append(addressList, address)
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetMyAddressResponse{
			AddressList: addressList,
		},
	}, nil
}
