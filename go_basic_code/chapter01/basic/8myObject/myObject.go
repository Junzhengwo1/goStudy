package _myObject

// 首字母大写 即可被其他包引用 ________学习面向对象 提供给包chapter02使用

type my struct {
	name string
	age  int
}

func (m *my) GetName() string {
	return m.name
}

func (m *my) SetName(name string) {
	m.name = name
}

func (m *my) GetAge() int {
	return m.age
}

func (m *my) SetAge(age int) {
	m.age = age
}

// 工厂模式 给其他包访问 实例 结构体首字母小写时 (类似与 java的构造函数)

func NewMy(name string, age int) *my {
	return &my{name: name, age: age}
}
