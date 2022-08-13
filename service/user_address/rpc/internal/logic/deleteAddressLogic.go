package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/user_address/model"

	"newbee-mall-gozero/service/user_address/rpc/internal/svc"
	"newbee-mall-gozero/service/user_address/rpc/useraddress"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAddressLogic {
	return &DeleteAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAddressLogic) DeleteAddress(in *useraddress.DeleteAddressRequest) (*useraddress.EmptyResponse, error) {
	// 查找
	res, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.Id)
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
	res.IsDeleted = 1
	err = l.svcCtx.UserAddressModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("地址删除失败" + err.Error())
		return nil, errors.New("地址删除失败")
	}

	return &useraddress.EmptyResponse{}, nil
}
