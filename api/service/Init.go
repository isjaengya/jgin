package service

import (
	"fmt"
	ginConf "jgin/api/config"
)

func InitService() {
	ginConf.Init()
	fmt.Println("配置初始化成功")
	RedisInit()
	fmt.Println("redis 初始化成功")
	MysqlInit()
	fmt.Println("mysql 初始化成功")
	MongoInit()
	fmt.Println("mongodb 初始化成功")
	MachineryInit()
	fmt.Println("machinery 初始化成功")

}
