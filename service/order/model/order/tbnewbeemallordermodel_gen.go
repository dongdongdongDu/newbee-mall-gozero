// Code generated by goctl. DO NOT EDIT!

package order

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbNewbeeMallOrderFieldNames          = builder.RawFieldNames(&TbNewbeeMallOrder{})
	tbNewbeeMallOrderRows                = strings.Join(tbNewbeeMallOrderFieldNames, ",")
	tbNewbeeMallOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(tbNewbeeMallOrderFieldNames, "`order_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tbNewbeeMallOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(tbNewbeeMallOrderFieldNames, "`order_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTbNewbeeMallOrderOrderIdPrefix = "cache:tbNewbeeMallOrder:orderId:"
)

type (
	tbNewbeeMallOrderModel interface {
		Insert(ctx context.Context, data *TbNewbeeMallOrder) (sql.Result, error)
		FindOne(ctx context.Context, orderId int64) (*TbNewbeeMallOrder, error)
		Update(ctx context.Context, data *TbNewbeeMallOrder) error
		Delete(ctx context.Context, orderId int64) error
	}

	defaultTbNewbeeMallOrderModel struct {
		sqlc.CachedConn
		table string
	}

	TbNewbeeMallOrder struct {
		OrderId     int64        `db:"order_id"`     // 订单表主键id
		OrderNo     string       `db:"order_no"`     // 订单号
		UserId      int64        `db:"user_id"`      // 用户主键id
		TotalPrice  int64        `db:"total_price"`  // 订单总价
		PayStatus   int64        `db:"pay_status"`   // 支付状态:0.未支付,1.支付成功,-1:支付失败
		PayType     int64        `db:"pay_type"`     // 0.无 1.支付宝支付 2.微信支付
		PayTime     sql.NullTime `db:"pay_time"`     // 支付时间
		OrderStatus int64        `db:"order_status"` // 订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
		ExtraInfo   string       `db:"extra_info"`   // 订单body
		IsDeleted   int64        `db:"is_deleted"`   // 删除标识字段(0-未删除 1-已删除)
		CreateTime  time.Time    `db:"create_time"`  // 创建时间
		UpdateTime  time.Time    `db:"update_time"`  // 最新修改时间
	}
)

func newTbNewbeeMallOrderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbNewbeeMallOrderModel {
	return &defaultTbNewbeeMallOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_newbee_mall_order`",
	}
}

func (m *defaultTbNewbeeMallOrderModel) Delete(ctx context.Context, orderId int64) error {
	tbNewbeeMallOrderOrderIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallOrderOrderIdPrefix, orderId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `order_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, orderId)
	}, tbNewbeeMallOrderOrderIdKey)
	return err
}

func (m *defaultTbNewbeeMallOrderModel) FindOne(ctx context.Context, orderId int64) (*TbNewbeeMallOrder, error) {
	tbNewbeeMallOrderOrderIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallOrderOrderIdPrefix, orderId)
	var resp TbNewbeeMallOrder
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallOrderOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", tbNewbeeMallOrderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, orderId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbNewbeeMallOrderModel) Insert(ctx context.Context, data *TbNewbeeMallOrder) (sql.Result, error) {
	tbNewbeeMallOrderOrderIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallOrderOrderIdPrefix, data.OrderId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tbNewbeeMallOrderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.OrderNo, data.UserId, data.TotalPrice, data.PayStatus, data.PayType, data.PayTime, data.OrderStatus, data.ExtraInfo, data.IsDeleted)
	}, tbNewbeeMallOrderOrderIdKey)
	return ret, err
}

func (m *defaultTbNewbeeMallOrderModel) Update(ctx context.Context, data *TbNewbeeMallOrder) error {
	tbNewbeeMallOrderOrderIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallOrderOrderIdPrefix, data.OrderId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `order_id` = ?", m.table, tbNewbeeMallOrderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.OrderNo, data.UserId, data.TotalPrice, data.PayStatus, data.PayType, data.PayTime, data.OrderStatus, data.ExtraInfo, data.IsDeleted, data.OrderId)
	}, tbNewbeeMallOrderOrderIdKey)
	return err
}

func (m *defaultTbNewbeeMallOrderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbNewbeeMallOrderOrderIdPrefix, primary)
}

func (m *defaultTbNewbeeMallOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", tbNewbeeMallOrderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbNewbeeMallOrderModel) tableName() string {
	return m.table
}
