package main

import "fmt"

func Loop1() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	// 这样也可以
	for i := 0; i < 10; {
		println(i)
		i++
	}
}

/**
第二种写法类似于别的语言中的 while 循环，基本语法是：
for condition {
// 循环内部操作
}
*/

func Loop2() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//for ; i < 10; i++ {
	//	println(i)
	//}

	// 死循环
	//for  {
	//	println(i)
	//	i++
	//}
}

func ForArr() {
	arr := [3]int{1, 2, 3}
	for index, val := range arr {
		println("下标：", index, "值：", val)
	}
	for index := range arr {
		println("下标：", index, "值：", arr[index])
	}
}

func ForSlice() {
	slice := []int{1, 2, 3}
	for index, val := range slice {
		println("下标：", index, "值：", val)
	}
	for index := range slice {
		println("下标：", index, "值：", slice[index])
	}
}

func ForMap() {
	m := map[string]int{
		"key1": 100,
		"key2": 102,
		"key3": 104,
		"key4": 106,
		"key5": 108,
		"key6": 110,
	}
	for key, val := range m {
		println("键：", key, "值：", val)
	}
	for key := range m {
		println("键：", key, "值：", m[key])
	}

	// !注：对map的遍历是随机的！不要指望在map遍历中按照key的定义顺序做逻辑
	// 对顺序有要求的场景，需要将map转为有序的数据结构
}

type User struct {
	Name string
}

func LoopBug() {
	/**
	！！！千万不要对迭代参数取地址！！！
	在内存里面，迭代参数都是放在一个同一个地方的，你循环开始就确定了，所以你一旦取地址，那么你拿到的就是这个地址。
	所以，右边的 map 中的键值对的值，最终都是同一个，也就是最后一个。
	*/
	users := []User{
		{
			Name: "张三",
		},
		{
			Name: "李四",
		},
	}
	m := make(map[string]*User)
	for _, v := range users {
		fmt.Printf("%p\n", &v) // 迭代里，v的地址始终都是同一个固定的
		m[v.Name] = &v
	}
	fmt.Printf("%v\n", m)

	// 所以for循环都是这样！
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t%p\n", i, &i)
	}
}

func LoopBug2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t%p\n", i, &i)
	}
}

func LoopBreak() {
	i := 0
	for true {
		if i > 10 {
			break
		}
		i++
	}
}

func LoopContinue() {
	i := 0
	for ; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		println(i)
	}
}
