package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyAddressLogic {
	return &GetMyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyAddressLogic) GetMyAddress(in *useraddress.GetMyAddressRequest) (*useraddress.GetMyAddressResponse, error) {
	addressLsit, err := l.svcCtx.UserAddressModel.FindMyAddress(l.ctx, in.User)
	if err != nil {
		logx.Error("地址获取失败" + err.Error())
		return nil, errors.New("地址获取失败" + err.Error())
	}
	// 构造列表
	MyAddressList := make([]*useraddress.Address, 0)
	for _, item := range addressLsit {
		var address useraddress.Address
		err := copier.Copy(&address, item)
		if err != nil {
			logx.Error("复制失败" + err.Error())
			return nil, errors.New("复制失败" + err.Error())
		}
		address.CreateTime = item.CreateTime.Unix()
		address.UpdateTime = item.UpdateTime.Unix()
		MyAddressList = append(MyAddressList, &address)
	}

	return &useraddress.GetMyAddressResponse{
		AddressList: MyAddressList,
	}, nil
}
