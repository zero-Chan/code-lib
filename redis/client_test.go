package redis

import (
	"testing"

	"code-lib/conf"
)

func Test_NewClientConnect(t *testing.T) {
	cfg := &RedisClientConf{
		Addr: conf.AddrConf{
			Host: "localhost",
			Port: 10379,
		},
	}

	cli, err := NewRedisClient(cfg)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	defer cli.Close()

	t.Logf("Redis Connect [%s] ok.", cfg.Addr.String())
}
