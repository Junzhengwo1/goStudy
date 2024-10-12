package _myFunc

import "fmt"

/**
go 语言中 、
	基础函数：
		1、函数不支持 重载
		2、可以返回多个值
		3、函数本身可以赋值给变量 此时该变量可以直接调用
		4、可以自定义数据类型 相当于起别名 可以理解为 java的对象 其实也是结构体
		5、支持对返回值 命名
	init函数：
		1、每个源文件都可以有一个init函数
		2、会在 main函数执行之前执行 全局变量->包init-> main的init
	匿名函数：
		1、匿名函数在定义时 就可以直接调用
	闭包：
		1、一个函数与其相关的
	defer关键字
		1、会将该关键字的 语句压入一个专门的栈中 去执行后面的语句
		2、defer 语句修饰的时候的变量 也是当时的时候的值 不受后面的 程序所影响
		3、用于 延迟操作相关的 语言 例如释放资源

*/

/*
*
全局匿名函数的使用
*/
var (
	f = func(n1 int, n2 int) int { return n1 + n2 }
)

func init() {
	println("init2……")
}

func MyFunc() {
	i, v := sum(1, 2)
	println(i, v)
	println(f(1, 3))

	a := sum // 函数本身可以赋值给变量
	a(1, 2)  // 等价与 sum(1, 2)
	fmt.Printf(" type is : %T\n", a)

	println("==================")
	// 匿名函数 定义的时候直接就调用
	res := func(n1 int, n2 int) int { return n1 + n2 }(1, 2)
	println(res)

	println("==================")
	f2 := getSum()
	println(f2(2))
	println(f2(3))
	println("闭包==================")
	println(getSum()(2))
	println("==================")
	s, mit := sum2(2, 3)
	println(s, mit)

}

/**
可以返回多个值
*/

func sum(a int, b int) (int, int) {
	return a + b, a * b
}

/*
*
函数返回值 命名
*/
func sum2(a int, b int) (sum int, mit int) {
	defer println("defer……")
	sum = a + b
	mit = a * b
	return
}

/**
斐波那契数列
*/

func Fbo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fbo(n-1) + Fbo(n-2)
	}
}

/**
猴子吃桃子问题
*/

func EatPeach(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (EatPeach(n+1) + 1) * 2
	}
}

/**
闭包
*/

func getSum() func(int) int {
	var sum int = 10
	return func(i int) int {
		sum = sum + i
		return sum
	}
}
