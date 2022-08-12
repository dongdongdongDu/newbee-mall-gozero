package goodsInfo

import (
	"context"
	"newbee-mall-gozero/common/response"
	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"

	"newbee-mall-gozero/service/admin/api/internal/svc"
	"newbee-mall-gozero/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlterGoodsStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlterGoodsStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlterGoodsStatusLogic {
	return &AlterGoodsStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlterGoodsStatusLogic) AlterGoodsStatus(req *types.AlterGoodsStatusRequest) (resp *types.Response, err error) {
	// 添加
	_, err = l.svcCtx.GoodsInfoRpc.AlterGoodsStatus(l.ctx, &goodsinfo.AlterGoodsStatusRequest{
		Ids:    req.Ids,
		Status: req.Status,
	})
	if err != nil {
		return &types.Response{
			ResultCode: response.ERROR,
			Msg:        "修改商品状态失败！" + err.Error(),
		}, nil
	}

	return &types.Response{
		ResultCode: response.SUCCESS,
		Msg:        "修改商品状态成功",
		Data:       map[string]interface{}{},
	}, nil
}
