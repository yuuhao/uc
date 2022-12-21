package utils

import (
	"fmt"
	"log"

	redigo "github.com/gomodule/redigo/redis"
)

var RedisPool *redigo.Pool

func InitRedis() *redigo.Pool {

	RedisPool = &redigo.Pool{
		MaxIdle:     3,   //最初的连接数量
		MaxActive:   5,   // 最大连接数
		IdleTimeout: 300, // 超过这个时间，关闭连接
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", "localhost:6379")
		},
	}
	rds := RedisPool.Get()
	defer rds.Close()

	_, err := rds.Do("PING")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("redis ping fail %+v\n", err)
	}

	//RedisPool = pool
	return RedisPool
}
