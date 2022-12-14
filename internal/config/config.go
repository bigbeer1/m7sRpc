package config

import (
	"fmt"
	"github.com/bigbeer1/m7sRpc/internal/server"
	"github.com/bigbeer1/m7sRpc/internal/svc"
	"github.com/bigbeer1/m7sRpc/m7sClient"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	. "m7s.live/engine/v4"
)

type RedisConfig struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
}

var plugin = InstallPlugin(&RedisConfig{})

func (c *RedisConfig) OnEvent(event any) {
	switch event.(type) {
	case FirstConfig:

		ctx := svc.NewServiceContext(*c)

		s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
			m7sClient.RegisterM7SServer(grpcServer, server.NewM7sServer(ctx))

			if c.Mode == service.DevMode || c.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		defer s.Stop()

		//rpc log
		fmt.Println(c.Name, c.ListenOn)
		s.Start()

	}
}
