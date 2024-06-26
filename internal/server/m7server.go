// Code generated by goctl. DO NOT EDIT.
// Source: m7.proto

package server

import (
	"context"
	"github.com/bigbeer1/m7sRpc/internal/logic"
	"github.com/bigbeer1/m7sRpc/internal/svc"
	"github.com/bigbeer1/m7sRpc/m7client"
)

type M7Server struct {
	svcCtx *svc.ServiceContext
	m7client.UnimplementedM7Server
}

func NewM7Server(svcCtx *svc.ServiceContext) *M7Server {
	return &M7Server{
		svcCtx: svcCtx,
	}
}

// 用户登录
func (s *M7Server) Test(ctx context.Context, in *m7client.TestReq) (*m7client.TestResp, error) {
	l := logic.NewTestLogic(ctx, s.svcCtx)
	return l.Test(in)
}
