package main

import "fmt"

// 切片 左闭右开 是对 内存连续片段的引用
// 切片中的每个元素的地址都是一样的

func MySlice() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var slice = arr[2:4]
	fmt.Println(slice)

	println(cap(arr))
	// 切片的容量
	fmt.Println(cap(slice))

	println("===========================")
	slice[1] = 100
	for i, i2 := range slice {
		println(i, i2)
	}
	println("===========================")

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
	fmt.Printf(" type is : %T\n", slice)
	fmt.Printf(" type is : %T\n", runeArr2)
	fmt.Printf(" type is : %T\n", runeArr)

}

/**
切片的扩容
*/

func MySlice3() {
	runeArr := []string{"a", "b", "c", "d", "f", "g"}
	//runeArr2 := [5]rune{1, 2, 3, 4}
	fmt.Println(runeArr)
	//fmt.Println(runeArr2)
	// 切片的定义 类型 长度 容量
	slice := runeArr[1:4]
	slice = append(slice, "king", "kou", "king") //本质是新数组 又赋值给slice了
	fmt.Println(slice)
	fmt.Println(runeArr)
	slice = slice[0:] //
	//slice = slice[:0] //
	fmt.Println(slice)
	fmt.Println(slice)
}

/**
切片的复制
*/

func MySlice4() {
	runeArr := []string{"a", "b", "c", "d", "f", "g"}
	slice := runeArr[1:4]
	slice2 := runeArr[3:4]
	fmt.Println(slice)
	fmt.Println(slice2)
	copy(slice2, slice)
	fmt.Println(slice)
	fmt.Println(slice2)

}

/**
make 和 new 的区别
	make 只能用于map，chan，slice
	new 可以任意类型 默认是nil值
*/

func main() {

	//MySlice()
	//MySlice2()
	//MySlice3()
	MySlice4()
}
