package main

import (
	jginroute "jgin/api/route"
	jginService "jgin/api/service"
	"runtime"
)

func runApp() {
	jginService.InitService()

	route := jginroute.InitRoute()

	//ginpprof.Wrap(route)  //火焰图

	runtime.GOMAXPROCS(1) // goland 调试协程必须指定cpu使用核数为1

	_ = route.Run(":8000")
}

func main() {
	runApp()
}
