package main

import (
	"fmt"
	"go-battle/routers"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	if err := configInit(); err != nil {
		panic(fmt.Sprintf("config init error ", err))
	}

	router := routers.RoutersInit()
	router.Run("localhost:3000")
}

func configInit() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			return fmt.Errorf("Config file not found")

		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("Config file was found but another error is %v", err)

		}
	}
	// Config file found and successfully parsed
	fmt.Println("config init success")

	//add config change event
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	return nil
}
