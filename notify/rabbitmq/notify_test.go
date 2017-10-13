package rabbitmq

import (
	"testing"
)

func Test_RabbitNotifyPublish(t *testing.T) {
	cfg := &RabbitNotifyConf{
		RabbitClientConf: &RabbitClientConf{
			Host:     "localhost",
			Port:     5672,
			UserName: "guest",
			Password: "guest",
			VHost:    "/",
		},
		Exchange:       "myExchange",
		RoutingKey:     "myRoutingKey",
		PublisherInuse: true,

		Kind: "direct",
	}

	notify, err := NewRabbitNotify(cfg)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	err = notify.Init()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	pdata := []byte(`
		{
			"key": 1
		}
	`)

	err = notify.Push(pdata)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
}

func Test_RabbitNotifyConsume(t *testing.T) {
	cfg := &RabbitNotifyConf{
		RabbitClientConf: &RabbitClientConf{
			Host:     "localhost",
			Port:     5672,
			UserName: "guest",
			Password: "guest",
			VHost:    "/",
		},
		Exchange:      "myExchange",
		RoutingKey:    "myRoutingKey",
		QueueName:     "myQueue",
		ConsumerInuse: true,

		Kind: "direct",
	}

	notify, err := NewRabbitNotify(cfg)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	err = notify.Init()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	err = notify.Receive()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	dataChan := notify.Pop()
	t.Logf("get data form rabbitmq: %s", string(<-dataChan))

}
