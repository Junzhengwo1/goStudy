package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testErr(10, 0)
	if err != nil {
		fmt.Println("self Error Is:", err)
		panic(err) // 出现的异常 则不会往下执行了
	}
	fmt.Println("正常执行下面的逻辑。。。")
	test()
}

func testErr(a int, b int) (err error) {
	if b == 0 {
		return errors.New("除数不能为零……") // 返回自定义的error
	} else {
		result := a / b
		fmt.Println(result)
		return nil
	}
}

// defer + recover 相当于是 java 的try catch finally
func test() {
	defer func() { // 匿名函数的调用
		// 调用recover内置函数
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获……")
			fmt.Println("errIs……", err)
		}
	}()

	num1 := 10
	num2 := 2
	result := num1 / num2
	fmt.Println(result)
}
