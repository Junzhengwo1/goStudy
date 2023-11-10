package main

import "fmt"

type Pes struct {
	id   int
	name string
	age  int
}

// 继承 通过匿名字段实现继承

type P struct {
	Pes
	addr string
}

// 方法接收器
// 表达了该方法唯一属于 Pes 要想使用该方法 则需要通过 s.Pes.read()
func (s Pes) read() {
	fmt.Println(s.name + "学生正在读书……")
}

func main() {
	p := Pes{22, "que", 23}

	s := P{Pes{2, "张三", 23}, "aksdia"}

	fmt.Println(p.name)
	p.read() // 对象调用方法 对象自己相当于也是传递了自己 函数属于包
	fmt.Println(s)
	s.read() // 继承了 Pes 所以可以直接调用 Pes 的方法

}
