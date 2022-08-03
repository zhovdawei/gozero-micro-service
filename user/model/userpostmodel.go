package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserPostModel = (*customUserPostModel)(nil)

type (
	// UserPostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPostModel.
	UserPostModel interface {
		userPostModel
		FindList(ctx context.Context, userId int64) ([]*UserPost, error)
	}

	customUserPostModel struct {
		*defaultUserPostModel
	}
)

// NewUserPostModel returns a model for the database table.
func NewUserPostModel(conn sqlx.SqlConn, c cache.CacheConf) UserPostModel {
	return &customUserPostModel{
		defaultUserPostModel: newUserPostModel(conn, c),
	}
}

func (m *defaultUserPostModel) FindList(ctx context.Context, userId int64) ([]*UserPost, error) {
	query := "select * from user_post where user_id = ?"
	var list []*UserPost
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, userId)
	switch err {
	case nil:
		return list, nil
	default:
		return nil, err
	}
}
