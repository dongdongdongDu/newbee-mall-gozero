package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbNewbeeMallGoodsInfoModel = (*customTbNewbeeMallGoodsInfoModel)(nil)

type (
	// TbNewbeeMallGoodsInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallGoodsInfoModel.
	TbNewbeeMallGoodsInfoModel interface {
		tbNewbeeMallGoodsInfoModel
		Search(ctx context.Context, pageNumber, goodsCategoryId int64, keyword, orderBy string) ([]*TbNewbeeMallGoodsInfo, int64, error)
		Find(ctx context.Context, name, status string, limit, offset int64) ([]*TbNewbeeMallGoodsInfo, error)
	}

	customTbNewbeeMallGoodsInfoModel struct {
		*defaultTbNewbeeMallGoodsInfoModel
	}
)

func (m *defaultTbNewbeeMallGoodsInfoModel) Search(ctx context.Context, pageNumber, goodsCategoryId int64, keyword, orderBy string) ([]*TbNewbeeMallGoodsInfo, int64, error) {
	var res []*TbNewbeeMallGoodsInfo
	query := fmt.Sprintf("select %s from %s", tbNewbeeMallGoodsInfoRows, m.table)
	var conditions string
	var order string
	if goodsCategoryId >= 0 && keyword != "" {
		conditions = fmt.Sprintf(" where `goods_category_id = %d and (goods_name like \"%%%s%%\" or goods_intro like \"%%%s%%\")", goodsCategoryId, keyword, keyword)
	} else if goodsCategoryId >= 0 && keyword == "" {
		conditions = fmt.Sprintf(" where `goods_category_id = %d", goodsCategoryId)
	} else if goodsCategoryId < 0 && keyword != "" {
		conditions = fmt.Sprintf(" where goods_name like \"%%%s%%\" or goods_intro like \"%%%s%%\"", keyword, keyword)
	}
	switch orderBy {
	case "new":
		order = " `goods_id` desc"
	case "price":
		order = " `selling_price` asc"
	default:
		order = " `stock_num` desc"
	}
	limit := 10
	offset := 10 * (pageNumber - 1)
	query = query + conditions + " order by" + order + " limit ? offset ?"
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, limit, offset)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	// 查询数量
	var total int64
	queryCount := fmt.Sprintf("select count(*) from %s", m.table)
	queryCount = queryCount + conditions
	err = m.QueryRowsNoCacheCtx(ctx, &total, query)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	return res, total, err
}

func (m *defaultTbNewbeeMallGoodsInfoModel) Find(ctx context.Context, name, status string, limit, offset int64) ([]*TbNewbeeMallGoodsInfo, error) {
	var res []*TbNewbeeMallGoodsInfo
	query := fmt.Sprintf("select %s from %s", tbNewbeeMallGoodsInfoRows, m.table)
	var conditions string
	if name != "" && status != "" {
		conditions = fmt.Sprintf(" where `goods_name` like \"%%%s%%\" and `goods_sell_status` = %s", name, status)
	} else if name != "" && status == "" {
		conditions = fmt.Sprintf(" where `goods_name` like \"%%%s%%\"", name)
	} else if name == "" && status != "" {
		conditions = fmt.Sprintf(" where `goods_sell_status` = %s", status)
	}
	query = query + conditions + " order by `goods_id` desc limit ? offset ?"
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, limit, offset)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}
	return res, nil
}

// NewTbNewbeeMallGoodsInfoModel returns a model for the database table.
func NewTbNewbeeMallGoodsInfoModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallGoodsInfoModel {
	return &customTbNewbeeMallGoodsInfoModel{
		defaultTbNewbeeMallGoodsInfoModel: newTbNewbeeMallGoodsInfoModel(conn, c),
	}
}
