package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
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
		Find(ctx context.Context, name, status string, limit, offset int64) ([]*TbNewbeeMallGoodsInfo, int64, error)
		CountAll(ctx context.Context) (int64, error)
	}

	customTbNewbeeMallGoodsInfoModel struct {
		*defaultTbNewbeeMallGoodsInfoModel
	}
)

func (m *customTbNewbeeMallGoodsInfoModel) Search(ctx context.Context, pageNumber, goodsCategoryId int64, keyword, orderBy string) ([]*TbNewbeeMallGoodsInfo, int64, error) {
	var res []*TbNewbeeMallGoodsInfo
	query := fmt.Sprintf("select %s from %s", tbNewbeeMallGoodsInfoRows, m.table)
	var conditions string
	var order string
	if goodsCategoryId >= 0 && keyword != "" {
		conditions = fmt.Sprintf(" where `goods_category_id` = %d and (`goods_name` like \"%%%s%%\" or `goods_intro` like \"%%%s%%\")", goodsCategoryId, keyword, keyword)
	} else if goodsCategoryId >= 0 && keyword == "" {
		conditions = fmt.Sprintf(" where `goods_category_id` = %d", goodsCategoryId)
	} else if goodsCategoryId < 0 && keyword != "" {
		conditions = fmt.Sprintf(" where `goods_name` like \"%%%s%%\" or `goods_intro` like \"%%%s%%\"", keyword, keyword)
	}
	switch orderBy {
	case "new":
		order = " `goods_id` desc"
	case "price":
		order = " `selling_price` asc"
	default:
		order = " `stock_num` desc"
	}

	query = query + conditions + " order by" + order
	logx.Info("SQL:", query)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		logx.Error("查找GoodsInfo失败")
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	total := int64(len(res))

	var limit int64 = 10
	offset := 10 * (pageNumber - 1)
	if offset+limit > total {
		res = res[offset:]
	} else {
		res = res[offset : offset+limit]
	}

	return res, total, err
}

func (m *customTbNewbeeMallGoodsInfoModel) CountAll(ctx context.Context) (int64, error) {
	tbNewbeeMallGoodsInfoAllKey := "cache:tbNewbeeMall:goodsInfo:countAll"
	var res int64
	err := m.QueryRowCtx(ctx, &res, tbNewbeeMallGoodsInfoAllKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) from %s", m.table)
		return conn.QueryRowCtx(ctx, v, query)
	})
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customTbNewbeeMallGoodsInfoModel) Find(ctx context.Context, name, status string, limit, offset int64) ([]*TbNewbeeMallGoodsInfo, int64, error) {
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
	//query = query + conditions + " order by `goods_id` desc limit ? offset ?"
	//err := m.QueryRowsNoCacheCtx(ctx, &res, query, limit, offset)
	query = query + conditions + " order by `goods_id` desc"
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, 0, ErrNotFound
		} else {
			return nil, 0, err
		}
	}
	total := int64(len(res))

	if offset >= total {
		return nil, 0, errors.New("查询页码超过最大范围")
	}
	if offset+limit > total {
		res = res[offset:]
	} else {
		res = res[offset : offset+limit]
	}

	return res, total, nil
}

// NewTbNewbeeMallGoodsInfoModel returns a model for the database table.
func NewTbNewbeeMallGoodsInfoModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallGoodsInfoModel {
	return &customTbNewbeeMallGoodsInfoModel{
		defaultTbNewbeeMallGoodsInfoModel: newTbNewbeeMallGoodsInfoModel(conn, c),
	}
}
