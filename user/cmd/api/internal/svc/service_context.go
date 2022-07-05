package svc

import (
	"github.com/zhovdawei/gozero-micro-service/user/cmd/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}