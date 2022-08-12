package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallUserAddressModel = (*customTbNewbeeMallUserAddressModel)(nil)

type (
	// TbNewbeeMallUserAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallUserAddressModel.
	TbNewbeeMallUserAddressModel interface {
		tbNewbeeMallUserAddressModel
		FindDefault(ctx context.Context, userId int64) (*TbNewbeeMallUserAddress, error)
		FindMyAddress(ctx context.Context, userId int64) ([]*TbNewbeeMallUserAddress, error)
	}

	customTbNewbeeMallUserAddressModel struct {
		*defaultTbNewbeeMallUserAddressModel
	}
)

func (m customTbNewbeeMallUserAddressModel) FindDefault(ctx context.Context, userId int64) (*TbNewbeeMallUserAddress, error) {
	tbNewbeeMallUserAddressAddressIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallUserAddressAddressIdPrefix, userId)
	var resp TbNewbeeMallUserAddress
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallUserAddressAddressIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `default_flag` =1 and `is_deleted` = 0 limit 1", tbNewbeeMallUserAddressRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

func (m customTbNewbeeMallUserAddressModel) FindMyAddress(ctx context.Context, userId int64) ([]*TbNewbeeMallUserAddress, error) {
	var res []*TbNewbeeMallUserAddress
	query := fmt.Sprintf("select %s from %s where `user_id` = %d and `is_deleted` = 0", tbNewbeeMallUserAddressRows, m.table, userId)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return res, nil
}

// NewTbNewbeeMallUserAddressModel returns a model for the database table.
func NewTbNewbeeMallUserAddressModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallUserAddressModel {
	return &customTbNewbeeMallUserAddressModel{
		defaultTbNewbeeMallUserAddressModel: newTbNewbeeMallUserAddressModel(conn, c),
	}
}
