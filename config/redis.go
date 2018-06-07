package config

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

var Cache redis.Conn

func init() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("Nice redis is running")
	// Assign the connection to the package level `cache` variable
	Cache = conn
}