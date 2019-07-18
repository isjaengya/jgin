package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Conf *viper.Viper = nil

func LoadConfig() *viper.Viper {
	if Conf == nil {
		MODE := os.Getenv("MODE")
		v := viper.New()
		v.AddConfigPath("./api/config/")
		v.SetConfigType("yaml")
		if MODE == "PRODUCTION" {
			v.SetConfigName("production")
		} else {
			v.SetConfigName("development")
		}
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("配置初始化失败，%s", err.Error())
			return nil
		}
		Conf = v
	}
	return Conf
}

func Init() {
	LoadConfig()
}
