package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallAdminUserTokenModel = (*customTbNewbeeMallAdminUserTokenModel)(nil)

type (
	// TbNewbeeMallAdminUserTokenModel is an interface.md to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallAdminUserTokenModel.
	TbNewbeeMallAdminUserTokenModel interface {
		tbNewbeeMallAdminUserTokenModel
	}

	customTbNewbeeMallAdminUserTokenModel struct {
		*defaultTbNewbeeMallAdminUserTokenModel
	}
)

// NewTbNewbeeMallAdminUserTokenModel returns a model for the database table.
func NewTbNewbeeMallAdminUserTokenModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallAdminUserTokenModel {
	return &customTbNewbeeMallAdminUserTokenModel{
		defaultTbNewbeeMallAdminUserTokenModel: newTbNewbeeMallAdminUserTokenModel(conn, c),
	}
}
