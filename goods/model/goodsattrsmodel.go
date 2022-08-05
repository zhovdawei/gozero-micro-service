package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsAttrsModel = (*customGoodsAttrsModel)(nil)

type (
	// GoodsAttrsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsAttrsModel.
	GoodsAttrsModel interface {
		goodsAttrsModel
		FindList(ctx context.Context, goodsId int64) ([]*GoodsAttrs, error)
	}

	customGoodsAttrsModel struct {
		*defaultGoodsAttrsModel
	}
)

// NewGoodsAttrsModel returns a model for the database table.
func NewGoodsAttrsModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsAttrsModel {
	return &customGoodsAttrsModel{
		defaultGoodsAttrsModel: newGoodsAttrsModel(conn, c),
	}
}

func (m *defaultGoodsAttrsModel) FindList(ctx context.Context, goodsId int64) ([]*GoodsAttrs, error) {
	query := "SELECT * FROM goods_attrs where goods_id = ?"
	var list []*GoodsAttrs
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, goodsId)
	switch err {
	case nil:
		return list, nil
	default:
		return nil, err
	}
}
