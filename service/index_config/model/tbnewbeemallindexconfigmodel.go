package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallIndexConfigModel = (*customTbNewbeeMallIndexConfigModel)(nil)

type (
	// TbNewbeeMallIndexConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallIndexConfigModel.
	TbNewbeeMallIndexConfigModel interface {
		tbNewbeeMallIndexConfigModel
		GetConfigList(ctx context.Context, configType, limit, offset int64) ([]*TbNewbeeMallIndexConfig, int64, error)
		GetConfigForIndex(ctx context.Context, configType, limit int64) ([]*TbNewbeeMallIndexConfig, error)
	}

	customTbNewbeeMallIndexConfigModel struct {
		*defaultTbNewbeeMallIndexConfigModel
	}
)

func (m customTbNewbeeMallIndexConfigModel) GetConfigList(ctx context.Context, configType, limit, offset int64) ([]*TbNewbeeMallIndexConfig, int64, error) {
	var res []*TbNewbeeMallIndexConfig
	if configType == 0 || configType > 5 {
		return nil, 0, errors.New("configType值不合法")
	}
	query := fmt.Sprintf("select %s from %s where `config_type` = %d ", tbNewbeeMallIndexConfigRows, m.table, configType)
	query = query + " order by `config_rank` desc"
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	total := int64(len(res))

	if offset >= total {
		return nil, 0, errors.New("查询页码超过最大范围")
	}
	if offset+limit > total {
		res = res[offset:]
	} else {
		res = res[offset : offset+limit]
	}

	return res, total, nil
}
func (m customTbNewbeeMallIndexConfigModel) GetConfigForIndex(ctx context.Context, configType, limit int64) ([]*TbNewbeeMallIndexConfig, error) {
	var res []*TbNewbeeMallIndexConfig
	if configType == 0 || configType > 5 {
		return nil, errors.New("configType值不合法")
	}
	query := fmt.Sprintf("select %s from %s where `is_deleted` = 0 and `config_type` = %d ", tbNewbeeMallIndexConfigRows, m.table, configType)
	query = query + " order by `config_rank` desc limit ?"
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, limit)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return res, nil
}

// NewTbNewbeeMallIndexConfigModel returns a model for the database table.
func NewTbNewbeeMallIndexConfigModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallIndexConfigModel {
	return &customTbNewbeeMallIndexConfigModel{
		defaultTbNewbeeMallIndexConfigModel: newTbNewbeeMallIndexConfigModel(conn, c),
	}
}
