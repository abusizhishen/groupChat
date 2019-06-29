package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"groupChat/msg"
	"os"
	"time"
)

func main() {
	var mqUrl string
	flag.StringVar(&mqUrl,"url","","rabbitmq连接信息 `amqp://username:password@host:post/vhost` 如 -url=amqp://guest:guest@localhost:5672/yourHost")
	flag.Parse()

	if mqUrl == ""{
		fmt.Println("无效的mq连接信息")
		os.Exit(0)
	}

	msg.Url = mqUrl
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	var name string
	for {

		fmt.Println("请输入昵称：")
		fmt.Scanln(&name)
		if len(name) != 0 {
			break
		}
	}
	fmt.Println("欢迎你：",name)
	var ch = make(chan amqp.Delivery, 10)
	go msg.Receive(ch)
	go func() {
		for{
			select {
			case ms := <-ch:
				var d msg.Msg
				json.Unmarshal(ms.Body, &d)
				fmt.Println(d.Time.Format("2006-01-02 15:04:05"))
				fmt.Println(fmt.Sprintf("%s：%s", d.Name, d.Msg))
			ms.Ack(false)
			}
		}
	}()

	var content string
	for{
		fmt.Println("请输入：")
		fmt.Scanln(&content)

		if len(content) == 0 {
			continue
		}
		if "" == content {
			continue
		}

		var ms = msg.Msg{
			Name:name,
			Msg:content,
			Time:time.Now(),
		}
		var by,_ = json.Marshal(ms)

		err :=msg.Send(string(by))
		if err != nil{
			fmt.Println("发送失败：",err)
		}
	}
}