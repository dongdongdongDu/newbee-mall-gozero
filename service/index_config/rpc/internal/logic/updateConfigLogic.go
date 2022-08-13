package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/index_config/model"
	"time"

	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"newbee-mall-gozero/service/index_config/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigLogic {
	return &UpdateConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigLogic) UpdateConfig(in *indexconfig.UpdateConfigRequest) (*indexconfig.EmptyResponse, error) {
	res, err := l.svcCtx.ConfigModel.FindOne(l.ctx, in.ConfigId)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的Config")
			return nil, errors.New("不存在的Config")
		}
		logx.Error("Config获取失败" + err.Error())
		return nil, errors.New("Config获取失败" + err.Error())
	}
	res.ConfigName = in.ConfigName
	res.ConfigType = in.ConfigType
	res.GoodsId = in.GoodsId
	res.RedirectUrl = in.RedirectUrl
	res.ConfigRank = in.ConfigRank
	res.UpdateUser = in.UserId
	res.UpdateTime = time.Now()

	err = l.svcCtx.ConfigModel.Update(l.ctx, res)
	if err != nil {
		logx.Error("Config更新失败" + err.Error())
		return nil, errors.New("Config更新失败")
	}

	return &indexconfig.EmptyResponse{}, nil
}
