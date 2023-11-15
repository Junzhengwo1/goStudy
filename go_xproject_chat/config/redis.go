package config

import (
	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
)

// redis 连接池
// 连接Redis
func init() {
	//Client = redis.NewClient(&redis.Options{
	//	Addr:        "localhost:6379",
	//	Password:    "",                // no password set
	//	DB:          2,                 // use default DB
	//	DialTimeout: 240 * time.Second, // 连接超时时间
	//	PoolSize:    10,                // PoolSize:    10,
	//})
	//fmt.Println("Client:",Client)
}
