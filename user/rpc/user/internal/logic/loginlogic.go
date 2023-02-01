package logic

import (
	"context"
	"github.com/cjstarcc/zeroframework/user/rpc/user/internal/svc"
	"github.com/cjstarcc/zeroframework/user/rpc/user/user"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.SCLoginReq) (*user.CSLoginResp, error) {
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		logx.Info("db error!")
		return nil, err
	}
	if res.Password == in.Password {
		return &user.CSLoginResp{Id: res.Id}, nil
	}

	return nil, errors.New("密码错误")

}
