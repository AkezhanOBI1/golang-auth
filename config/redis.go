package config

import (
	"github.com/go-redis/redis"
	"fmt"
)

var Cache *redis.Client

func init() {
	// Initialize the redis connection to a redis instance running on your local machine
	/*conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}*/

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

