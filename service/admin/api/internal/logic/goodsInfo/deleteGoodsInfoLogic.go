package goodsInfo

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsInfoLogic {
	return &DeleteGoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGoodsInfoLogic) DeleteGoodsInfo(req *types.DeleteGoodsInfoRequest) (resp *types.Response, err error) {
	// 删除
	_, err = l.svcCtx.GoodsInfoRpc.DeleteGoodsInfo(l.ctx, &goodsinfo.DeleteGoodsInfoRequest{
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "删除失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "删除成功",
		Data:       map[string]interface{}{},
	}, nil
}
