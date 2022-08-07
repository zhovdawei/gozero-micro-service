package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsModel = (*customGoodsModel)(nil)

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
		FindListByCategory(ctx context.Context, category string) ([]*Goods, error)
		FindListByGoodsName(ctx context.Context, goodsName string) ([]*Goods, error)
		FindListGoods(ctx context.Context, category string, goodsName string, lastId int64, size int32) ([]*Goods, error)
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn sqlx.SqlConn, c cache.CacheConf) GoodsModel {
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn, c),
	}
}

func (m *defaultGoodsModel) FindListByCategory(ctx context.Context, category string) ([]*Goods, error) {
	query := "SELECT * FROM goods WHERE category = ?"
	var goodsList []*Goods
	err := m.QueryRowsNoCacheCtx(ctx, &goodsList, query, category)
	switch err {
	case nil:
		return goodsList, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsModel) FindListByGoodsName(ctx context.Context, goodsName string) ([]*Goods, error) {
	goodsName = "'%" + goodsName + "%'"
	query := "SELECT * FROM goods WHERE goods_name like ?"
	var goodsList []*Goods
	err := m.QueryRowsNoCacheCtx(ctx, &goodsList, query, goodsName)
	switch err {
	case nil:
		return goodsList, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsModel) FindListGoods(ctx context.Context, category string, goodsName string, lastId int64, size int32) ([]*Goods, error) {
	if category != "" {
		category = "and category = '" + category + "'"
	}
	if goodsName != "" {
		goodsName = "and goods_name like '%" + goodsName + "%'"
	}
	query := fmt.Sprintf("SELECT * FROM goods where goods_id > ? %s %s ORDER BY goods_id LIMIT ?", category, goodsName)
	var goodsList []*Goods
	err := m.QueryRowsNoCacheCtx(ctx, &goodsList, query, lastId, size)
	switch err {
	case nil:
		return goodsList, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
