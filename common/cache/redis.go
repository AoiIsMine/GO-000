package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

//可以用一个struct包含context和client
var redisClient *redis.Client
var redisCtx context.Context

func Init() (err error) {
	//上下文init
	redisCtx = context.Background()
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	db := viper.GetInt("redis.db")
	password := viper.GetString("redis.password")

	redisAddr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("redis address = %s \n", redisAddr)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err = redisClient.Ping(redisCtx).Result()
	return
	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	//     panic(err)
	// }
}

func GetInstance() (*redis.Client, context.Context) {
	return redisClient, redisCtx
}
