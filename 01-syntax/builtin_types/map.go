package main

import "fmt"

func Map() {
	m1 := map[string]int{
		"key1": 123, // 最后需要一个尾分号：Need a trailing comma before a newline in the composite literal
	}
	m1["hello"] = 345

	fmt.Println(m1["hello"])

	// 第二个参数: 容量
	m2 := make(map[string]int, 12)
	m2["key2"] = 678
	fmt.Println(m2["key2"])

	// 空map不支持赋值
	var m3 map[string]int // nil map
	//m3["key3"] = 890      // panic: assignment to entry in nil map
	fmt.Println(m3["key3"]) // 但可以访问，访问任何key都是其类型的零值

	/**
	读取元素：有两个返回值
	• 第一个是值，第二个是这个元素是否存在。
	• 如果只用一个返回值，那么就是对应的元素；元素
	不存在，那么就是对应类型的零值
	*/
	val, ok := m3["key3"] // 推荐写法
	if !ok {
		// 没有这个键值对
		fmt.Println("key3 not exist in m3, 零值为：", m3["key3"])
	}
	fmt.Println("m3[key3]: ", val)

	val2 := m3["hello"]
	fmt.Println("m3[hello]: ", val2)
}

/*
*
读取长度：len
• 遍历：for，第一个是 key，第二个是 value
• 删除：使用 delete 方法
注意：map 的遍历是随机的，也就是说你遍历两遍，输出的结果都不一样。
*/

func Map2() {
	m := make(map[string]string, 4)
	m["key1"] = "value1"
	m["key2"] = "value2"
	m["key3"] = "value3"
	m["key4"] = "value4"

	println(len(m))

	for k, v := range m {
		println(k, v)
	}
	for k := range m {
		println(k)
	}

	delete(m, "key1")
	println(len(m))
}

type User struct {
	name string
	//abc int // 可比较
	abc [1]int // 如果数组内的项是可比较的，则数组也是可比较
	//abc  []int // 但切片始终是不可比较
}

func DemoComparable(u User) {
	switch u {
	case User{}:
		fmt.Println("case User{}")
	case User{name: "张三", abc: [1]int{1}}:
		fmt.Println("case User{name: 张三}")
	case User{name: "liuh"}:

	}
}
