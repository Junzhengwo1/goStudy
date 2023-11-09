package main

import (
	"fmt"
	"time"
)

/*
*
日期相关的
*/
func main() {

	now := time.Now()
	fmt.Printf("%v ~~~~  %T \n", now, now)

	println(now.Year())

	println(now.Month())

	// 日期格式化
	println("================================")
	formatted := now.Format(`2006-01-02 15:04:05`) // 必须是这样写
	println(formatted)

}
