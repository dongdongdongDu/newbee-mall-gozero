package order_item

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallOrderItemModel = (*customTbNewbeeMallOrderItemModel)(nil)

type (
	// TbNewbeeMallOrderItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallOrderItemModel.
	TbNewbeeMallOrderItemModel interface {
		tbNewbeeMallOrderItemModel
		FindItemsByOrderId(ctx context.Context, orderId int64) ([]*TbNewbeeMallOrderItem, error)
	}

	customTbNewbeeMallOrderItemModel struct {
		*defaultTbNewbeeMallOrderItemModel
	}
)

func (m customTbNewbeeMallOrderItemModel) FindItemsByOrderId(ctx context.Context, orderId int64) ([]*TbNewbeeMallOrderItem, error) {
	var res []*TbNewbeeMallOrderItem
	query := fmt.Sprintf("select %s from %s where `order_id` = %d", tbNewbeeMallOrderItemRows, m.table, orderId)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	return res, nil
}

// NewTbNewbeeMallOrderItemModel returns a model for the database table.
func NewTbNewbeeMallOrderItemModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallOrderItemModel {
	return &customTbNewbeeMallOrderItemModel{
		defaultTbNewbeeMallOrderItemModel: newTbNewbeeMallOrderItemModel(conn, c),
	}
}
