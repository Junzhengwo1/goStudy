package main

import "fmt"

func main() {
	myArray()
}

// go 语言中 数组的长度也属于类型的一部分
func myArray() {

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	var c = [5]int{1, 2, 3, 4, 5}
	var d = [...]int{1, 2, 3, 4, 5, 10}
	// 1: 3: 5: 代表对应下标赋值为 45,22,100 | 给对应下标赋值
	var e = [...]int{1: 45, 3: 22, 5: 100, 4, 5}

	// 二维数组
	f := [][]int{{123, 123}, {5, 6, 7}}

	println(a[3])
	println(sumArray(a))
	println(b[3])
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

func sumArray(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
