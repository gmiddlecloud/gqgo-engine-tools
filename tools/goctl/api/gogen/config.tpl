package config

import (
    {{if .useCasbin}}"gqgo-engine-common/plugins/casbin"
    "gqgo-engine-common/config"
    "github.com/zeromicro/go-zero/core/stores/redis"{{end}}
    "github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth   rest.AuthConf
	{{if .useCasbin}}DatabaseConf config.DatabaseConf
    RedisConf    redis.RedisConf
	CasbinConf   casbin.CasbinConf{{end}}
	{{.jwtTrans}}
}
