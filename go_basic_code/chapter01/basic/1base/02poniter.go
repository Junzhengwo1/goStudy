package _base

import "fmt"

// MyPointer /*
func MyPointer() {
	//指针
	var inr int = 520
	var ptr *int = &inr
	fmt.Println("int的地址：", &inr)
	fmt.Println("int的地址：", ptr)
	//取出指针的值
	fmt.Printf("ptr指向的值=%v \n", *ptr)
	inr = 100
	fmt.Printf("ptr指向的值=%v \n", *ptr)
	var p *int
	p = new(int)
	*p = 10
	fmt.Println("p=", p)
	fmt.Println("p=", *p)

}
