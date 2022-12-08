package svc

import (
	"github.com/bigbeer1/m7sRpc/internal/config"
)

type ServiceContext struct {
	Config config.RedisConfig
}

func NewServiceContext(c config.RedisConfig) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
