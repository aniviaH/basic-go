package main

import (
	"errors"
	"fmt"
)

/**

Go 的泛型语法很简单。

右边是一个泛型接口，其中类型参数 T 可以
是任意类型。

*/

// T 类型参数，名字叫T，约束是 any，等于没有约束

type List[T any] interface {
	Add(idx int, val T)
	Append(val T)
}

type ListInt interface {
	Add(idx int, val int)
	Append(val int)
}

type ListFloat interface {
	Add(idx int, val float64)
	Append(val float64)
}

func UseList() {
	//var l List // Use of generic type without instantiation
	var l List[int]
	l.Append(10)
	//l.Append(12.3) // 编译会报错：cannot use 12.3 (untyped float constant) as int value in argument to l.Append (truncated)
	//l.Append("hello") // Cannot use '"hello"' (type string) as the type int64

	var lany List[any]
	lany.Append(12.3)
	lany.Append("hello")

	lk := LinkedList[int]{}
	intVal := lk.head.val
	println(intVal)
}

type LinkedList[T any] struct {
	//head *node[int]
	head *node[T]
	t    T
}

type node[T any] struct {
	val T
}

func Max[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		var t T
		return t, errors.New("未传入数据进行查找最大值")
	}

	var res T

	for i := 0; i < len(vals); i++ {
		if vals[i] > res {
			res = vals[i]
		}
	}

	return res, nil
}

func AddSlice[T any](s []T, idx int, val T) ([]T, error) {
	//var ss = []int{1, 2, 3, 4}

	// 如果 idx 是负数，或者超过了 slice 的界限
	if idx < 0 || idx > len(s) {
		return nil, errors.New("下标出错")
	}

	res := make([]T, 0, len(s)+1)
	// 写法1
	//for i, v := range s {
	//	if i < idx {
	//		res = append(res, v)
	//	} else if i == idx {
	//		res = append(res, val, v)
	//	} else {
	//		res = append(res, v)
	//	}
	//}

	// 写法2
	//for i := 0; i < idx; i++ {
	//	res = append(res, s[i])
	//}
	//res = append(res, val)
	//for i := idx; i < len(s); i++ {
	//	res = append(res, s[i])
	//}

	// 写法3
	//res = append(s, s[0:idx]..., val, s[idx: (len(s) + 1)]...)
	res = append(res, s[0:idx]...)
	res = append(res, val)
	res = append(res, s[idx:len(s)]...)

	// 切片可能发生扩容，需要将其返回出去
	return res, nil
}

func AddSliceV2[T any](s []T, idx int, val T) ([]T, error) {
	// 如果 idx 是负数，或者超过了 slice 的界限
	if idx < 0 || idx > len(s) {
		return nil, errors.New("下标出错")
	}

	// 原切片的前半部分
	sliceStart := s[0:idx]
	var res = make([]T, 0, len(s)+1)

	res = append(res, sliceStart...)
	res = append(res, val)

	// 原切片的后半部分
	sliceEnd := s[idx:len(s)]
	res = append(res, sliceEnd...)

	return res, nil
}

func AddSliceV3[T any](s []T, idx int, val T) ([]T, error) {
	if idx < 0 || idx > len(s) {
		return nil, errors.New("下标出错")
	}

	// 先对s扩展一个元素
	var zeroValue T
	s = append(s, zeroValue)
	// 从后往前遍历，直到插入位置的下一位置
	for i := len(s) - 1; i > idx; i-- {
		if i-1 > 0 {
			s[i] = s[i-1]
		}
	}
	// 最后赋值要插入的元素
	s[idx] = val
	return s, nil
}

func UseAddSlice() {
	s1, err := AddSlice([]int{1, 2, 3, 4, 5}, 2, 10) // 期望：[1, 2, 10, 3, 4, 5]
	if err != nil {
		fmt.Println("切片添加异常", err)
	}
	fmt.Println(s1)

	s2, err := AddSliceV2([]int{1, 2, 3, 4, 5}, 2, 10)
	if err != nil {
		fmt.Println("切片添加异常", err)
	}
	fmt.Println(s2)

	s3, err := AddSliceV3([]int{1, 2, 3, 4, 5}, 2, 10)
	if err != nil {
		fmt.Println("切片添加异常", err)
	}
	fmt.Println(s3)
}
