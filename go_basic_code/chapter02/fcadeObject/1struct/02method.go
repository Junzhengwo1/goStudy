package _struct

import (
	"fmt"
	_myObject "goStudy/go_basic_code/chapter01/basic/8myObject"
)

// 定义结构体 通常直接 绑定String 方法

type Pes struct {
	id   int
	name string
	age  int
}

// 重写了 string

func (p Pes) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

// 继承 通过匿名字段实现继承

type P struct {
	Pes  // 匿名字段实现继承
	addr string
}

// 方法接收器
// 表达了该方法唯一属于 Pes 要想使用该方法 则需要通过 s.Pes.read()
func (p Pes) read() { // 如果大写 可以被其他包调用
	fmt.Println(p.name + "学生正在读书……")
	fmt.Println(p.age)
	p.age = 10000
	fmt.Println(p.age)
}

func TestMethod() {
	p := Pes{22, "que", 23}
	s := P{Pes{2, "张三", 23}, "aksdia"}

	fmt.Println(p.name)
	p.read() // 对象调用方法 对象自己相当于也是传递了自己 函数属于包
	fmt.Println(s)
	s.Pes.read() // 继承了 Pes 所以可以直接调用 Pes 的方法
	fmt.Println(s)

	println("-------------")
	s.Pes.name = "lishi"
	fmt.Println(s)

	println("------------------------------------")
	fmt.Println(p)

	println("------------------------------------")
	// 跨包访问
	//my := _myObject.My{Name: "张三", Age: 23}
	newMy := _myObject.NewMy("张三", 23)
	fmt.Println(*newMy)
}
