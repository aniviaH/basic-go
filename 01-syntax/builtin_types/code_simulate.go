package main

import "fmt"

/**
代码演示
1. Go 是强类型语言，能不能设计一个方法，可以计算任意数字类型切片的和的方法？
• func SumInt64([]int64) 只能用于 int64
• func SumInt32([]int32) 只能用于 int32
• ...
2. 获得 map 的所有 key、所有 value。
*/

func UseSumInt64() {
	s1 := []int{1, 2, 3, 4}
	//res1 := SumInt64(s1) // s1是[]int切片，不能传给SumInt64。所以需要再写一个函数，处理[]int切片求和
	res := SumInt(s1)
	fmt.Printf("s1的所有项的和为：%d\n", res)

	// 那如果切片的项的类型是int32、int8、uint、float64...???

	// 可以方式1：进行类型转换
	//var s2 []int64 = s1.([]int64) // 类型断言：Invalid type assertion: s1.([]int64) (non-interface type []int on the left)
	//var s2 []int64 = ([]int64)(s1) // Cannot convert an expression of the type '[]int' to the type '[]int64'
	// []int 和 []int64 是两个完全不同的类型。如果要类型转换，老老实实一个个项进行转换
	var s2 []int64 = make([]int64, len(s1))
	for _, val := range s1 {
		s2 = append(s2, int64(val))
	}
	res64 := SumInt64(s2) // 这里可以编译通过，上面不能
	fmt.Printf("s2的所有项的和为：%d\n", res64)

	//更好的方案是可以使用泛型
}

func UseMapKeys() {
	// map也同样面临这样的问题，见MapStringKeys、MapIntKeys
	m1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	//m1Keys := MapIntKeys(m1) // m1不能传给MapIntKeys
	keys := MapStringKeys(m1)
	fmt.Printf("keys: %v\n", keys)
	//keys2 := MapAnyKeys(m1) // Cannot use 'm1' (type map[string]int) as the type map[any]any

	//keys3 := KeysByGeneric(m1)
}

func SumInt64(vals []int64) (sum int64) {
	//var sum int64
	for _, val := range vals {
		sum += val
	}
	return sum
}

func SumInt(vals []int) (sum int) {
	//var sum int
	for _, val := range vals {
		sum += val
	}
	return sum
}

func MapStringKeys(m map[string]int) []string {
	var res []string
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

func MapIntKeys(m map[int]string) []int {
	var res []int
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

func MapAnyKeys(m map[any]any) []any {
	var res []any
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

func SumByGeneric[T any](s T) int {
	return 0
}

func KeysByGeneric[T any](m T) []string {
	return []string{"hello"}
}
