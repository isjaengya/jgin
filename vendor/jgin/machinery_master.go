package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	Config "jgin/api/config"
	"jgin/tasks"
)

// 启动Machinery的Master
func main() {
	var cnf config.Config
	conf := Config.Conf
	broker := conf.GetString("machinery.broker")
	default_queue := conf.GetString("machinery.default_queue")
	result_backend := conf.GetString("machinery.result_backend")
	binding_key := conf.GetString("amqp.binding_key")
	exchange := conf.GetString("amqp.exchange")
	exchange_type := conf.GetString("amqp.exchange_type")
	prefetch_count := conf.GetInt("amqp.prefetch_count")

	cnf = config.Config{
  	Broker:             broker,
  	DefaultQueue:       default_queue,
  	ResultBackend:      result_backend,
  	AMQP:               &config.AMQPConfig{
  	 Exchange:     exchange,
  	 ExchangeType: exchange_type,
  	 BindingKey:   binding_key,
  	 PrefetchCount: prefetch_count,
  		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
	  // do something with the error
	  fmt.Println(err.Error())
	}

	err = server.RegisterTask("add", tasks.HelloWorld)



}