package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var _ TbNewbeeMallShoppingCartItemModel = (*customTbNewbeeMallShoppingCartItemModel)(nil)

type (
	// TbNewbeeMallShoppingCartItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallShoppingCartItemModel.
	TbNewbeeMallShoppingCartItemModel interface {
		tbNewbeeMallShoppingCartItemModel
		FindItem(ctx context.Context, goodsId, userId int64) ([]*TbNewbeeMallShoppingCartItem, error)
		FindItems(ctx context.Context, cartItemIds []int64, userId int64) ([]*TbNewbeeMallShoppingCartItem, error)
		FindNotDelItemById(ctx context.Context, cartItemId int64) (*TbNewbeeMallShoppingCartItem, error)
		FindNotDelItemByUser(ctx context.Context, userId int64) ([]*TbNewbeeMallShoppingCartItem, error)
		CountUserItemAll(ctx context.Context, userId int64) (int64, error)
	}

	customTbNewbeeMallShoppingCartItemModel struct {
		*defaultTbNewbeeMallShoppingCartItemModel
	}
)

func (m customTbNewbeeMallShoppingCartItemModel) FindItem(ctx context.Context, goodsId, userId int64) ([]*TbNewbeeMallShoppingCartItem, error) {
	var res []*TbNewbeeMallShoppingCartItem
	query := fmt.Sprintf("select %s from %s where `goods_id` = %d and `user_id` = %d and `is_deleted` = 0", tbNewbeeMallShoppingCartItemRows, m.table, goodsId, userId)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m customTbNewbeeMallShoppingCartItemModel) FindItems(ctx context.Context, cartItemIds []int64, userId int64) ([]*TbNewbeeMallShoppingCartItem, error) {
	var res []*TbNewbeeMallShoppingCartItem
	if len(cartItemIds) == 0 {
		return res, nil
	}
	idstr := "("
	for i := 0; i < len(cartItemIds); i++ {
		idstr += strconv.FormatInt(cartItemIds[i], 10)
		if i != len(cartItemIds)-1 {
			idstr += ","
		}
	}
	idstr += ")"

	query := fmt.Sprintf("select %s from %s", tbNewbeeMallShoppingCartItemRows, m.table)
	conditions := fmt.Sprintf(" where `cart_item_id` in %s and `user_id` = %d and `is_deleted` = 0", idstr, userId)
	query = query + conditions
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbNewbeeMallShoppingCartItemModel) FindNotDelItemById(ctx context.Context, cartItemId int64) (*TbNewbeeMallShoppingCartItem, error) {
	tbNewbeeMallShoppingCartItemCartItemIdKey := fmt.Sprintf("cache:tbNewbeeMallShoppingCartItemNotDel:cartItemId:%v", cartItemId)
	var resp TbNewbeeMallShoppingCartItem
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallShoppingCartItemCartItemIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `cart_item_id` = ? and `is_deleted`= 0 limit 1", tbNewbeeMallShoppingCartItemRows, m.table)
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

func (m *customTbNewbeeMallShoppingCartItemModel) FindNotDelItemByUser(ctx context.Context, userId int64) ([]*TbNewbeeMallShoppingCartItem, error) {
	var res []*TbNewbeeMallShoppingCartItem
	query := fmt.Sprintf("select %s from %s where `user_id` = %d and `is_deleted` = 0", tbNewbeeMallShoppingCartItemRows, m.table, userId)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTbNewbeeMallShoppingCartItemModel) CountUserItemAll(ctx context.Context, userId int64) (int64, error) {
	tbNewbeeMallShoppingCartItemAllKey := "cache:tbNewbeeMall:shoppingCartItems:countAll" + fmt.Sprintf(":%d", userId)
	var res int64
	err := m.QueryRowCtx(ctx, &res, tbNewbeeMallShoppingCartItemAllKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) from %s where `user_id` = %d and `is_deleted` = 0", m.table, userId)
		return conn.QueryRowCtx(ctx, v, query)
	})
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// NewTbNewbeeMallShoppingCartItemModel returns a model for the database table.
func NewTbNewbeeMallShoppingCartItemModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallShoppingCartItemModel {
	return &customTbNewbeeMallShoppingCartItemModel{
		defaultTbNewbeeMallShoppingCartItemModel: newTbNewbeeMallShoppingCartItemModel(conn, c),
	}
}
