package main

import (
	"fmt"
	"time"
)

func main() {

	go test() // 开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("main ", i)
		time.Sleep(time.Second)

	}

}

func test() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
