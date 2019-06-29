package msg

import "github.com/streadway/amqp"

func Receive(msgChan chan amqp.Delivery) (err error) {
	conn := getConn()
	defer conn.Close()
	channel ,err := conn.Channel()
	if err != nil{
		return
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(exchangeAll,"fanout",true,false,false,false,nil)
	if err != nil{
		return
	}
	queue,err := channel.QueueDeclare("",false,true,false,false,nil)
	if err != nil{
		return
	}

	err = channel.QueueBind(queue.Name,"",exchangeAll,false,nil)
	if err != nil{
		return
	}

	msgs,err := channel.Consume(queue.Name,"",false,false,false,false,nil)
	if err != nil{
		return
	}

	for {
		select {
		case msg := <- msgs:
			msgChan <- msg
		}
	}
}
