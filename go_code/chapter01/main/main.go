package main

import (
	_base "goStudy/go_code/chapter01/main/1base"
	_String "goStudy/go_code/chapter01/main/3String"
)

/*
*
init() 函数会先执行被引入包的
*/
func init() {
	println("init……")
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
}
