package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallUserModel = (*customTbNewbeeMallUserModel)(nil)

type (
	// TbNewbeeMallUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallUserModel.
	TbNewbeeMallUserModel interface {
		tbNewbeeMallUserModel
	}

	customTbNewbeeMallUserModel struct {
		*defaultTbNewbeeMallUserModel
	}
)

// NewTbNewbeeMallUserModel returns a model for the database table.
func NewTbNewbeeMallUserModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallUserModel {
	return &customTbNewbeeMallUserModel{
		defaultTbNewbeeMallUserModel: newTbNewbeeMallUserModel(conn, c),
	}
}
