package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindUserByPhone(ctx context.Context, phone string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) FindUserByPhone(ctx context.Context, phone string) (*User, error) {
	query := "SELECT * FROM `user` where `phone` = ?"
	var user User
	err := m.QueryRowNoCache(&user, query, phone)
	switch err {
	case nil:
		return &user, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
