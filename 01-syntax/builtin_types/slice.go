package main

import "fmt"

/**
切片
切片，语法是：[]type
1. 直接初始化
2. make初始化:make([]type, length,
capacity)
3. arr[i] 的形式访问元素
4. append 追加元素
5. len 获取元素数量
6. cap 获取切片容容量
7. 推荐写法：s1 := make([]type, 0,capacity)
8. 使用 for range 来遍历

最佳实践：在初始化切片的时候要预估容量。
*/

func Slice() {
	s1 := []int{9, 8, 7}
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	s2 := []int{9, 8}
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 3, 4)
	fmt.Printf("s3: %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	s4 := make([]int, 4) // {0, 0, 0, 0}
	fmt.Printf("s4: %v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	s5 := make([]int, 0, 4)
	s5 = append(s5, 1)
	fmt.Printf("s5: %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
}

/**
子切片
数组和切片都可以通过[start:end] 的形式来获取子切片：
1. arr[start:end]，获得[start, end)之间的元素。
2. arr[:end]，获得[0, end) 之间的元素。
3. arr[start:]，获得[start, len(arr))之间的元素。

都是左闭右开！
*/

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	//s2 := s1[1:2]
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	// 内存共享 - 只有一个判断点，有没有扩容
	//s2 = append(s2, 99, 98)

	// 容量就是 start 开始往后，包括原本 s1 的底层数组的个数
	s3 := s1[2:]
	fmt.Printf("s3: %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	s5 := s4[:]
	fmt.Printf("s5: %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
	s5 = append(s5, 12, 14, 16)
	fmt.Printf("s5: %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
}

/**
内存共享问题
核心：共享数组。
子切片和切片究竟会不会互相影响，就抓住一点：它们是不是还共享数组？
什么意思？
• 就是如果它们结构没有变化，那肯定是共享的；但是结构变化了，就可能不是共享了。
什么情况下结构会发生变化？扩容了。
所以，切片与子切片，切片作为参数传递到别的方法、结构体里面，任何情况下你要判断是否内存共享，那么就一个点：有没有扩容。
*/

func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	//s2 := s1[2:3]
	fmt.Printf("s1: %v, len=%d, cap=%d\t", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s1: %v, len=%d, cap=%d\t", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("s1: %v, len=%d, cap=%d\t", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("s1: %v, len=%d, cap=%d\t", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
}

// 一些原则
func Input(arr []int) {
	// 你不要直接修改 arr。可以复制一份，修改复制的那份，返回复制的这份
	//arr[0] = 1

}
