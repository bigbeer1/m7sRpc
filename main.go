package m7sRpc

import (
	"fmt"
	"github.com/bigbeer1/m7sRpc/internal/config"
	"github.com/bigbeer1/m7sRpc/internal/server"
	"github.com/bigbeer1/m7sRpc/internal/svc"
	"github.com/bigbeer1/m7sRpc/m7client"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	. "m7s.live/engine/v4"
	"time"
)

type M7sRpcConfig struct {
	Name     string
	ListenOn string
	Timeout  int64 `json:",default=2000"`
}

var plugin = InstallPlugin(&M7sRpcConfig{})

func (m M7sRpcConfig) OnEvent(event any) {

	switch event.(type) {
	case FirstConfig:
		var c config.Config

		_ = copier.Copy(&c, m)
		ctx := svc.NewServiceContext(c)

		s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
			m7client.RegisterM7Server(grpcServer, server.NewM7Server(ctx))

			if c.Mode == service.DevMode || c.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		defer s.Stop()

		// 设置日志输出 接口慢时间  rpc
		zrpc.SetServerSlowThreshold(time.Second * 2000)

		//rpc log
		//s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

		fmt.Println(c.Name, c.ListenOn)
		s.Start()

	}
}
