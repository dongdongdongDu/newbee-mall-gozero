package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"newbee-mall-gozero/service/index_config/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigListLogic {
	return &GetConfigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigListLogic) GetConfigList(in *indexconfig.GetConfigListRequest) (*indexconfig.GetConfigListResponse, error) {
	limit := in.PageSize
	offset := in.PageSize * (in.PageNumber - 1)
	list, total, err := l.svcCtx.ConfigModel.GetConfigList(l.ctx, in.ConfigType, limit, offset)
	if err != nil {
		logx.Error("分类列表获取失败" + err.Error())
		return nil, errors.New("分类列表获取失败" + err.Error())
	}
	// 构造列表
	configList := make([]*indexconfig.Config, 0)
	for _, item := range list {
		var c indexconfig.Config
		err = copier.Copy(&c, item)
		configList = append(configList, &c)
	}

	return &indexconfig.GetConfigListResponse{
		ConfigList: configList,
		Total:      total,
	}, nil
}
