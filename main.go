package main

import (
   route2 "tebu_go/api/route"
   Service "tebu_go/api/service"
)

func main() {
	Service.RedisInit()
	Service.MysqlInit()

	route := route2.InitRoute()
	route.Run(":8000")
}