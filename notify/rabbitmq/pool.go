package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

// amqp.Connection连接池
var (
	// 记录已经建立的连接
	rmqConnMarker RmqConnMarker = RmqConnMarker{
		// map[amqpUrl]*amqp.Connection
		conns: make(map[string]*connRecord),
	}
)

type RmqConnMarker struct {
	conns map[string]*connRecord
	mutex sync.Mutex
}

type connRecord struct {
	// 连接个数
	amount int64

	// amqp 连接
	conn *amqp.Connection
}

func (this *connRecord) Set(conn *amqp.Connection) {
	this.conn = conn
	this.amount = 1
}

func (this *connRecord) Get() (conn *amqp.Connection) {
	if this.amount == 0 {
		return nil
	}

	return this.conn
}

func (this *connRecord) Incr() {
	this.amount++
}

func (this *connRecord) Decr() {
	if this.amount > 0 {
		this.amount--
	}
}

func (this *connRecord) Amount() int64 {
	return this.amount
}

func (this *connRecord) Close() (err error) {
	this.amount = 0
	if this.conn != nil {
		err = this.conn.Close()
		return
	}

	return
}

// Set：　记录一个新连接
func (this *RmqConnMarker) Set(key string, val *amqp.Connection) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.conns[key] = &connRecord{}
	this.conns[key].Set(val)
}

// Get：　获取一个连接
func (this *RmqConnMarker) Get(key string) (val *amqp.Connection, exist bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	connr, exist := this.conns[key]
	if !exist {
		return
	}

	if connr.amount <= 0 || connr.conn == nil {
		exist = false
		return
	}

	return connr.Get(), true
}

// GetSet：　如果连接不存在才写入
func (this *RmqConnMarker) GetSet(key string, val *amqp.Connection) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if connr, exist := this.conns[key]; !exist {
		this.conns[key].Set(val)
	} else {
		if connr.amount <= 0 || connr.conn == nil {
			this.conns[key].Set(val)
		}
	}
}

// Incr: 增加一个连接数
func (this *RmqConnMarker) Incr(key string, val *amqp.Connection) (err error) {
	ErrorPrefix := "[InitError] `Func: rmqConnMarker.Incr` "

	this.mutex.Lock()
	defer this.mutex.Unlock()

	if connr, exist := this.conns[key]; !exist {
		this.conns[key].Set(val)
	} else {
		if connr.conn != val {
			err = fmt.Errorf(ErrorPrefix + "`Reason: Connection not the same.`")
			return
		}

		// 如果是同一个连接对象，才自增
		connr.Incr()
	}

	return
}

func (this *RmqConnMarker) Decr(key string, val *amqp.Connection) (err error) {
	ErrorPrefix := "[InitError] `Func: rmqConnMarker.Decr` "

	this.mutex.Lock()
	defer this.mutex.Unlock()

	if connr, exist := this.conns[key]; !exist {
		return
	} else {
		if connr.conn != val {
			err = fmt.Errorf(ErrorPrefix + "`Reason: Connection not the same.`")
			return
		}

		// 如果是同一个连接对象，才自减
		connr.Decr()
	}

	return
}

func (this *RmqConnMarker) GetConnAmount(key string) int64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	connr, exist := this.conns[key]
	if !exist {
		return 0
	}
	return connr.Amount()
}

func (this *RmqConnMarker) Close(key string, val *amqp.Connection) (err error) {
	ErrorPrefix := "[InitError] `Func: rmqConnMarker.Close` "

	this.mutex.Lock()
	defer this.mutex.Unlock()

	if connr, exist := this.conns[key]; !exist {
		return
	} else {
		if connr.conn != val {
			err = fmt.Errorf(ErrorPrefix + "`Reason: Connection not the same.`")
			return
		}

		err = connr.Close()
		if err != nil {
			return
		}

		delete(this.conns, key)
	}

	return
}
