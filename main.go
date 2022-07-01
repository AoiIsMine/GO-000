package main

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"go-battle/common/db"
	"go-battle/model"
	"go-battle/router"
	"go-battle/service"
)

func main() {
	var err error
	if err = configInit(); err != nil {
		panic(fmt.Sprintln("config init error ", err))
	}

	if err = db.Init(); err != nil {
		panic(fmt.Sprintln("db init error ", err))
	}

	if err = migrate(db.DBConn()); err != nil {
		panic(fmt.Sprintln("db migrate error ", err))
	}

	serviceInit(db.DBConn())

	router := router.RoutersInit()
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	router.Run(fmt.Sprintf("%s:%s", host, port))
}

func configInit() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			return errors.New("config file not found")
		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("config file was found but another error is %v", err)

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

func serviceInit(dbConn *gorm.DB) {
	service.TestServiceInit(dbConn)
}

//TODO优化,获取包内所有结构体
func migrate(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(&model.Test{})
}
