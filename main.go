package main

import (
	jginroute "jgin/api/route"
	jginService "jgin/api/service"
	"runtime"
)

func runApp() {
	jginService.InitService()

	route := jginroute.InitRoute()

	//ginpprof.Wrap(route)

	runtime.GOMAXPROCS(1)
	_ = route.Run(":8000")
}

func main() {
	runApp()
}
