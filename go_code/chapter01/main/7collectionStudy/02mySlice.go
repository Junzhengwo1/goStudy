package main

import "fmt"

func main() {

	//MySlice()
	//MySlice2()
	MySlice3()
}

// 切片 左闭右开 是对 内存连续片段的引用
// 切片中的每个元素的地址都是一样的

func MySlice() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var slice = arr[1:3]
	fmt.Println(slice)

	println(cap(arr))
	// 切片的容量
	fmt.Println(cap(slice))

	slice[1] = 100
	for i, i2 := range slice {
		println(i2, i)
	}

	fmt.Println(arr)
	fmt.Println(slice)

}

// 	切片的定义

func MySlice2() {

	runeArr := []string{"a", "b"}
	runeArr2 := [5]rune{1, 2, 3, 4}
	fmt.Println(runeArr)
	fmt.Println(runeArr2)
	// 切片的定义 类型 长度 容量
	var slice2 []int = make([]int, 2, 10)
	fmt.Println(slice2)
	println(cap(slice2))

	var slice []int
	fmt.Println(slice)
}

func MySlice3() {
	runeArr := []string{"a", "b", "c", "d", "f", "g"}
	runeArr2 := [5]rune{1, 2, 3, 4}
	fmt.Println(runeArr)
	fmt.Println(runeArr2)
	// 切片的定义 类型 长度 容量
	slice := runeArr[1:4]
	slice = append(slice, "king") //本质是新数组 又赋值给slice了
	fmt.Println(slice)
	slice = slice[:0]
	fmt.Println(slice)
}
