package service

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	ginConf "jgin/api/config"
	"log"
)

// 任务队列 类似python的celeryj

var machineryServer *machinery.Server

func MachineryInit() {
	conf := ginConf.Conf
	broker := conf.GetString("machinery.broker")
	defaultQueue := conf.GetString("machinery.default_queue")
	resultBackend := conf.GetString("machinery.result_backend")
	bindingKey := conf.GetString("amqp.binding_key")
	exchange := conf.GetString("amqp.exchange")
	exchangeType := conf.GetString("amqp.exchange_type")
	prefetchCount := conf.GetInt("amqp.prefetch_count")

	var cnf = &config.Config{
		Broker:        broker,
		DefaultQueue:  defaultQueue,
		ResultBackend: resultBackend,
		AMQP: &config.AMQPConfig{
			Exchange:      exchange,
			ExchangeType:  exchangeType,
			BindingKey:    bindingKey,
			PrefetchCount: prefetchCount,
		},
	}

	var err error
	machineryServer, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal("machinery 初始化失败，", err.Error())
	}
}

func GetMachinery() (s *machinery.Server) {
	return machineryServer
}
