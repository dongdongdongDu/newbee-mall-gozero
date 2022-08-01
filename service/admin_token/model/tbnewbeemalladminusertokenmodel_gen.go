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
	tbNewbeeMallAdminUserTokenFieldNames          = builder.RawFieldNames(&TbNewbeeMallAdminUserToken{})
	tbNewbeeMallAdminUserTokenRows                = strings.Join(tbNewbeeMallAdminUserTokenFieldNames, ",")
	tbNewbeeMallAdminUserTokenRowsExpectAutoSet   = strings.Join(stringx.Remove(tbNewbeeMallAdminUserTokenFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tbNewbeeMallAdminUserTokenRowsWithPlaceHolder = strings.Join(stringx.Remove(tbNewbeeMallAdminUserTokenFieldNames, "`admin_user_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTbNewbeeMallAdminUserTokenAdminUserIdPrefix = "cache:tbNewbeeMallAdminUserToken:adminUserId:"
	cacheTbNewbeeMallAdminUserTokenTokenPrefix       = "cache:tbNewbeeMallAdminUserToken:token:"
)

type (
	tbNewbeeMallAdminUserTokenModel interface {
		Insert(ctx context.Context, data *TbNewbeeMallAdminUserToken) (sql.Result, error)
		FindOne(ctx context.Context, adminUserId int64) (*TbNewbeeMallAdminUserToken, error)
		FindOneByToken(ctx context.Context, token string) (*TbNewbeeMallAdminUserToken, error)
		Update(ctx context.Context, data *TbNewbeeMallAdminUserToken) error
		Delete(ctx context.Context, adminUserId int64) error
	}

	defaultTbNewbeeMallAdminUserTokenModel struct {
		sqlc.CachedConn
		table string
	}

	TbNewbeeMallAdminUserToken struct {
		AdminUserId int64     `db:"admin_user_id"` // 用户主键id
		Token       string    `db:"token"`         // token值(32位字符串)
		UpdateTime  time.Time `db:"update_time"`   // 修改时间
		ExpireTime  time.Time `db:"expire_time"`   // token过期时间
	}
)

func newTbNewbeeMallAdminUserTokenModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTbNewbeeMallAdminUserTokenModel {
	return &defaultTbNewbeeMallAdminUserTokenModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tb_newbee_mall_admin_user_token`",
	}
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) Delete(ctx context.Context, adminUserId int64) error {
	data, err := m.FindOne(ctx, adminUserId)
	if err != nil {
		return err
	}

	tbNewbeeMallAdminUserTokenAdminUserIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenAdminUserIdPrefix, adminUserId)
	tbNewbeeMallAdminUserTokenTokenKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenTokenPrefix, data.Token)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `admin_user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, adminUserId)
	}, tbNewbeeMallAdminUserTokenAdminUserIdKey, tbNewbeeMallAdminUserTokenTokenKey)
	return err
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) FindOne(ctx context.Context, adminUserId int64) (*TbNewbeeMallAdminUserToken, error) {
	tbNewbeeMallAdminUserTokenAdminUserIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenAdminUserIdPrefix, adminUserId)
	var resp TbNewbeeMallAdminUserToken
	err := m.QueryRowCtx(ctx, &resp, tbNewbeeMallAdminUserTokenAdminUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `admin_user_id` = ? limit 1", tbNewbeeMallAdminUserTokenRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, adminUserId)
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

func (m *defaultTbNewbeeMallAdminUserTokenModel) FindOneByToken(ctx context.Context, token string) (*TbNewbeeMallAdminUserToken, error) {
	tbNewbeeMallAdminUserTokenTokenKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenTokenPrefix, token)
	var resp TbNewbeeMallAdminUserToken
	err := m.QueryRowIndexCtx(ctx, &resp, tbNewbeeMallAdminUserTokenTokenKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `token` = ? limit 1", tbNewbeeMallAdminUserTokenRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, token); err != nil {
			return nil, err
		}
		return resp.AdminUserId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) Insert(ctx context.Context, data *TbNewbeeMallAdminUserToken) (sql.Result, error) {
	tbNewbeeMallAdminUserTokenAdminUserIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenAdminUserIdPrefix, data.AdminUserId)
	tbNewbeeMallAdminUserTokenTokenKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenTokenPrefix, data.Token)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, tbNewbeeMallAdminUserTokenRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AdminUserId, data.Token, data.ExpireTime)
	}, tbNewbeeMallAdminUserTokenAdminUserIdKey, tbNewbeeMallAdminUserTokenTokenKey)
	return ret, err
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) Update(ctx context.Context, newData *TbNewbeeMallAdminUserToken) error {
	data, err := m.FindOne(ctx, newData.AdminUserId)
	if err != nil {
		return err
	}

	tbNewbeeMallAdminUserTokenAdminUserIdKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenAdminUserIdPrefix, data.AdminUserId)
	tbNewbeeMallAdminUserTokenTokenKey := fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenTokenPrefix, data.Token)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `admin_user_id` = ?", m.table, tbNewbeeMallAdminUserTokenRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Token, newData.ExpireTime, newData.AdminUserId)
	}, tbNewbeeMallAdminUserTokenAdminUserIdKey, tbNewbeeMallAdminUserTokenTokenKey)
	return err
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTbNewbeeMallAdminUserTokenAdminUserIdPrefix, primary)
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `admin_user_id` = ? limit 1", tbNewbeeMallAdminUserTokenRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTbNewbeeMallAdminUserTokenModel) tableName() string {
	return m.table
}
