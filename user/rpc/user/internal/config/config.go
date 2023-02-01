package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	CacheRedis cache.CacheConf
	Mysql      struct {
		DataSource string
	}
	zrpc.RpcServerConf
}
