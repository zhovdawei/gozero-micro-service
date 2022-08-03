package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserPostModel = (*customUserPostModel)(nil)

type (
	// UserPostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPostModel.
	UserPostModel interface {
		userPostModel
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
