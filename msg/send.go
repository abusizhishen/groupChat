package msg

import (
	"fmt"
	"github.com/streadway/amqp"
)

func Send(msg string) (err error) {
	conn := getConn()
	defer conn.Close()
	channel,err := conn.Channel()
	if err != nil{
		return fmt.Errorf("创建信道失败：%s",err)
	}
	defer channel.Close()
	err = channel.ExchangeDeclare(exchangeAll,"fanout",true,false,false,false,nil)
	if err != nil{
		return
	}

	pub := amqp.Publishing{
		Body:[]byte(msg),
	}



	return channel.Publish(exchangeAll,"",false,false,pub)
}
