package svc

import (
	"github.com/cjstarcc/zeroframework/user/rpc/user/internal/config"
	"github.com/cjstarcc/zeroframework/user/rpc/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
		Config:    c,
	}
}
