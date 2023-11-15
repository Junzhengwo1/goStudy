package main

import (
	"fmt"
)

// 管道：有它自己类型 本质是 队列 数据结构
// 管道本质就是保证全局变量线程安全
func myChannel() {
	// 管道的声明
	// 管道的声明格式：
	// var 管道名 chan 数据类型
	// 管道的使用格式：
	// 管道名 <- 值
	// <- 管道名
	// 管道的特点：
	// 1. 管道是一种通信机制
	// 2. 管道是一种数据结构
	// 在没有使用协程的情况下，如果数据已经取完了就会报 死锁错误

	v := make(chan int, 10)
	fmt.Printf("type is %T \n", v)

	//var ch chan int
	//if ch == nil {
	//	ch = make(chan int, 10)
	//}
	// 向管道写数据
	v <- 10
	v <- 20
	s := cap(v)
	fmt.Printf("len is %v \n, cap is %v \n", len(v), s)
	// 取一个元素
	num := <-v
	fmt.Println(num)
	fmt.Printf("len is %v \n, cap is %v \n", len(v), s)

	//for {
	//	select {
	//	case n := <-v:
	//		fmt.Printf("get a value %v \n", n)
	//	default:
	//		fmt.Println("no value")
	//		break
	//	}
	//}

}

type stu struct {
	Name string
	Age  int
}

// 管道可以存放任意类型的数据 但是需要配合 空接口使用
func myChannel2() {
	ch := make(chan interface{}, 5)
	a := stu{"xiaoMing", 18}
	b := 10
	c := "hello"
	ch <- a
	ch <- b
	ch <- c

	xiao := <-ch

	fmt.Printf("xiao type is %T \n value is %v\n", xiao, xiao)
	fmt.Println(xiao.(stu).Name) // 空接口 要用.(类型) 才能取值 // 类型断言

}

// 管道的 关闭 遍历
func myChannel3() {
	ch := make(chan int, 10)
	fmt.Printf("type is %T \n", ch)
	// 向管道写数据
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
	ch <- 60
	close(ch) // 关闭管道 关闭管道之后不可以写 但是可以读
	x, ok := <-ch
	if ok {
		fmt.Println(x)
	}
	// 遍历的时候 之后当管道关闭之后才可以正常遍历 如果没有关闭 会报错 死锁
	for i := range ch {
		fmt.Println("v=", i)
	}

}

func writeData(intChan chan int) {
	defer close(intChan) // 写完数据 关闭管道
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("add value %v \n", i)

	}
}

func readData(intChan chan int, exit chan bool) {
	// 读管道
	for {
		n, ok := <-intChan
		if !ok {
			fmt.Println("readData exit")
			break
		}
		fmt.Printf("get a value %v \n", n)
	}
	exit <- true
	close(exit)
}

//var waitGroup sync.WaitGroup

func main() {
	//myChannel()
	//myChannel2()
	//myChannel3()

	// 管道 配合协程 应用案例
	//ints := make(chan int, 50)
	//exit := make(chan bool, 1)
	//go writeData(ints)
	//go readData(ints, exit)
	//// 等待 读结束
	//for {
	//	_, ok := <-exit
	//	if !ok {
	//		break
	//	}
	//}

	testChannelEg()

}

// 管道 配合协程 应用案例 统计素数
func testChannelEg() {
	ch := make(chan int, 10)
	ch2 := make(chan string, 10)
	fmt.Printf("type is %T \n", ch)
	fmt.Printf("type is %T \n", ch2)
	// 向管道写数据
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
	ch <- 60

	ch2 <- "hello"
	ch2 <- "world"

	// 读管道
	for {
		select {
		case n := <-ch:
			fmt.Printf("get a value %v \n", n)
		case s := <-ch2:
			fmt.Printf("get a value %v \n", s)
		default:
			fmt.Println("退出")
			return
		}

	}

}
