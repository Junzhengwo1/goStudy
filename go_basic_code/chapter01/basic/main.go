package main

import (
	_base "goStudy/go_basic_code/chapter01/basic/1base"
	_String "goStudy/go_basic_code/chapter01/basic/3String"
	_interFunc "goStudy/go_basic_code/chapter01/basic/5interFunc"
)

/*
*
init() 函数会先执行被引入包的
*/
func init() {
	println("util……")
}

func main() {
	//_base.MyPointer()
	//println("=============================")
	//_base.Calculate()
	//println("=============================")
	//_base.Switch(1)
	//println("=============================")
	//_base.MyFor()
	//println("=============================")
	k := "king"
	_base.ForRange(k)
	println("=============================")
	//_myFunc.MyFunc()
	_String.MyString()
	println("=============================")
	_interFunc.InterFunc()
}
