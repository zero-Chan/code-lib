package redis

import (
	"code-lib/conf"
)

type RedisClientConf struct {
	Addr conf.AddrConf
}

type RedisClusterClientConf struct {
	Addrs []conf.AddrConf
}
