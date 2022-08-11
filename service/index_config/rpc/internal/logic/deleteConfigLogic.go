package logic

import (
	"context"
	"errors"
	"newbee-mall-gozero/service/index_config/model"

	"newbee-mall-gozero/service/index_config/rpc/indexconfig"
	"newbee-mall-gozero/service/index_config/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigLogic {
	return &DeleteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteConfigLogic) DeleteConfig(in *indexconfig.DeleteConfigRequest) (*indexconfig.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	// 遍历Ids修改IsDeleted值
	for _, id := range in.Ids {
		// 查找
		res, err := l.svcCtx.ConfigModel.FindOne(l.ctx, id)
		if err != nil {
			if err == model.ErrNotFound {
				logx.Error("不存在的Config")
				return nil, errors.New("不存在的Config")
			}
			logx.Error("Config获取失败" + err.Error())
			return nil, errors.New("Config获取失败" + err.Error())
		}
		res.IsDeleted = 1
		err = l.svcCtx.ConfigModel.Update(l.ctx, res)
		if err != nil {
			logx.Error("Config删除失败" + err.Error())
			return nil, errors.New("config删除失败")
		}
	}

	return &indexconfig.EmptyResponse{}, nil
}
