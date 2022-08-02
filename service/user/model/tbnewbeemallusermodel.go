package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallUserModel = (*customTbNewbeeMallUserModel)(nil)

type (
	// TbNewbeeMallUserModel is an interface.md to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallUserModel.
	TbNewbeeMallUserModel interface {
		tbNewbeeMallUserModel
		TakePart(ctx context.Context, limit, offset int64) ([]*TbNewbeeMallUser, error)
		CountAll(ctx context.Context) (int64, error)
	}

	customTbNewbeeMallUserModel struct {
		*defaultTbNewbeeMallUserModel
	}
)

func (m *customTbNewbeeMallUserModel) TakePart(ctx context.Context, limit, offset int64) ([]*TbNewbeeMallUser, error) {
	var resp []*TbNewbeeMallUser
	query := fmt.Sprintf("select %s from %s order by `create_time` desc limit ? offset ?", tbNewbeeMallUserRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, limit, offset)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customTbNewbeeMallUserModel) CountAll(ctx context.Context) (int64, error) {
	tbNewbeeMallUserAllKey := "cache:tbNewbeeMallUser:users:countAll"
	var resp int64
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallUserAllKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) from %s", m.table)
		return conn.QueryRowCtx(ctx, v, query)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// NewTbNewbeeMallUserModel returns a model for the database table.
func NewTbNewbeeMallUserModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallUserModel {
	return &customTbNewbeeMallUserModel{
		defaultTbNewbeeMallUserModel: newTbNewbeeMallUserModel(conn, c),
	}
}
