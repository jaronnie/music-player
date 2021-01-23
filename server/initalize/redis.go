package initalize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
	Addr:"47.95.1.229:6379",
	DB:0,
	PoolSize: 100,
	MinIdleConns: 20,
	DialTimeout: 5 * time.Second,
	})

	pong, err := client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	if pong != "PONG" {
		log.Fatal("连接redis失败")
	} else {
		log.Println("已连接redis服务")
	}

	return client
}

