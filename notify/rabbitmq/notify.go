package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// RabbitNotify
// a RabbitNotify must be only use for only one queue push or pop.
type RabbitNotify struct {
	*RabbitClient

	cfg            *RabbitNotifyConf
	Exchange       string
	RoutingKey     string
	QueueName      string
	PublisherInuse bool
	ConsumerInuse  bool

	stop    chan bool
	data    chan []byte
	deliver <-chan amqp.Delivery

	curDeliver amqp.Delivery
	isAck      chan bool

	Advance *RabbitNotifyAdvance
}

func NewRabbitNotify(cfg *RabbitNotifyConf) (notify *RabbitNotify, err error) {
	ErrorPrefix := "[InitError] `Func: NewRabbitNotify` "

	if cfg == nil {
		err = fmt.Errorf(ErrorPrefix + "`Reason: cfg is nil.`")
		return
	}

	cli, err := NewRabbitClient(cfg.RabbitClientConf)
	if err != nil {
		return
	}

	notify = &RabbitNotify{
		RabbitClient: cli,

		cfg:            cfg,
		Exchange:       cfg.Exchange,
		RoutingKey:     cfg.RoutingKey,
		QueueName:      cfg.QueueName,
		PublisherInuse: cfg.PublisherInuse,
		ConsumerInuse:  cfg.ConsumerInuse,

		stop:  make(chan bool),
		data:  make(chan []byte),
		isAck: make(chan bool),

		Advance: NewRabbitNotifyAdvanceParams(),
	}

	if notify.PublisherInuse {
		// TODO Advance / Banned user declare exchange
		err = notify.AmqpChan.ExchangeDeclare(notify.Exchange, cfg.Kind, false, false, false, false, nil)
		if err != nil {
			err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", err)
			return
		}
	}

	if notify.ConsumerInuse {
		// TODO Advance / Banned user declare exchange
		rmqQueue, nerr := notify.AmqpChan.QueueDeclare(notify.QueueName, false, false, false, false, nil)
		if nerr != nil {
			err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", nerr)
			return
		}

		// TODO Advance / Banned user declare exchange
		nerr = notify.AmqpChan.QueueBind(rmqQueue.Name, notify.RoutingKey, notify.Exchange, false, nil)
		if nerr != nil {
			err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", nerr)
			return
		}
	}

	return
}

func (this *RabbitNotify) Name() string {
	return this.QueueName
}

func (this *RabbitNotify) StopPop() {
	this.stop <- true
}

func (this *RabbitNotify) Pop() <-chan []byte {
	return this.data
}

func (this *RabbitNotify) Ack() (err error) {
	ErrorPrefix := "[AckFail] `Func: RabbitNotify.Ack` "

	err = this.curDeliver.Ack(this.Advance.Multiple)
	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", err)
		return
	}

	this.isAck <- true

	return
}

func (this *RabbitNotify) Receive() (err error) {
	ErrorPrefix := "[ReceiveFail] `Func: RabbitNotify.Receive` "

	// TODO
	// Advance 参数
	this.deliver, err = this.AmqpChan.Consume(this.QueueName, "", false, false, false, false, nil)
	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", err)
		return
	}

	go this.pop()

	return
}

func (this *RabbitNotify) pop() {
	var ok bool
	for {
		select {
		case <-this.stop:
			return
		case this.curDeliver, ok = <-this.deliver:
			if !ok {
				continue
			}

			this.data <- this.curDeliver.Body
			<-this.isAck
		}
	}
}

func (this *RabbitNotify) Push(data []byte) (err error) {
	ErrorPrefix := "[PushFail] `Func: RabbitNotify.Push` "

	if this.PublisherInuse != true {
		err = fmt.Errorf(ErrorPrefix + "`Reason: Publisher not Open.`")
		return
	}

	err = this.AmqpChan.Publish(this.Exchange, this.RoutingKey, this.Advance.Mandatory, this.Advance.Immediate, amqp.Publishing{
		Body: data,
	})

	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", err)
		return
	}

	return
}

func (this *RabbitNotify) PushNative(mandatory, immediate bool, msg amqp.Publishing) (err error) {
	ErrorPrefix := "[PushFail] `Func: RabbitNotify.PushNative` "

	if this.PublisherInuse != true {
		err = fmt.Errorf(ErrorPrefix + "`Reason: Publisher not Open.`")
		return
	}

	err = this.AmqpChan.Publish(this.Exchange, this.RoutingKey, mandatory, immediate, msg)
	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"`Reason: %s`", err)
		return
	}

	return
}

type RabbitNotifyAdvance struct {
	// Push 参数
	// 如果 mandatory　为 true，
	// 当 消息推送给一个没有绑定队列的routerKey　时
	// 消息会被丢弃
	// 默认值为false
	Mandatory bool

	// Push参数
	// 如果 immediate 为 true.
	// 当　消息推送给一个没有消费者的队列　时
	// 消息会被丢弃
	// 默认值为false
	Immediate bool

	// Consume 参数

	// Ack 参数
	// 是否对该channel上的所有消息发出回应
	Multiple bool
}

func NewRabbitNotifyAdvanceParams() *RabbitNotifyAdvance {
	adv := RabbitNotifyAdvance{
		Mandatory: false,
		Immediate: false,
		Multiple:  false,
	}

	return &adv
}
