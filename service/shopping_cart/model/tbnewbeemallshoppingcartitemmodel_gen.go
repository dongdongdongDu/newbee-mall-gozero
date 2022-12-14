// Code generated by goctl. DO NOT EDIT!

package model

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
	tbNewbeeMallShoppingCartItemFieldNames          = builder.RawFieldNames(&TbNewbeeMallShoppingCartItem{})
	tbNewbeeMallShoppingCartItemRows                = strings.Join(tbNewbeeMallShoppingCartItemFieldNames, ",")
	tbNewbeeMallShoppingCartItemRowsExpectAutoSet   = strings.Join(stringx.Remove(tbNewbeeMallShoppingCartItemFieldNames, "`cart_item_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tbNewbeeMallShoppingCartItemRowsWithPlaceHolder = strings.Join(stringx.Remove(tbNewbeeMallShoppingCartItemFieldNames, "`cart_item_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTbNewbeeMallShoppingCartItemCartItemIdPrefix = "cache:tbNewbeeMallShoppingCartItem:cartItemId:"
)

type (
	tbNewbeeMallShoppingCartItemModel interface {
		Insert(ctx context.Context, data *TbNewbeeMallShoppingCartItem) (sql.Result, error)
		FindOne(ctx context.Context, cartItemId int64) (*TbNewbeeMallShoppingCartItem, error)
		Update(ctx context.Context, data *TbNewbeeMallShoppingCartItem) error
		Delete(ctx context.Context, cartItemId int64) error
	}

	defaultTbNewbeeMallShoppingCartItemModel struct {
		sqlc.CachedConn
		table string
	}

	TbNewbeeMallShoppingCartItem struct {
		CartItemId int64     `db:"cart_item_id"` // 购物项主键id
		UserId     int64     `db:"user_id"`      // 用户主键id
		GoodsId    int64     `db:"goods_id"`     // 关联商品id
		GoodsCount int64     `db:"goods_count"`  // 数量(最大为5)
		IsDeleted  int64     `db:"is_deleted"`   // 删除标识字段(0-未删除 1-已删除)
		CreateTime time.Time `db:"create_time"`  // 创建时间
		UpdateTime time.Time `db:"update_time"`  // 最新修改时间
	}
)

func newTbNewbeeMallShoppingCartItemModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbNewbeeMallShoppingCartItemModel {
	return &defaultTbNewbeeMallShoppingCartItemModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_newbee_mall_shopping_cart_item`",
	}
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) Delete(ctx context.Context, cartItemId int64) error {
	tbNewbeeMallShoppingCartItemCartItemIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallShoppingCartItemCartItemIdPrefix, cartItemId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `cart_item_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, cartItemId)
	}, tbNewbeeMallShoppingCartItemCartItemIdKey)
	return err
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) FindOne(ctx context.Context, cartItemId int64) (*TbNewbeeMallShoppingCartItem, error) {
	tbNewbeeMallShoppingCartItemCartItemIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallShoppingCartItemCartItemIdPrefix, cartItemId)
	var resp TbNewbeeMallShoppingCartItem
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallShoppingCartItemCartItemIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `cart_item_id` = ? limit 1", tbNewbeeMallShoppingCartItemRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, cartItemId)
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

func (m *defaultTbNewbeeMallShoppingCartItemModel) Insert(ctx context.Context, data *TbNewbeeMallShoppingCartItem) (sql.Result, error) {
	tbNewbeeMallShoppingCartItemCartItemIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallShoppingCartItemCartItemIdPrefix, data.CartItemId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tbNewbeeMallShoppingCartItemRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.GoodsId, data.GoodsCount, data.IsDeleted)
	}, tbNewbeeMallShoppingCartItemCartItemIdKey)
	return ret, err
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) Update(ctx context.Context, data *TbNewbeeMallShoppingCartItem) error {
	tbNewbeeMallShoppingCartItemCartItemIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallShoppingCartItemCartItemIdPrefix, data.CartItemId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `cart_item_id` = ?", m.table, tbNewbeeMallShoppingCartItemRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.GoodsId, data.GoodsCount, data.IsDeleted, data.CartItemId)
	}, tbNewbeeMallShoppingCartItemCartItemIdKey)
	return err
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbNewbeeMallShoppingCartItemCartItemIdPrefix, primary)
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `cart_item_id` = ? limit 1", tbNewbeeMallShoppingCartItemRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) tableName() string {
	return m.table
}
