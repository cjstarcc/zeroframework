// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/cjstarcc/zeroframework/user/rpc/user/internal/logic"
	"github.com/cjstarcc/zeroframework/user/rpc/user/internal/svc"
	"github.com/cjstarcc/zeroframework/user/rpc/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.SCLoginReq) (*user.CSLoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
