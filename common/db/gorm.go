package db

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbConn *gorm.DB

func Init() (err error) {
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	name := viper.GetString("db.dbName")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	fmt.Println("mysql dns ==", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "resource_",
			SingularTable: true, //避免orm使用复数表名进行sql操作
		},
	})
	if err != nil {
		fmt.Printf("db init error=%v", err)
		return
	}
	dbConn = db

	//连接池
	sqlDB, err := dbConn.DB()
	if err != nil {
		return
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}

func DBConn() *gorm.DB {
	return dbConn
}
