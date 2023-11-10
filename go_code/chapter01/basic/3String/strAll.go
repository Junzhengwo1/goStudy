package _String

import (
	"fmt"
	"strconv"
	"strings"
)

func MyString() {

	var str string = "abc美king"
	fmt.Println(len(str))
	println("=================================")

	// 先转成切片
	a := []rune(str)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c \n", a[i])
	}
	println("=================================")

	// 将string 转成数字
	b, _ := strconv.Atoi("12")
	i := 1
	b += i
	fmt.Println(b)
	println("=================================")

	sum := strconv.Itoa(i)
	fmt.Println("-----" + sum)
	// 切片 左闭右开
	s := str[0:8]
	println(s) //abc美ki
	// 字符串拼接 +
	king := "kou"

	println(str + king)
	println("=================================")

	//  ====================strings包 一系列的字符串转换过程
	k := "chen"
	// s := "yu"
	replace := strings.Replace(str, "美", "年后", 2)
	println("replace", replace)
	println("=================================")

	//  ====================strings包 字符串长度
	println(len(k))

	//====================strings包 字符串遍历
	for i2 := range k {
		println(i2)
	}
	println("=================================")

	for _, v := range k {
		//fmt.Println(index)
		fmt.Printf("%c \n", v)
	}
	println("=================================")

	//====================strings包 统计字符串中 这个字符的个数

	kins := "chenkingabcdking"
	println(strings.Count(kins, "king"))

	println("=================================")

	//====================strings包 不区分大小比较
	cbc := "nihao"

	bcd := "NIHao"
	// 不区分
	println(strings.EqualFold(cbc, bcd))
	println("=================================")

	println(strings.Index(bcd, "NI"))

}
