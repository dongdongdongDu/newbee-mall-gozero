package logic

import (
	"context"
	"newbee-mall-gozero/service/user_address/model"
	"time"

	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAddressLogic {
	return &SaveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveAddressLogic) SaveAddress(in *useraddress.SaveAddressRequest) (*useraddress.EmptyResponse, error) {
	address := model.TbNewbeeMallUserAddress{
		UserId:        in.User,
		UserName:      in.UserName,
		UserPhone:     in.UserPhone,
		DefaultFlag:   in.DefaultFlag,
		ProvinceName:  in.ProvinceName,
		CityName:      in.CityName,
		RegionName:    in.RegionName,
		DetailAddress: in.DetailAddress,
		IsDeleted:     0,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}

	_, err := l.svcCtx.UserAddressModel.Insert(l.ctx, &address)
	if err != nil {
		logx.Error("添加失败：" + err.Error())
		return nil, err
	}

	return &useraddress.EmptyResponse{}, nil
}
