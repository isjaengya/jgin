package main

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	ginConf "jgin/api/config"
	jginService "jgin/api/service"
	"jgin/tasks"
	"log"
)

// 启动Machinery的Master
func main() {
	conf := ginConf.LoadConfig()
	broker := conf.GetString("machinery.broker")
	default_queue := conf.GetString("machinery.default_queue")
	result_backend := conf.GetString("machinery.result_backend")
	results_expire_in := conf.GetInt("machinery.results_expire_in")

	var cnf = &config.Config{
		Broker:          broker,
		DefaultQueue:    default_queue,
		ResultBackend:   result_backend,
		ResultsExpireIn: results_expire_in,
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Fatal("machinery 启动失败，%s", err.Error())
	} else {
		// 注册任务
		err = server.RegisterTasks(tasks.GetTasks())

		// 注册service
		jginService.InitService()
		worker := server.NewWorker("worker_name", 10)
		err = worker.Launch()
	}
}
