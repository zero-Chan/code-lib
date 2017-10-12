package rabbitmq

import (
	"testing"
)

func Test_ClientConnect(t *testing.T) {
	cfg := &RabbitClientConf{
		Host:     "localhost",
		Port:     5672,
		UserName: "guest",
		Password: "guest",
		VHost:    "/",
	}

	cli, err := NewRabbitClient(cfg)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	t.Logf("Rabbitmq Client Connect ok!!!")

	err = cli.Close()
	if err != nil {
		t.Error("%s", err)
		t.FailNow()
	}
	t.Logf("Rabbitmq Client Close ok!!!")
}
