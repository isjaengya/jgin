package service

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	ginConf "jgin/api/config"
	"log"
)

// 任务队列 类似python的celeryj

var MachineryServer *machinery.Server

func MachineryInit() {
	conf := ginConf.Conf
	broker := conf.GetString("machinery.broker")
	default_queue := conf.GetString("machinery.default_queue")
	result_backend := conf.GetString("machinery.result_backend")
	binding_key := conf.GetString("amqp.binding_key")
	exchange := conf.GetString("amqp.exchange")
	exchange_type := conf.GetString("amqp.exchange_type")
	prefetch_count := conf.GetInt("amqp.prefetch_count")

	var cnf = &config.Config{
		Broker:        broker,
		DefaultQueue:  default_queue,
		ResultBackend: result_backend,
		AMQP: &config.AMQPConfig{
			Exchange:      exchange,
			ExchangeType:  exchange_type,
			BindingKey:    binding_key,
			PrefetchCount: prefetch_count,
		},
	}

	var err error
	MachineryServer, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal("machinery 初始化失败， %s", err.Error())
	}
}

func GetMachinerty() (s *machinery.Server) {
	return MachineryServer
}
