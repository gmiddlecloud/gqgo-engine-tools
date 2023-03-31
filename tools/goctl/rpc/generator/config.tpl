package config

import (
{{if .isEnt}}   "github.com/gmiddlecloud/gqgo-engine-common/config"
{{end}}
    "github.com/zeromicro/go-zero/core/stores/redis"
    "github.com/zeromicro/go-zero/zrpc"

)

type Config struct {
	zrpc.RpcServerConf
{{if .isEnt}}   DatabaseConf config.DatabaseConf
    RedisConf    redis.RedisConf
{{end}}
}

