package _myFunc

/**
go 语言中 函数不支持 重载
         支持可变参数
*/

func MyFunc() {
	i, v := sum(1, 2)
	println(i, v)
}

/**
可以返回多个值
*/

func sum(a int, b int) (int, int) {
	return a + b, a * b
}

/**
斐波那契数列
*/

func Fbo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fbo(n-1) + Fbo(n-2)
	}
}

/**
猴子吃桃子问题
*/

func EatPeach(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (EatPeach(n+1) + 1) * 2
	}
}
