package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	// 主死从随
	//go test() // 开启协程
	//for i := 0; i < 10; i++ {
	//	fmt.Println("main ", i)
	//	time.Sleep(time.Second)
	//
	//}

	// 第二种方式 启动协程  匿名函数的方式
	//for i := 0; i < 10; i++ {
	//	go func(n int) {
	//		fmt.Println("go ", n)
	//	}(i)
	//}
	//time.Sleep(time.Second)

	// waitGroup := sync.WaitGroup{} 来管理协程
	testWaitGroup()

	test2()

}

func test() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func testWaitGroup() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(n int) {
			fmt.Println("go ", n)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}

// 多个协程操作同一个变量
func test2() {
	var num int
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(n int) {
			num = n
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	fmt.Println(num) // 理论上应该是 9
}
