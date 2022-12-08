package svc

import (
	"gitee.com/dabenxiong11/m7s-rpc"
)

type ServiceContext struct {
	Config m7sRpc.RedisConfig
}

func NewServiceContext(c m7sRpc.RedisConfig) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
