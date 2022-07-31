package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminTokenModel = (*customAdminTokenModel)(nil)

type (
	// AdminTokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminTokenModel.
	AdminTokenModel interface {
		adminTokenModel
	}

	customAdminTokenModel struct {
		*defaultAdminTokenModel
	}
)

// NewAdminTokenModel returns a model for the database table.
func NewAdminTokenModel(conn sqlx.SqlConn, c cache.CacheConf) AdminTokenModel {
	return &customAdminTokenModel{
		defaultAdminTokenModel: newAdminTokenModel(conn, c),
	}
}
