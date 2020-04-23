package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func Connect(host string, port string, pwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pwd,
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
