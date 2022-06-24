package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func TestReadConfig(t *testing.T) {
	if err := configInit(); err != nil {
		t.Fatal("config init error ", err)
	}

	name := viper.GetString("testName")
	fmt.Println("name before  == ", name)
	// if name != "ha" {
	// 	t.Fatal("value is error = ", name)
	// }
	time.Sleep(10 * time.Second) //没用??
	fmt.Println("name after  == ", name)
}

func configInit() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
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
