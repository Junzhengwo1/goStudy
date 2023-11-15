package _interFunc

import "fmt"

func InterFunc() {

	// new 函数通常是 用于创建基本数据类型
	num := new(int) // 返回的是对应类型的指针
	*num = 10
	fmt.Printf("type is %T \n, num is %v \n,num address is %v \n ,num pointer for is %v \n",
		num, num, &num, *num)

}
