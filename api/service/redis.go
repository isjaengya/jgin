package service

/*
 * Redis client for temp storage and cache.
 */
import (
	"github.com/go-redis/redis"
	ginConf "jgin/api/config"
	"log"
	"strconv"
	"strings"
	"time"
)

var defaultRedis *redis.Client

var RedisPool map[string]*redis.Client

func RedisInit() {
	RedisPool = make(map[string]*redis.Client)
	config := ginConf.Conf
	Map := config.GetStringMapString("redis.address")
	for key, value := range Map {
		confs := strings.Split(value, " ")
		if len(confs) <= 0 {
			continue
		}
		array := strings.Split(confs[0], "/")
		if len(array[0]) <= 0 {
			continue
		}
		host := array[0]
		db := 0
		password := ""
		if len(confs) == 2 {
			password = confs[1]
		}
		if len(array) == 2 {
			db, _ = strconv.Atoi(array[1])
		}
		cli := redis.NewClient(&redis.Options{
			Addr: host,
			DB:   db,
			//IdleTimeout:  -1,
			Password:     password,
			PoolSize:     config.GetInt("redis.poolsize"),
			MinIdleConns: config.GetInt("redis.minIdleConns"),
			DialTimeout:  time.Duration(config.GetInt("redis.dialtimeout")),
			ReadTimeout:  time.Duration(config.GetInt("redis.readtimeout")),
			WriteTimeout: time.Duration(config.GetInt("redis.writetimeout"))})

		_, err := cli.Ping().Result()
		if err != nil {
			log.Fatalf("%s redis_init_error:%s", key, err.Error())
		}
		RedisPool[key] = cli
	}
	defaultRedis, _ = RedisPool["default"]
}

func GetRedisClient() *redis.Client {
	return defaultRedis
}
