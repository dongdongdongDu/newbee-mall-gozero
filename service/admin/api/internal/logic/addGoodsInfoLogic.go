package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsInfoLogic {
	return &AddGoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodsInfoLogic) AddGoodsInfo(req *types.AddGoodsInfoRequest) (resp *types.Response, err error) {
	// 复制
	var goodsInfo goodsinfo.GoodsInfo
	err = copier.Copy(&goodsInfo, req)
	// 添加
	_, err = l.svcCtx.GoodsInfoRpc.AddGoodsInfo(l.ctx, &goodsinfo.AddGoodsInfoRequest{
		GoodsInfo: &goodsInfo,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "创建失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "创建成功",
		Data:       map[string]interface{}{},
	}, nil
}
