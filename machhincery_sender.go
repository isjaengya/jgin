package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	ginConf "jgin/api/config"
	"time"
)

var MachineryServer *machinery.Server

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
		// do something with the error
		fmt.Println(err.Error())
	} else {
		MachineryServer = server
		for i := 0; i < 1; i++ {
			go async_send(i)
			//fmt.Println(i)
		}
	}
	time.Sleep(time.Second * 1)
}

func async_send(i int) {
	fmt.Println(i)
	signature := &tasks.Signature{
		Name: "hello",
		Args: []tasks.Arg{
			{
				Type:  "int32",
				Value: i,
			},
		},
	}

	asyncResult, err := MachineryServer.SendTask(signature)
	if err != nil {
		fmt.Println(err.Error(), "1111111")
	} else {
		fmt.Println(asyncResult.GetState().IsSuccess(), "2222222")
	}
}
