package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"newbee-mall-gozero/service/mqueue/job/jobtype"
	"newbee-mall-gozero/service/order/model/order_address"
	"newbee-mall-gozero/service/order/model/order_item"
	"time"

	"newbee-mall-gozero/common/nums"
	goodsInfo "newbee-mall-gozero/service/goods_info/model"
	orderModel "newbee-mall-gozero/service/order/model/order"
	"newbee-mall-gozero/service/order/rpc/internal/svc"
	"newbee-mall-gozero/service/order/rpc/order"
	"newbee-mall-gozero/service/shopping_cart/rpc/shoppingcart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOrderLogic) AddOrder(in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	if len(in.CartItemIds) == 0 {
		logx.Error("数据无效")
		return nil, errors.New("数据无效")
	}
	// 根据CartItemIds查询商品
	itemsRes, err := l.svcCtx.ShoppingCartRpc.GetCartItems(l.ctx, &shoppingcart.GetCartItemsRequest{
		CartItemIds: in.CartItemIds,
		UserId:      in.UserId,
	})
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	var goodsIds []int64
	for _, item := range itemsRes.CartList {
		goodsIds = append(goodsIds, item.GoodsId)
	}
	if len(goodsIds) == 0 {
		logx.Error("数据无效")
		return nil, errors.New("数据无效")
	}
	fmt.Println(goodsIds)
	// 根据goodsIds查询商品详情
	goodsInfos, err := l.svcCtx.GoodsInfoModel.GetForIndex(l.ctx, goodsIds)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	// 检查是否包含已下架商品
	for _, goods := range goodsInfos {
		if goods.GoodsSellStatus != 0 {
			logx.Error("有商品已下架，无法生成订单")
			return nil, errors.New("有商品已下架，无法生成订单")
		}
	}
	// 构造goodsMap [商品id]商品
	goodsMap := make(map[int64]*goodsInfo.TbNewbeeMallGoodsInfo)
	for _, goods := range goodsInfos {
		goodsMap[goods.GoodsId] = goods
	}
	// 判断商品库存
	for _, item := range itemsRes.CartList {
		// 查出的商品中不存在购物车中的这条关联商品数据，直接返回错误提醒
		goods, ok := goodsMap[item.GoodsId]
		if !ok {
			logx.Error("购物车数据异常！")
			return nil, errors.New("购物车数据异常！")
		}
		if item.GoodsCount > goods.StockNum {
			logx.Error("库存不足！")
			return nil, errors.New("库存不足！")
		}
	}
	// 删除购物项
	for _, id := range in.CartItemIds {
		_, err = l.svcCtx.ShoppingCartRpc.DeleteCartItem(l.ctx, &shoppingcart.DeleteCartItemRequest{
			Id:     id,
			UserId: in.UserId,
		})
		if err != nil {
			logx.Error("删除购物项失败！")
			return nil, errors.New("删除购物项失败！")
		}
	}
	// 修改库存
	for _, item := range itemsRes.CartList {
		// 查找商品
		goods, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, item.GoodsId)
		if err != nil {
			logx.Error("查询失败" + err.Error())
			return nil, errors.New("查询失败" + err.Error())
		}
		// 检查库存和销售状态
		if goods.StockNum < item.GoodsCount || goods.GoodsSellStatus != 0 {
			logx.Error("库存不足" + err.Error())
			return nil, errors.New("库存不足" + err.Error())
		}
		// 减库存
		goods.StockNum = goods.StockNum - item.GoodsCount
		err = l.svcCtx.GoodsInfoModel.Update(l.ctx, goods)
		if err != nil {
			logx.Error("减库存失败" + err.Error())
			return nil, errors.New("减库存失败" + err.Error())
		}
	}
	//生成订单号
	orderNo := nums.GenOrderNo()
	//总价
	var priceTotal int64
	for _, item := range itemsRes.CartList {
		priceTotal += item.GoodsCount * item.SellingPrice
	}
	if priceTotal < 1 {
		logx.Error("订单价格异常！")
		return nil, errors.New("订单价格异常！")
	}
	// 构造订单对象
	newOrder := orderModel.TbNewbeeMallOrder{
		OrderNo:    orderNo,
		UserId:     in.UserId,
		TotalPrice: priceTotal,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	// 插入数据
	sqlRes, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		logx.Error("订单入库失败！" + err.Error())
		return nil, errors.New("订单入库失败！" + err.Error())
	}
	// 查询地址
	address, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.AddressId)
	if err != nil {
		logx.Error("查询失败" + err.Error())
		return nil, errors.New("查询失败" + err.Error())
	}
	// 生成订单收货地址快照，并保存至数据库
	var newOrderAddress order_address.TbNewbeeMallOrderAddress
	err = copier.Copy(&newOrderAddress, address)
	if err != nil {
		logx.Error("地址复制失败" + err.Error())
		return nil, errors.New("地址复制失败" + err.Error())
	}
	newOrderAddress.OrderId, err = sqlRes.LastInsertId()
	if err != nil {
		logx.Error("订单最后插入id获取失败！" + err.Error())
		return nil, errors.New("订单最后插入id获取失败！" + err.Error())
	}
	// 插入数据
	_, err = l.svcCtx.OrderAddressModel.Insert(l.ctx, &newOrderAddress)
	if err != nil {
		logx.Error("订单地址入库失败！" + err.Error())
		return nil, errors.New("订单地址入库失败！" + err.Error())
	}
	//生成所有的订单项快照，并保存至数据库
	for _, item := range itemsRes.CartList {
		var orderItem order_item.TbNewbeeMallOrderItem
		err := copier.Copy(&orderItem, item)
		if err != nil {
			logx.Error("订单项复制失败" + err.Error())
			return nil, errors.New("订单项复制失败" + err.Error())
		}
		orderItem.OrderId, err = sqlRes.LastInsertId()
		if err != nil {
			logx.Error("订单最后插入id获取失败！" + err.Error())
			return nil, errors.New("订单最后插入id获取失败！" + err.Error())
		}
		orderItem.CreateTime = time.Now()
		// 插入数据
		_, err = l.svcCtx.OrderItemModel.Insert(l.ctx, &orderItem)
		if err != nil {
			logx.Error("订单项入库失败！" + err.Error())
			return nil, errors.New("订单项入库失败！" + err.Error())
		}
	}

	// 超时关闭订单任务
	payload, err := json.Marshal(jobtype.DeferCloseOrderPayload{OrderNo: orderNo})
	if err != nil {
		logx.Error("超时关闭订单" + orderNo + "json序列化失败！" + err.Error())
		return nil, errors.New("超时关闭订单" + orderNo + "json序列化失败！" + err.Error())
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseOrder, payload), asynq.ProcessIn(20*time.Minute))
	if err != nil {
		logx.Error("超时关闭订单" + orderNo + "任务加入队列失败！" + err.Error())
		return nil, errors.New("超时关闭订单" + orderNo + "任务加入队列失败！" + err.Error())
	}

	return &order.AddOrderResponse{
		OrderNo: orderNo,
	}, nil
}
