package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// redis 连接池
// 连接Redis
var (
	client = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",                // no password set
		DB:          2,                 // use default DB
		DialTimeout: 240 * time.Second, // 连接超时时间
		PoolSize:    10,                // PoolSize:    10,
	})
)

func main() {

	// 使用连接池
	_, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 设置key-value对
	client.Set("key", "wang", 0)
	get := client.Get("key")
	fmt.Println(get.Val())

}
