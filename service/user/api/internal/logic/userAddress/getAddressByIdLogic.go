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

type GetAddressByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressByIdLogic {
	return &GetAddressByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressByIdLogic) GetAddressById(req *types.GetAddressByIdRequest) (resp *types.Response, err error) {
	// 获取当前用户
	token, err := l.svcCtx.UserTokenRpc.GetExistToken(l.ctx, &usertoken.GetExistTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取地址失败！" + err.Error(),
		}, nil
	}
	// 查找
	res, err := l.svcCtx.UserAddressRpc.GetAddressById(l.ctx, &useraddress.GetAddressByIdRequest{
		Id:     req.AddressId,
		UserId: token.Token.UserId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "获取地址失败！",
		}, nil
	}
	// 复制
	var address types.Address
	err = copier.Copy(&address, res.Address)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "复制失败" + err.Error(),
		}, nil
	}
	address.CreateTime = time.Unix(res.Address.CreateTime, 0).Format("2006-01-02 15:04:05")
	address.UpdateTime = time.Unix(res.Address.UpdateTime, 0).Format("2006-01-02 15:04:05")

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "SUCCESS",
		Data: types.GetAddressByIdResponse{
			Address: address,
		},
	}, nil
}
