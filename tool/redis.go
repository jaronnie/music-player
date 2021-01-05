package main

import (

	"fmt"

	"github.com/go-redis/redis/v7"

	"log"
)

var client = redis.NewClient(&redis.Options{
	Addr:"47.95.1.229:6379",
	DB:0,
})

func ConnectRedis() {

	pong, err := client.Ping().Result()

	if err != nil {

		log.Fatal(err)

	}

	if pong != "PONG" {

		log.Fatal("连接redis失败")

	} else {

		log.Println("已连接redis服务")

	}
}

func main() {

	ConnectRedis()

}

func set(key , value string) {
	//redis的string操作
	result, err := client.Set(key, value, 0).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Set %s %s %v", key, value, result )
}

func get(key string)  {
	result, err := client.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}