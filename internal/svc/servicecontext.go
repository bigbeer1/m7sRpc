package svc

import (
	"github.com/bigbeer1/m7sRpc"
)

type ServiceContext struct {
	Config m7sRpc.RedisConfig
}

func NewServiceContext(c m7sRpc.RedisConfig) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
