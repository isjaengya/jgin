package main

import (
	//"github.com/DeanThompson/ginpprof"
	//_ "net/http/pprof"
	Config "jgin/api/config"
	route2 "jgin/api/route"
	Service "jgin/api/service"
	"runtime"
)

func main() {
	Config.Init()
	Service.RedisInit()
	Service.MysqlInit()

	route := route2.InitRoute()

	//ginpprof.Wrap(route)

	runtime.GOMAXPROCS(1)
	_ = route.Run(":8000")
}
