package logic

import (
	"context"
	"github.com/bigbeer1/m7sRpc/internal/svc"
	"github.com/bigbeer1/m7sRpc/m7client"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *TestLogic) Test(in *m7client.TestReq) (*m7client.TestResp, error) {

	return &m7client.TestResp{
		Message: in.Message + "11111",
	}, nil
}
