package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsSpecModel = (*customGoodsSpecModel)(nil)

var cacheGoGoodsGoodsSpecSpecPrefix = "cache:goGoods:goodsSpec:spec:"

type (
	// GoodsSpecModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsSpecModel.
	GoodsSpecModel interface {
		goodsSpecModel
		FindList(ctx context.Context, goodsId int64) ([]*GoodsSpec, error)
		QueryGoodsSpecBySpec(ctx context.Context, goodsId int64, specs []string) (*GoodsSpec, error)
	}

	customGoodsSpecModel struct {
		*defaultGoodsSpecModel
	}
)

// NewGoodsSpecModel returns a model for the database table.
func NewGoodsSpecModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsSpecModel {
	return &customGoodsSpecModel{
		defaultGoodsSpecModel: newGoodsSpecModel(conn, c),
	}
}

func (m *defaultGoodsSpecModel) FindList(ctx context.Context, goodsId int64) ([]*GoodsSpec, error) {
	query := "SELECT * FROM goods_spec WHERE goods_id = ?"
	var list []*GoodsSpec
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, goodsId)
	switch err {
	case nil:
		return list, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsSpecModel) QueryGoodsSpecBySpec(ctx context.Context, goodsId int64, specs []string) (*GoodsSpec, error) {
	//query := "select * from goods_spec where goods_id = 1 and json_contains(json_array('8GB+256GB','冰川蓝'), spec->'$') LIMIT 1"
	specVal := ""
	cacheKey := string(goodsId)
	for _, spec := range specs {
		specVal = specVal + "'" + spec + "',"
		cacheKey += spec
	}
	specVal = specVal[:len(specVal)-1]
	var resp GoodsSpec
	cacheGoGoodsGoodsSpecSpecPrefix += cacheKey
	err := m.QueryRowCtx(ctx, &resp, cacheGoGoodsGoodsSpecSpecPrefix, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select * from goods_spec where goods_id = ? and json_contains(json_array(%s), spec->'$') LIMIT 1", specVal)
		return conn.QueryRowCtx(ctx, v, query, goodsId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
	return nil, nil
}
