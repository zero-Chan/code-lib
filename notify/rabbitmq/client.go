package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitClient struct {
	cfg      *RabbitClientConf
	Vhost    string
	conn     *amqp.Connection
	AmqpChan *amqp.Channel
}

func NewRabbitClient(cfg *RabbitClientConf) (cli *RabbitClient, err error) {
	ErrorPrefix := "[InitError] `Func: NewRabbitClient` "

	if cfg == nil {
		err = fmt.Errorf(ErrorPrefix + "`Reason: cfg is nil.`")
		return
	}

	cli = &RabbitClient{
		cfg: cfg,
	}

	// 当vhost为空默认连接 /
	if cfg.VHost == "" {
		cli.Vhost = "/"
	}

	url := cfg.String()
	if url == "" {
		err = fmt.Errorf(ErrorPrefix+"`Reason: make amqpUri fail.` `params: cfg=%+v\n`", *cfg)
		return
	}

	// 获取连接
	conn, exist := rmqConnMarker.Get(url)
	if !exist {
		conn, err = amqp.Dial(url)
		if err != nil {
			err = fmt.Errorf(ErrorPrefix+"`Reason: amqp dial url[%s] fail: %s`", url, err)
			return
		}
		rmqConnMarker.Set(url, conn)
	} else {
		// 增加对同一个连接的连接数记录
		err = rmqConnMarker.Incr(url, conn)
		if err != nil {
			return
		}
	}

	// 多个rabbitClient可以复用同一个connection,但是一个rabbitClient只使用一个amqpChan
	amqpChan, err := conn.Channel()
	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", err)
	}

	cli.conn = conn
	cli.AmqpChan = amqpChan

	return
}

func (this *RabbitClient) Close() (err error) {
	err = this.AmqpChan.Close()
	if err != nil {
		return
	}

	url := this.cfg.String()

	// 标示减１
	err = rmqConnMarker.Decr(url, this.conn)
	if err != nil {
		return
	}

	// 如果是最后的连接，则断开
	if rmqConnMarker.GetConnAmount(url) == 0 {
		err = rmqConnMarker.Close(url, this.conn)
		if err == nil {
			return
		}
	}

	return
}
