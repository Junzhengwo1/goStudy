package _struct

// 接口的定义
// 接口里面不能有任何变量
// 接口的定义是一种类型，它定义了一种行为集合，接口的定义包含了接口的名称和方法的集合
// 接口的定义包含了接口的名称和方法的集合
// 一个结构体可以实现多个接口
// 接口可以继承另一个接口

type SayHello interface {
	sayHello()
}

type A struct {
}

// 实现接口的方法 就是绑定到结构体就行了

func (a A) sayHello() {
	println("你好")
}

func (a A) Ni() {
	println("你好2")
}

type B struct {
}

func (b B) sayHello() {
	println("hello")
}

func Greet(s SayHello) {
	s.sayHello()
}
