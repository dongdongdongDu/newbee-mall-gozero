package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"newbee-mall-gozero/common/nums"
)

var _ TbNewbeeMallGoodsCategoryModel = (*customTbNewbeeMallGoodsCategoryModel)(nil)

type (
	// TbNewbeeMallGoodsCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbNewbeeMallGoodsCategoryModel.
	TbNewbeeMallGoodsCategoryModel interface {
		tbNewbeeMallGoodsCategoryModel
		GetCategoryListByParentIdAndLevel(ctx context.Context, parentId, level string, limit, offset int64) ([]*TbNewbeeMallGoodsCategory, int64, error)
		GetCategoriesByParentIdsAndLevel(ctx context.Context, ids []int64, level int64, limit int64) ([]*TbNewbeeMallGoodsCategory, error)
	}

	customTbNewbeeMallGoodsCategoryModel struct {
		*defaultTbNewbeeMallGoodsCategoryModel
	}
)

func (m customTbNewbeeMallGoodsCategoryModel) GetCategoryListByParentIdAndLevel(ctx context.Context, parentId, level string, limit, offset int64) ([]*TbNewbeeMallGoodsCategory, int64, error) {
	var res []*TbNewbeeMallGoodsCategory
	query := fmt.Sprintf("select %s from %s where `is_deleted` = 0 ", tbNewbeeMallGoodsCategoryRows, m.table)
	var conditions string
	if parentId != "" && level != "" {
		conditions = fmt.Sprintf(" and `parent_id` = %s and `category_level` = %s", parentId, level)
	} else if parentId != "" && level == "" {
		conditions = fmt.Sprintf(" and `parent_id` = %s", parentId)
	} else if parentId == "" && level != "" {
		conditions = fmt.Sprintf(" and `category_level` = %s", level)
	}
	query = query + conditions + " order by `category_rank` desc"
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

func (m customTbNewbeeMallGoodsCategoryModel) GetCategoriesByParentIdsAndLevel(ctx context.Context, ids []int64, level int64, limit int64) ([]*TbNewbeeMallGoodsCategory, error) {
	var res []*TbNewbeeMallGoodsCategory
	if len(ids) == 0 {
		return res, nil
	}
	idstr := nums.IntToStr(ids)
	query := fmt.Sprintf("select %s from %s where `is_deleted` = 0", tbNewbeeMallGoodsCategoryRows, m.table)
	conditions := fmt.Sprintf(" and `parent_id` in %s and `category_level` = %d", idstr, level)
	query = query + conditions + " order by `category_rank` desc"
	//logx.Info("SQL:", query)
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}

	total := int64(len(res))
	if limit < total {
		res = res[:limit]
	}

	return res, nil
}

// NewTbNewbeeMallGoodsCategoryModel returns a model for the database table.
func NewTbNewbeeMallGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) TbNewbeeMallGoodsCategoryModel {
	return &customTbNewbeeMallGoodsCategoryModel{
		defaultTbNewbeeMallGoodsCategoryModel: newTbNewbeeMallGoodsCategoryModel(conn, c),
	}
}
