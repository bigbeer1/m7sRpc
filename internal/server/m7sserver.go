// Code generated by goctl. DO NOT EDIT!
// Source: m7s.proto

package server

import (
	"github.com/bigbeer1/m7sRpc/internal/svc"
)

type M7sServer struct {
	svcCtx *svc.ServiceContext
}

func NewM7sServer(svcCtx *svc.ServiceContext) *M7sServer {
	return &M7sServer{
		svcCtx: svcCtx,
	}
}
