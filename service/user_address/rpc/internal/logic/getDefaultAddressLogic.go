package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/user_address/model"

	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDefaultAddressLogic {
	return &GetDefaultAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDefaultAddressLogic) GetDefaultAddress(in *useraddress.GetDefaultAddressRequest) (*useraddress.GetDefaultAddressResponse, error) {
	res, err := l.svcCtx.UserAddressModel.FindDefault(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的地址")
			return nil, errors.New("不存在的地址")
		}
		logx.Error("地址获取失败" + err.Error())
		return nil, errors.New("地址获取失败" + err.Error())
	}
	if res.UserId != in.UserId {
		logx.Error("userId不相等")
		return nil, errors.New("禁止该操作")
	}
	// 复制
	var address useraddress.Address
	err = copier.Copy(&address, res)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}
	address.CreateTime = res.CreateTime.Unix()
	address.UpdateTime = res.UpdateTime.Unix()

	return &useraddress.GetDefaultAddressResponse{
		Address: &address,
	}, nil
}
