package config

import (
	"github.com/go-redis/redis"
	"fmt"
)

var Cache *redis.Client

func init() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0, 
	})


	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	fmt.Println("Nice, redis is running")
	Cache = client
}

