package main

import "fmt"

func main() {
	//MyMap()
	MyMap2()
}

func MyMap() {
	var m = make(map[string]int, 10)
	m["king"] = 23
	m["que"] = 24
	m["aa"] = 50
	fmt.Println(m["king"])
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	// 删除map中的元素
	delete(m, "aa")
	fmt.Println(m)

}

// map 的创建

func MyMap2() {
	var m = map[string]int{
		"king": 23,
		"que":  24,
		"aa":   50,
	}
	fmt.Println(m["king"])
	fmt.Println(m)

	println("===============")

	value, ok := m["aa"]
	fmt.Println(ok)
	fmt.Println(value)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}
