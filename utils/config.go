package utils

import (
	"fmt"
	"strings"

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
