package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallGoodsCategoryModel = (*customTbNewbeeMallGoodsCategoryModel)(nil)

type (
	// TbNewbeeMallGoodsCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallGoodsCategoryModel.
	TbNewbeeMallGoodsCategoryModel interface {
		tbNewbeeMallGoodsCategoryModel
	}

	customTbNewbeeMallGoodsCategoryModel struct {
		*defaultTbNewbeeMallGoodsCategoryModel
	}
)

// NewTbNewbeeMallGoodsCategoryModel returns a model for the database table.
func NewTbNewbeeMallGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallGoodsCategoryModel {
	return &customTbNewbeeMallGoodsCategoryModel{
		defaultTbNewbeeMallGoodsCategoryModel: newTbNewbeeMallGoodsCategoryModel(conn, c),
	}
}
