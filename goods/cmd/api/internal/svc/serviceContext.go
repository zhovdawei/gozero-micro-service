package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/config"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/api/internal/middleware"
	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/goods"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	GoodsRpc       goods.Goods
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		GoodsRpc:       goods.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
	}
}
