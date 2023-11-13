package main

// 轻量级的 redis 操作包
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//redis 操作已经成功
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	do, err := conn.Do("SET", "name", "koujiajun")
	if err != nil {
		return
	}
	fmt.Println(do)
	reply, _ := conn.Do("GET", "name")
	replyStr, _ := redis.String(reply, nil) // 转成 String 类型
	fmt.Println(replyStr)

}
