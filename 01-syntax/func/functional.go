package main

func Functional4() string {
	println("hello, functional 4")

	return "hello"
}

func Functional5(age int) {

}

var Abc = func() string {
	return "hello"
}

func UseFunctional4() {
	myFunc := Functional4
	myFunc()

	//Abc = func(a int) string { // Cannot use 'func (a int) string { }' (type func(a int) string) as the type func() string
	//
	//}

	myFunc5 := Functional5
	myFunc5(18)
}

func Functional6() {
	// 局部函数，作用域只在当前函数内
	fn := func() string {
		return "hello"
	}
	fn()
}

// Functional7 函数作为返回值 返回一个 func() string
func Functional7() func() string {
	return func() string {
		return "hello, world"
	}
}

// Functional8 匿名函数
func Functional8() {
	/*
		在方法内部可以声明一个匿名方法，但是需要立刻发起调用。
		为什么需要立刻发起调用？因为匿名，即没有名字，不立刻调用的话后面你都没办法调用了。
	*/
	hello := func() string {
		return "hello"
	}()
	println(hello)
}
