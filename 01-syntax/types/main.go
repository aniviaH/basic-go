package main

import "fmt"

func main() {
	//var list List
	//list.Add(1, 1)

	//NewUser()

	//ChangeUser()

	// 衍生类型
	//UseInt()
	//UseFish()

	// 类型别名
	//UseYu()

	// 面向接口编程：
	var l List
	l = &ArrayList{}
	l = &LinkedList{}
	fmt.Println(l)
	DoSomething(l)
}

func DoSomething(l List) {
	l.Append(12.3)
	l.Append(10)
	l.Append("hello")
}
