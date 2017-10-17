package conf

import (
	"strconv"
)

type AddrConf struct {
	Host string
	Port int64
}

func (this AddrConf) String() string {
	if this.Port <= 0 {
		this.Port = 80
	}

	if this.Host == "" {
		this.Host = "localhost"
	}

	return this.Host + ":" + strconv.FormatInt(this.Port, 10)
}
