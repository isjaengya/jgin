package config

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

var conf *viper.Viper = nil

func LoadConfig() *viper.Viper {
	if conf == nil {
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
			fmt.Println(err)
			return nil
		}
		conf = v
	}
	return conf
}
