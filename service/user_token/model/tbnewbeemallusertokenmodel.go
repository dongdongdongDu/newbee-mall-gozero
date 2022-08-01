package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallUserTokenModel = (*customTbNewbeeMallUserTokenModel)(nil)

type (
	// TbNewbeeMallUserTokenModel is an interface.md to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallUserTokenModel.
	TbNewbeeMallUserTokenModel interface {
		tbNewbeeMallUserTokenModel
	}

	customTbNewbeeMallUserTokenModel struct {
		*defaultTbNewbeeMallUserTokenModel
	}
)

// NewTbNewbeeMallUserTokenModel returns a model for the database table.
func NewTbNewbeeMallUserTokenModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallUserTokenModel {
	return &customTbNewbeeMallUserTokenModel{
		defaultTbNewbeeMallUserTokenModel: newTbNewbeeMallUserTokenModel(conn, c),
	}
}
