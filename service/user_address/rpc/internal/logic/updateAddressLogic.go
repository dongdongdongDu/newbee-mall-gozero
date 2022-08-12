package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/user_address/model"
	"time"

	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressLogic) UpdateAddress(in *useraddress.UpdateAddressRequest) (*useraddress.EmptyResponse, error) {
	// 查找
	res, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.AddressId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的地址")
			return nil, errors.New("不存在的地址")
		}
		logx.Error("地址获取失败" + err.Error())
		return nil, errors.New("地址获取失败" + err.Error())
	}
	if res.UserId != in.User {
		logx.Error("userId不相等")
		return nil, errors.New("禁止该操作")
	}
	// 修改为默认地址，需要将原有的默认地址修改掉
	if in.DefaultFlag == 1 {
		defaultAddress, err := l.svcCtx.UserAddressModel.FindDefault(l.ctx, in.User)
		if err == nil {
			defaultAddress.DefaultFlag = 0
			err := l.svcCtx.UserAddressModel.Update(l.ctx, defaultAddress)
			if err != nil {
				logx.Error("更新原有默认地址失败")
				return nil, errors.New("默认地址获取失败" + err.Error())
			}
		} else if err == model.ErrNotFound {
			logx.Error("不存在默认地址")
		} else {
			logx.Error("默认地址获取失败" + err.Error())
			return nil, errors.New("默认地址获取失败" + err.Error())
		}
	}

	res.UserName = in.UserName
	res.UserPhone = in.UserPhone
	res.DefaultFlag = in.DefaultFlag
	res.ProvinceName = in.ProvinceName
	res.CityName = in.CityName
	res.RegionName = in.RegionName
	res.DetailAddress = in.DetailAddress
	res.UpdateTime = time.Now()

	err = l.svcCtx.UserAddressModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("地址更新失败" + err.Error())
		return nil, errors.New("地址更新失败")
	}

	return &useraddress.EmptyResponse{}, nil
}
