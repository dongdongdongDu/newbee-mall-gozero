package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallAdminUserModel = (*customTbNewbeeMallAdminUserModel)(nil)

type (
	// TbNewbeeMallAdminUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallAdminUserModel.
	TbNewbeeMallAdminUserModel interface {
		tbNewbeeMallAdminUserModel
	}

	customTbNewbeeMallAdminUserModel struct {
		*defaultTbNewbeeMallAdminUserModel
	}
)

// NewTbNewbeeMallAdminUserModel returns a model for the database table.
func NewTbNewbeeMallAdminUserModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallAdminUserModel {
	return &customTbNewbeeMallAdminUserModel{
		defaultTbNewbeeMallAdminUserModel: newTbNewbeeMallAdminUserModel(conn, c),
	}
}
