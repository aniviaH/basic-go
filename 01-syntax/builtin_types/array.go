package main

import "fmt"

func Array() {
	arr := [3]int{9, 8}
	fmt.Printf("arr: %v, len=%d, cap=%d\n", arr, len(arr), cap(arr))

	var arr3 [3]int
	fmt.Printf("arr3: %v, len=%d, cap=%d\n", arr3, len(arr3), cap(arr3))

	// 数组不支持 append 操作
	//arr3 = append(arr3, 1) // Cannot use 'arr3' (type [3]int) as the type []Type

	// 下标访问，编译器能识别出下标是否越界，会编译报错。如果不能，则会在运行时候报错出现panic
	fmt.Printf("arr[0]: %d\n", arr[0])
}
