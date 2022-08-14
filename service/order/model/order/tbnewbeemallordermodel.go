package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallOrderModel = (*customTbNewbeeMallOrderModel)(nil)

type (
	// TbNewbeeMallOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallOrderModel.
	TbNewbeeMallOrderModel interface {
		tbNewbeeMallOrderModel
		FindOneByNo(ctx context.Context, orderNo string) (*TbNewbeeMallOrder, error)
		FindByUserAndStatus(ctx context.Context, pageNumber, userId int64, status string) ([]*TbNewbeeMallOrder, int64, error)
		FindByNoAndStatus(ctx context.Context, orderNo, status string, limit, offset int64) ([]*TbNewbeeMallOrder, int64, error)
	}

	customTbNewbeeMallOrderModel struct {
		*defaultTbNewbeeMallOrderModel
	}
)

func (m customTbNewbeeMallOrderModel) FindOneByNo(ctx context.Context, orderNo string) (*TbNewbeeMallOrder, error) {
	tbNewbeeMallOrderOrderIdKey := fmt.Sprintf("cache:tbNewbeeMallOrder:orderNo:%s", orderNo)
	var resp TbNewbeeMallOrder
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallOrderOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `order_no` = ? and `is_deleted` = 0 limit 1", tbNewbeeMallOrderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, orderNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m customTbNewbeeMallOrderModel) FindByUserAndStatus(ctx context.Context, pageNumber, userId int64, status string) ([]*TbNewbeeMallOrder, int64, error) {
	var res []*TbNewbeeMallOrder
	query := fmt.Sprintf("select %s from %s", tbNewbeeMallOrderRows, m.table)
	conditions := fmt.Sprintf(" where `user_id` = %d and `is_deleted` = 0", userId)
	if status != "" {
		conditions = fmt.Sprintf(" and `order_status` = %s", status)
	}
	query = query + conditions
	logx.Info("SQL:", query)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		logx.Error("查找订单失败")
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	total := int64(len(res))
	var limit int64 = 5
	offset := 5 * (pageNumber - 1)
	if total <= offset {
		return nil, 0, errors.New("无效页码范围")
	}
	if total > offset && total < offset+limit {
		res = res[offset:]
	} else {
		res = res[offset : offset+limit]
	}

	return res, total, nil
}

func (m customTbNewbeeMallOrderModel) FindByNoAndStatus(ctx context.Context, orderNo, status string, limit, offset int64) ([]*TbNewbeeMallOrder, int64, error) {
	var res []*TbNewbeeMallOrder
	query := fmt.Sprintf("select %s from %s", tbNewbeeMallOrderRows, m.table)
	var conditions string
	if orderNo != "" && status != "" {
		conditions = fmt.Sprintf(" where `order_no` = %s and `order_status` = %s", orderNo, status)
	} else if orderNo == "" && status != "" {
		conditions = fmt.Sprintf(" where `order_status` = %s", status)
	} else if orderNo != "" && status == "" {
		conditions = fmt.Sprintf(" where `order_no` = %s", orderNo)
	}
	query = query + conditions + " order by `update_time` desc"
	logx.Info("SQL:", query)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		logx.Error("查找订单失败")
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	total := int64(len(res))
	if total <= offset {
		return nil, 0, errors.New("无效页码范围")
	}
	if total > offset && total < offset+limit {
		res = res[offset:]
	} else {
		res = res[offset : offset+limit]
	}

	return res, total, nil
}

// NewTbNewbeeMallOrderModel returns a model for the database table.
func NewTbNewbeeMallOrderModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallOrderModel {
	return &customTbNewbeeMallOrderModel{
		defaultTbNewbeeMallOrderModel: newTbNewbeeMallOrderModel(conn, c),
	}
}
