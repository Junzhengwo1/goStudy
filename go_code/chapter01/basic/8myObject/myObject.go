package _myObject

// 首字母大写 即可被其他包引用

type my struct {
	name string
	Age  int
}

// 工厂模式 给其他包访问 实例 结构体首字母小写时 (类似与 java的构造函数)

func NewMy(name string, age int) *my {
	return &my{name: name, Age: age}
}
