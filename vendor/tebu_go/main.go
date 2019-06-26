package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
	route2 "tebu_go/api/route"
	Service "tebu_go/api/service"
)

func main() {
	Service.RedisInit()
	Service.MysqlInit()

	type ruote *gin.Engine

	route := route2.InitRoute()

	ginpprof.Wrap(route)

	route.Run(":8000")
}
