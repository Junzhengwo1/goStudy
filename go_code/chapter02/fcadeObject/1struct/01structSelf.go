package main

import "fmt"

// 结构体 定义在方法外面可以全局使用

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	// 结构体名 定义结构体变量 赋值 类似java的构造器
	var s = Student{23, "kou", 22}
	s.Name = "king" // 修改结构体属性
	println(s.Age)  // 访问结构体属性
	fmt.Println(s)
}
