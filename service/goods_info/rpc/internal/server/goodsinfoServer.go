// Code generated by goctl. DO NOT EDIT!
// Source: goodsinfo.proto

package server

import (
	"context"

	"newbee-mall-gozero/service/goods_info/rpc/goodsinfo"
	"newbee-mall-gozero/service/goods_info/rpc/internal/logic"
	"newbee-mall-gozero/service/goods_info/rpc/internal/svc"
)

type GoodsinfoServer struct {
	svcCtx *svc.ServiceContext
	goodsinfo.UnimplementedGoodsinfoServer
}

func NewGoodsinfoServer(svcCtx *svc.ServiceContext) *GoodsinfoServer {
	return &GoodsinfoServer{
		svcCtx: svcCtx,
	}
}

//  普通用户
func (s *GoodsinfoServer) GetGoodsDetail(ctx context.Context, in *goodsinfo.GetGoodsDetailRequest) (*goodsinfo.GetGoodsDetailResponse, error) {
	l := logic.NewGetGoodsDetailLogic(ctx, s.svcCtx)
	return l.GetGoodsDetail(in)
}

func (s *GoodsinfoServer) SearchGoods(ctx context.Context, in *goodsinfo.SearchGoodsRequest) (*goodsinfo.SearchGoodsResponse, error) {
	l := logic.NewSearchGoodsLogic(ctx, s.svcCtx)
	return l.SearchGoods(in)
}

//  管理员
func (s *GoodsinfoServer) AddGoodsInfo(ctx context.Context, in *goodsinfo.AddGoodsInfoRequest) (*goodsinfo.EmptyResponse, error) {
	l := logic.NewAddGoodsInfoLogic(ctx, s.svcCtx)
	return l.AddGoodsInfo(in)
}

func (s *GoodsinfoServer) UpdateGoodsInfo(ctx context.Context, in *goodsinfo.UpdateGoodsInfoRequest) (*goodsinfo.EmptyResponse, error) {
	l := logic.NewUpdateGoodsInfoLogic(ctx, s.svcCtx)
	return l.UpdateGoodsInfo(in)
}

func (s *GoodsinfoServer) AlterGoodsStatus(ctx context.Context, in *goodsinfo.AlterGoodsStatusRequest) (*goodsinfo.EmptyResponse, error) {
	l := logic.NewAlterGoodsStatusLogic(ctx, s.svcCtx)
	return l.AlterGoodsStatus(in)
}

func (s *GoodsinfoServer) DeleteGoodsInfo(ctx context.Context, in *goodsinfo.DeleteGoodsInfoRequest) (*goodsinfo.EmptyResponse, error) {
	l := logic.NewDeleteGoodsInfoLogic(ctx, s.svcCtx)
	return l.DeleteGoodsInfo(in)
}

func (s *GoodsinfoServer) GetGoodsInfo(ctx context.Context, in *goodsinfo.GetGoodsInfoRequest) (*goodsinfo.GetGoodsInfoResponse, error) {
	l := logic.NewGetGoodsInfoLogic(ctx, s.svcCtx)
	return l.GetGoodsInfo(in)
}

func (s *GoodsinfoServer) GetGoodsList(ctx context.Context, in *goodsinfo.GetGoodsListRequest) (*goodsinfo.GetGoodsListResponse, error) {
	l := logic.NewGetGoodsListLogic(ctx, s.svcCtx)
	return l.GetGoodsList(in)
}
