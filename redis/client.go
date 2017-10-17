package redis

import (
	"fmt"

	"gopkg.in/redis.v5"
)

type RedisClient struct {
	cfg *RedisClientConf

	*redis.Client
}

func NewRedisClient(cfg *RedisClientConf) (cli *RedisClient, err error) {
	ErrorPrefix := "[InitError] `Func: NewRedisClient` "

	if cfg == nil {
		err = fmt.Errorf(ErrorPrefix + "`Reason: cfg is nil.`")
		return
	}

	rCli := redis.NewClient(&redis.Options{
		Addr: cfg.Addr.String(),
	})

	statusCmd := rCli.Ping()
	if rerr := statusCmd.Err(); rerr != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", rerr)
		return
	}

	cli = &RedisClient{
		cfg:    cfg,
		Client: rCli,
	}

	return
}
