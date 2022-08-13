package logic

import (
	"context"
	"newbee-mall-gozero/service/index_config/model"
	"time"

	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"newbee-mall-gozero/service/index_config/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigLogic {
	return &AddConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddConfigLogic) AddConfig(in *indexconfig.AddConfigRequest) (*indexconfig.EmptyResponse, error) {
	config := model.TbNewbeeMallIndexConfig{
		ConfigName:  in.ConfigName,
		ConfigType:  in.ConfigType,
		GoodsId:     in.GoodsId,
		RedirectUrl: in.RedirectUrl,
		ConfigRank:  in.ConfigRank,
		IsDeleted:   0,
		CreateTime:  time.Now(),
		CreateUser:  in.UserId,
		UpdateTime:  time.Now(),
		UpdateUser:  in.UserId,
	}
	_, err := l.svcCtx.ConfigModel.Insert(l.ctx, &config)
	if err != nil {
		logx.Error("添加失败：" + err.Error())
		return nil, err
	}

	return &indexconfig.EmptyResponse{}, nil
}
