package main

import (
	"go-battle/routers"

	"github.com/spf13/viper"
)

func main() {
	configInit()
	router := routers.RoutersInit()
	router.Run("localhost:3000")
}

func configInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
}
