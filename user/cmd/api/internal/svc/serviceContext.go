package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/config"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/middleware"
	"github.com/zhovdawei/gozero-micro-service/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	UserRpc        user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
