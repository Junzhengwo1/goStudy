package _base

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
* go语言中数据类型
go语言中 首字母的大小写来区别是否是私有的
1、基本数据类型 四个：

	1)数值  默认值  0 保小保大原则
	2)字符 byte  默认值
	3)bool  默认值 false
	4)string 默认值 ""

2、复杂数据类型：

	1)pointer
	2)数组
	3)结构体
	4)管道
	5)函数
	6)切片
	7)接口
	8)map
*/

func DataType() {

	var name = false
	println(name)

	var i = 1
	fmt.Println(i)

	println("=================================")

	var wang uint = 20
	println(wang)
	println("=================================")

	var a int8 = -128
	fmt.Println(a)
	fmt.Printf("%T %d \n", a, unsafe.Sizeof(a))

	println("=================================")

	// 浮点型
	var price float32 = 12.34
	println(price)
	fmt.Println(price)

	println("=================================")

	// 字符类型 一般字符就用byte来存；特殊的就用int
	var c byte = 'k'
	var c2 int = '北'
	fmt.Printf("c=%c %c \n", c, c2)

	println("=================================")

	// 字符串使用
	var str = "king"
	fmt.Println(len(str))
	println("=================================")

	var k = `var a int8 = -128`
	fmt.Println(k)
	println("=================================")

	// 字符串拼接
	var w = "hello"
	var e = "world"

	var t = "king" + "prince"
	fmt.Println(w + e)
	fmt.Println(t)

	println("=================================")

	// 基本数据类型的相互转换
	g := 100.00
	//把g转成float 类型
	f := int8(g)
	fmt.Println(f)

	println("=================================")
	//将int转成string
	var v = 100
	re := fmt.Sprintf("%d", v)
	fmt.Printf("%T %q \n", re, re)

	println("=================================")

	var v2 = 222
	str = strconv.FormatInt(int64(v2), 10) //十进制
	fmt.Printf("type %T str %q  \n", str, str)
	str = strconv.Itoa(v2) // 与上面的等价
	fmt.Println(str)

	println("=================================")

	//将string转成其他基本数据类型
	var s = "100"
	fmt.Printf("v = %v  %T \n", s, s)
	parseInt, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("type is %T \n", parseInt)

	println("=================================")

	var s2 = "100.2134"
	fmt.Printf("v = %v  %T \n", s2, s2)
	float, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("type is %T \n", float)

	println("=================================")

	// 将string 转成数字
	b, _ := strconv.Atoi("12")
	println(b)
	fmt.Printf("type is %T \n", b)

	println("=================================")

	var userName = "king"
	fmt.Println(userName)
	println("=================================")

	// 键盘输入语句
	// fmt.Scanln()

}
