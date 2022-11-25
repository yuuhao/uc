package utils

import "github.com/go-redis/redis/v8"

var Redis *redis.Client

func InitRedis() *redis.Client {

	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return Redis
}
