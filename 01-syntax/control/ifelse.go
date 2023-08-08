package main

func IfOnly(age int) string {
	if age >= 18 {
		println("成年了")
	}
	return "他还是个孩子"
}

func IfElse(age int) string {
	if age >= 18 {
		println("成年了")
		return ""
	} else {
		println("他还是个孩子")
		return ""
	}
}

func IfElseIf(age int) string {
	if age >= 18 {
		println("成年了")
		return ""
	} else if age >= 12 {
		println("骚年")
		return ""
	} else {
		println("他还是个孩子")
		return ""
	}
}

func IfElseIfV2(age int) string {
	if age >= 18 {
		println("成年了")
		return ""
	} else if age >= 12 {
		println("骚年")
		return ""
	}
	println("他还是个孩子")
	return ""
}

func IfNewVariable(start int, end int) string {
	/**
	Go 的 if else 支持一种新的写法，可以在 if -else块里面定义一个新的局部变量。
	在右边 distance 的只作用于 if -else 块，离开了这个范围就无法使用了。
	*/
	if distance := end - start; distance > 100 {
		println(distance)
		return "距离太远了"
	} else {
		println(distance)
		return "距离比较近"
	}

	//
	//if xx, err := xxx; err != nil {
	//
	//}

	// 编译错误
	//println(distance)
}
