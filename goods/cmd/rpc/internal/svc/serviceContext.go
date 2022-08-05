package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/internal/config"
	"github.com/zhovdawei/gozero-micro-service/goods/model"
)

type ServiceContext struct {
	Config          config.Config
	GoodsModel      model.GoodsModel
	GoodsAttrsModel model.GoodsAttrsModel
	GoodsSpecModel  model.GoodsSpecModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		GoodsModel:      model.NewGoodsModel(conn, c.CacheRedis),
		GoodsAttrsModel: model.NewGoodsAttrsModel(conn, c.CacheRedis),
		GoodsSpecModel:  model.NewGoodsSpecModel(conn, c.CacheRedis),
	}
}
