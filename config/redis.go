package config

import (
	"github.com/go-redis/redis"
	"fmt"
)

var Cache *redis.Client

func init() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})


	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	fmt.Println("Nice redis is running")
	// Assign the connection to the package level `cache` variable
	Cache = client
}

