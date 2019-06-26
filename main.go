package main

import (
	//"github.com/DeanThompson/ginpprof"
	//_ "net/http/pprof"
	route2 "tebu_go/api/route"
	Service "tebu_go/api/service"
)

func main() {
	Service.RedisInit()
	Service.MysqlInit()

	route := route2.InitRoute()

	//ginpprof.Wrap(route)

	route.Run(":8000")
}
