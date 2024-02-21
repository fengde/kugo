package svc

import (
	"time"
	"{{template}}/internal/config"
	"{{template}}/internal/middleware"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config        config.Config
	LogMiddleware rest.Middleware
	DB            sqlx.SqlConn
	Redis         *redis.Redis
	LocalCache    *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache, err := collection.NewCache(5 * time.Minute)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		LogMiddleware: middleware.NewLogMiddleware().Handle,
		DB:            sqlx.NewMysql(c.Mysql.Dsn),
		Redis:         redis.MustNewRedis(c.Redis),
		LocalCache:    cache,
	}
}
