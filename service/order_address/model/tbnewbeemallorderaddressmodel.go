package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallOrderAddressModel = (*customTbNewbeeMallOrderAddressModel)(nil)

type (
	// TbNewbeeMallOrderAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallOrderAddressModel.
	TbNewbeeMallOrderAddressModel interface {
		tbNewbeeMallOrderAddressModel
	}

	customTbNewbeeMallOrderAddressModel struct {
		*defaultTbNewbeeMallOrderAddressModel
	}
)

// NewTbNewbeeMallOrderAddressModel returns a model for the database table.
func NewTbNewbeeMallOrderAddressModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallOrderAddressModel {
	return &customTbNewbeeMallOrderAddressModel{
		defaultTbNewbeeMallOrderAddressModel: newTbNewbeeMallOrderAddressModel(conn, c),
	}
}
