package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallCarouselModel = (*customTbNewbeeMallCarouselModel)(nil)

type (
	// TbNewbeeMallCarouselModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallCarouselModel.
	TbNewbeeMallCarouselModel interface {
		tbNewbeeMallCarouselModel
		TakePart(ctx context.Context, limit, offset int64) ([]*TbNewbeeMallCarousel, error)
		CountAll(ctx context.Context) (int64, error)
		GetForIndex(ctx context.Context, limit int64) ([]*TbNewbeeMallCarousel, error)
	}

	customTbNewbeeMallCarouselModel struct {
		*defaultTbNewbeeMallCarouselModel
	}
)

func (m customTbNewbeeMallCarouselModel) TakePart(ctx context.Context, limit, offset int64) ([]*TbNewbeeMallCarousel, error) {
	var res []*TbNewbeeMallCarousel
	query := fmt.Sprintf("select %s from %s order by `carousel_rank` desc limit ? offset ?", tbNewbeeMallCarouselRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, limit, offset)
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customTbNewbeeMallCarouselModel) CountAll(ctx context.Context) (int64, error) {
	tbNewbeeMallCarouselAllKey := "cache:tbNewbeeMallCarousel:carousels:countAll"
	var res int64
	err := m.QueryRowCtx(ctx, &res, tbNewbeeMallCarouselAllKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) from %s", m.table)
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

func (m customTbNewbeeMallCarouselModel) GetForIndex(ctx context.Context, limit int64) ([]*TbNewbeeMallCarousel, error) {
	var res []*TbNewbeeMallCarousel
	query := fmt.Sprintf("select %s from %s where `is_deleted` = 0 order by `carousel_rank` desc limit ?", tbNewbeeMallCarouselRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, limit)
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewTbNewbeeMallCarouselModel returns a model for the database table.
func NewTbNewbeeMallCarouselModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallCarouselModel {
	return &customTbNewbeeMallCarouselModel{
		defaultTbNewbeeMallCarouselModel: newTbNewbeeMallCarouselModel(conn, c),
	}
}
