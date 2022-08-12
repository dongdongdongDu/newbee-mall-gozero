package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/index_config/model"

	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"newbee-mall-gozero/service/index_config/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigLogic {
	return &GetConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigLogic) GetConfig(in *indexconfig.GetConfigRequest) (*indexconfig.GetConfigResponse, error) {
	res, err := l.svcCtx.ConfigModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的Config")
			return nil, errors.New("不存在的Config")
		}
		logx.Error("Config获取失败" + err.Error())
		return nil, errors.New("Config获取失败" + err.Error())
	}
	// 复制
	var config indexconfig.Config
	err = copier.Copy(&config, res)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}
	config.CreateTime = res.CreateTime.Unix()
	config.UpdateTime = res.UpdateTime.Unix()

	return &indexconfig.GetConfigResponse{
		Config: &config,
	}, nil
}
