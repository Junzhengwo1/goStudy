package _base

func IfElse(a int) {
	if a == 1 {
		println("a is 1")
	} else if a == 2 {
		println("a is 2")
	} else {
		println("a is not 1 or 2")
	}

}

func Switch(a int) {
	switch a {
	case 1:
		println("a is 1")
		fallthrough // 匹配到了 还会走到下一个case 语句 switch 穿透
	case 2:
		println("a is 2")
	default:
		println("a is not 1 or 2")
	}
}
