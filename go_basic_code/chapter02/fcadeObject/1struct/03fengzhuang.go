package _struct

import (
	"fmt"
	_myObject "goStudy/go_basic_code/chapter01/basic/8myObject"
)

// 封装

func Encapsulation() {

	newMy := _myObject.NewMy("张三", 23)
	{
	} // 返回的是 指针
	newMy.SetAge(100000) // 封装的方法就是 指针操作的
	fmt.Println(*newMy)

}
