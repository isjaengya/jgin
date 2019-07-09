package main

import (
	//"github.com/DeanThompson/ginpprof"
	//_ "net/http/pprof"
	Config "tebu_go/api/config"
	route2 "tebu_go/api/route"
	Service "tebu_go/api/service"
)

func main() {
	Config.Init()
	Service.RedisInit()
	Service.MysqlInit()

	route := route2.InitRoute()

	//ginpprof.Wrap(route)

	_ = route.Run(":8000")
}
