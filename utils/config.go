package utils

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func ViperInit() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func RedisInit() *redis.Client {
	rDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rDB
}
