// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbNewbeeMallGoodsCategoryFieldNames          = builder.RawFieldNames(&TbNewbeeMallGoodsCategory{})
	tbNewbeeMallGoodsCategoryRows                = strings.Join(tbNewbeeMallGoodsCategoryFieldNames, ",")
	tbNewbeeMallGoodsCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(tbNewbeeMallGoodsCategoryFieldNames, "`category_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tbNewbeeMallGoodsCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(tbNewbeeMallGoodsCategoryFieldNames, "`category_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTbNewbeeMallGoodsCategoryCategoryIdPrefix = "cache:tbNewbeeMallGoodsCategory:categoryId:"
)

type (
	tbNewbeeMallGoodsCategoryModel interface {
		Insert(ctx context.Context, data *TbNewbeeMallGoodsCategory) (sql.Result, error)
		FindOne(ctx context.Context, categoryId int64) (*TbNewbeeMallGoodsCategory, error)
		Update(ctx context.Context, data *TbNewbeeMallGoodsCategory) error
		Delete(ctx context.Context, categoryId int64) error
	}

	defaultTbNewbeeMallGoodsCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	TbNewbeeMallGoodsCategory struct {
		CategoryId    int64     `db:"category_id"`    // 分类id
		CategoryLevel int64     `db:"category_level"` // 分类级别(1-一级分类 2-二级分类 3-三级分类)
		ParentId      int64     `db:"parent_id"`      // 父分类id
		CategoryName  string    `db:"category_name"`  // 分类名称
		CategoryRank  int64     `db:"category_rank"`  // 排序值(字段越大越靠前)
		IsDeleted     int64     `db:"is_deleted"`     // 删除标识字段(0-未删除 1-已删除)
		CreateTime    time.Time `db:"create_time"`    // 创建时间
		CreateUser    int64     `db:"create_user"`    // 创建者id
		UpdateTime    time.Time `db:"update_time"`    // 修改时间
		UpdateUser    int64     `db:"update_user"`    // 修改者id
	}
)

func newTbNewbeeMallGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbNewbeeMallGoodsCategoryModel {
	return &defaultTbNewbeeMallGoodsCategoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_newbee_mall_goods_category`",
	}
}

func (m *defaultTbNewbeeMallGoodsCategoryModel) Delete(ctx context.Context, categoryId int64) error {
	tbNewbeeMallGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallGoodsCategoryCategoryIdPrefix, categoryId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `category_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, categoryId)
	}, tbNewbeeMallGoodsCategoryCategoryIdKey)
	return err
}

func (m *defaultTbNewbeeMallGoodsCategoryModel) FindOne(ctx context.Context, categoryId int64) (*TbNewbeeMallGoodsCategory, error) {
	tbNewbeeMallGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallGoodsCategoryCategoryIdPrefix, categoryId)
	var resp TbNewbeeMallGoodsCategory
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallGoodsCategoryCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `category_id` = ? limit 1", tbNewbeeMallGoodsCategoryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, categoryId)
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

func (m *defaultTbNewbeeMallGoodsCategoryModel) Insert(ctx context.Context, data *TbNewbeeMallGoodsCategory) (sql.Result, error) {
	tbNewbeeMallGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallGoodsCategoryCategoryIdPrefix, data.CategoryId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tbNewbeeMallGoodsCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CategoryLevel, data.ParentId, data.CategoryName, data.CategoryRank, data.IsDeleted, data.CreateUser, data.UpdateUser)
	}, tbNewbeeMallGoodsCategoryCategoryIdKey)
	return ret, err
}

func (m *defaultTbNewbeeMallGoodsCategoryModel) Update(ctx context.Context, data *TbNewbeeMallGoodsCategory) error {
	tbNewbeeMallGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallGoodsCategoryCategoryIdPrefix, data.CategoryId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `category_id` = ?", m.table, tbNewbeeMallGoodsCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CategoryLevel, data.ParentId, data.CategoryName, data.CategoryRank, data.IsDeleted, data.CreateUser, data.UpdateUser, data.CategoryId)
	}, tbNewbeeMallGoodsCategoryCategoryIdKey)
	return err
}

func (m *defaultTbNewbeeMallGoodsCategoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbNewbeeMallGoodsCategoryCategoryIdPrefix, primary)
}

func (m *defaultTbNewbeeMallGoodsCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `category_id` = ? limit 1", tbNewbeeMallGoodsCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbNewbeeMallGoodsCategoryModel) tableName() string {
	return m.table
}