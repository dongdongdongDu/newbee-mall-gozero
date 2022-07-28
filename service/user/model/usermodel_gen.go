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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheUserUserIdPrefix    = "cache:user:userId:"
	cacheUserLoginNamePrefix = "cache:user:loginName:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*User, error)
		FindOneByLoginName(ctx context.Context, loginName string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		UserId        int64     `db:"user_id"`
		NikeName      string    `db:"nike_name"`      // 用户昵称
		LoginName     string    `db:"login_name"`     // 登陆名称(默认为手机号)
		PasswordMd5   string    `db:"password_md5"`   // MD5加密后的密码
		IntroduceSign string    `db:"introduce_sign"` // 个性签名
		IsDeleted     int64     `db:"is_deleted"`     // 注销标识字段(0-正常 1-已注销)
		LockedFlag    int64     `db:"locked_flag"`    // 锁定标识字段(0-未锁定 1-已锁定)
		CreateTime    time.Time `db:"create_time"`    // 注册时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, userId int64) error {
	data, err := m.FindOne(ctx, userId)
	if err != nil {
		return err
	}

	userLoginNameKey := fmt.Sprintf("%s%v", cacheUserLoginNamePrefix, data.LoginName)
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, userId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, userLoginNameKey, userUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userId int64) (*User, error) {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, userId)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

func (m *defaultUserModel) FindOneByLoginName(ctx context.Context, loginName string) (*User, error) {
	userLoginNameKey := fmt.Sprintf("%s%v", cacheUserLoginNamePrefix, loginName)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userLoginNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `login_name` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, loginName); err != nil {
			return nil, err
		}
		return resp.UserId, nil
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userLoginNameKey := fmt.Sprintf("%s%v", cacheUserLoginNamePrefix, data.LoginName)
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.NikeName, data.LoginName, data.PasswordMd5, data.IntroduceSign, data.IsDeleted, data.LockedFlag)
	}, userLoginNameKey, userUserIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.UserId)
	if err != nil {
		return err
	}

	userLoginNameKey := fmt.Sprintf("%s%v", cacheUserLoginNamePrefix, data.LoginName)
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.NikeName, newData.LoginName, newData.PasswordMd5, newData.IntroduceSign, newData.IsDeleted, newData.LockedFlag, newData.UserId)
	}, userLoginNameKey, userUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
