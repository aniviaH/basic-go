package main

import "math"

func main() {
	// int 类型
	IntUse()

	//数字类型的极值
	Extremum()

	// string 类型
	String()

	// byte 类型
	ByteUse()

	// bool 类型
	BoolUse()
}

func IntUse() {
	// 加减乘除：+、-、*、/
	var a int = 100
	var b int = 2
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
	//println(a / 0) // runtime error: invalid operation: division by zero
	// 正常做除法的时候，都判断一下！！！
	if b != 0 {
		println(a / b)
	}

	// 取余
	println(a % b)
	//println(a % 0) // runtime error: invalid operation: division by zero
	// 正常做取余的操作时，也都判断一下
	if b != 0 {
		println(a % b)
	}

	// Go 里面，只有同类型才可以执行加减乘除。
	// 某些语言有自动类型转换，Go 是没有的。
	var c float64 = 12.3
	//println(a + c) // compile error: Invalid operation: a + c (mismatched types int and float64)
	println(a + int(c))
	println(float64(a) + c)

	var d int32 = 12
	//println(a + d) // compile error: Invalid operation: a + d (mismatched types int and int32)
	println(a + int(d))

	// 更加复杂的运算都在 math 包里面，包括三角函数、聚合函数等。复杂数学操作找 math 包
	//math.Ceil()
}

// Extremum 极值
func Extremum() {
	/**
	数字类型的极值都在 math 包里面，作为常量。
	• int 和 uint 族都有最大值和最小值。
	• float32 和 float64 只有最大值和最小正数，
	没有最小值。
	注意，低版本的 Go SDK 不一定有全部的极值的常量。
	*/

	println("float64 最大值", math.MaxFloat64)
	// 没有 float64 最小值
	//println("float64 最小值", math.MinFloat64)
	println("float64 最小的正数", math.SmallestNonzeroFloat64)

	println("float32 最大值", math.MaxFloat32)
	// 没有 float32 最小值
	//println("float32 最小值", math.MinFloat32)
	println("float32 最小的正数", math.SmallestNonzeroFloat32)

	// int 族和 uint 族都有最大值最小值
	println(math.MaxInt)
}
