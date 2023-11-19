package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

type RedisConfig struct {
	Host        string
	Port        int
	Password    string
	Database    string
	PoolSize    int
	DialTimeout int
}

// redis 连接池
// 连接Redis
func InitRedis() {
	//Client = redis.NewClient(&redis.Options{
	//	Addr:        "localhost:6379",
	//	Password:    "",                // no password set
	//	DB:          2,                 // use default DB
	//	DialTimeout: 240 * time.Second, // 连接超时时间
	//	PoolSize:    10,                // PoolSize:    10,
	//})
	fmt.Println("Client:", RedisClient)
}
