package redis

import (
	"fmt"

	"gopkg.in/redis.v5"
)

type RedisClusterClient struct {
	cfg *RedisClusterClientConf

	*redis.ClusterClient
}

func NewRedisClusterClient(cfg *RedisClusterClientConf) (cli *RedisClusterClient, err error) {
	ErrorPrefix := "[InitError] `Func: NewRedisClusterClient` "

	if cfg == nil {
		err = fmt.Errorf(ErrorPrefix + "`Reason: cfg is nil.`")
		return
	}

	addrs := make([]string, len(cfg.Addrs))
	for idx, addr := range cfg.Addrs {
		addrs[idx] = addr.String()
	}

	rCli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})

	statusCmd := rCli.Ping()
	if rerr := statusCmd.Err(); rerr != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", rerr)
		return
	}

	cli = &RedisClusterClient{
		cfg:           cfg,
		ClusterClient: rCli,
	}

	return
}
