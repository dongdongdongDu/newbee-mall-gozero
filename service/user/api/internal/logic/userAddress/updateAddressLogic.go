package userAddress

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"
	"newbee-mall-gozero/service/user_token/rpc/usertoken"

	"newbee-mall-gozero/service/user/api/internal/svc"
	"newbee-mall-gozero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAddressLogic) UpdateAddress(req *types.UpdateAddressRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新用户地址失败！" + err.Error(),
		}, nil
	}
	// 更新
	_, err = l.svcCtx.UserAddressRpc.UpdateAddress(l.ctx, &useraddress.UpdateAddressRequest{
		AddressId:     req.AddressId,
		UserName:      req.UserName,
		UserPhone:     req.UserPhone,
		DefaultFlag:   req.DefaultFlag,
		ProvinceName:  req.ProvinceName,
		CityName:      req.CityName,
		RegionName:    req.RegionName,
		DetailAddress: req.DetailAddress,
		User:          token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "更新用户地址失败！",
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "更新用户地址成功",
		Data:       map[string]interface{}{},
	}, nil
}
