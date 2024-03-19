package client

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

type Config struct {
	Addr    string
	Limit   int64
	Timeout time.Duration
}

var Cfg Config

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join("."))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}

	Cfg = Config{
		Addr:    viper.GetString("redis.addr"),
		Limit:   viper.GetInt64("floodcontrol.limit"),
		Timeout: time.Duration(viper.GetInt64("floodcontrol.timeout")) * time.Second,
	}
}

func init() {
	initConfig()
}
