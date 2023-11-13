package main

// 这个操作redis 更牛逼
import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {

	// 连接Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	defer client.Close()
	fmt.Println(client)
	// 设置key-value对
	client.Set("key", "KING", 0)
	get := client.Get("key")
	fmt.Println(get.Val())

	client.MSet("name", "KING", "age", 18)
	mGet := client.MGet("name", "age")
	for i := range mGet.Val() {
		fmt.Println(mGet.Val()[i])
	}
	client.HSet("user", "name", "KING")
	client.HSet("user", "age", 18)
	val := client.HGetAll("user").Val()
	println("------------------------------")
	println(val["name"])
	for s, s2 := range val {
		fmt.Println(s, s2)
	}

	for s := range val {
		fmt.Println(val[s])
	}

}
