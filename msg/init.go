package msg

import (
	"github.com/streadway/amqp"
	"time"
)

var Url string
func getConn() *amqp.Connection{
	conn,err :=amqp.Dial(Url)
	if err != nil{
		panic(err)
	}

	return conn
}

var exchangeAll = "all"
type Msg struct {
	Name string
	Msg string
	Time time.Time
}