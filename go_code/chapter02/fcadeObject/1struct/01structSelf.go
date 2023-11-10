package main

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var a Student //直接定义
	println(a)
	// 结构体名 定义结构体变量 赋值 类似java的构造器
	var s = Student{23, "kou", 22}
	var s1 = Student{}
	s.Name = "king" // 修改结构体属性
	println(s.Age)  // 访问结构体属性
	fmt.Println(s)
	fmt.Println(s1)
	// 通过指针操作
	var p *Student = new(Student)
	(*p).Name = "nihao"
	p.Age = 20 // 底层简化了赋值操作
	fmt.Printf(" type is %T \n", p)
	fmt.Printf(" value is %v ", *p)
}

// 结构体可以转换 但是必须两个结构体属性一致

func myStudent() {

}
