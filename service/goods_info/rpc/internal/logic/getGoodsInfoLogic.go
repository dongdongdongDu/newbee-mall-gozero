package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/goods_info/model"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsInfoLogic {
	return &GetGoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsInfoLogic) GetGoodsInfo(in *goodsinfo.GetGoodsInfoRequest) (*goodsinfo.GetGoodsInfoResponse, error) {
	res, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			logx.Error("不存在的商品")
			return nil, errors.New("不存在的商品")
		}
		logx.Error("商品信息获取失败" + err.Error())
		return nil, errors.New("商品信息获取失败" + err.Error())
	}
	// 复制
	var goodsInfo goodsinfo.GoodsInfo
	err = copier.Copy(&goodsInfo, res)
	if err != nil {
		logx.Error("复制失败" + err.Error())
		return nil, errors.New("复制失败" + err.Error())
	}

	return &goodsinfo.GetGoodsInfoResponse{
		GoodsInfo: &goodsInfo,
	}, nil
}
