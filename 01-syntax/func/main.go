package main

func main() {
	//res, err := Func3()
	//fmt.Println(res, err)

	/**
	方法调用：接收返回值
	方法调用只需要用足够的参数去接收返回值，并且传入了调用参数就可以。
	在多个返回值的时候，如果你想忽略一些返回值，可以使用 _
	记住一条核心原则：同一个作用域内，如果左边出现了新的变量，那么就需要使用 := 来接收返回值。
	*/

	//str, err := Func2(1, 2)
	//println(str, err)
	// 忽略返回值
	//_, _ = Func2(1, 3)

	// 部分忽略返回值
	// str 是已经声明好了
	//str, _ = Func2(3, 4)
	// str1 是新变量，需要使用 :=
	//str1, _ := Func2(3, 4)
	//println(str1, err)

	// str2 是新变量，需要使用 :=
	// 使用 := 的前提，就是左边必须有至少一个新变量
	//str2, _ := Func2(3, 4)
	//println(str2)

	//Recursive(0)

	//UseFunctional4()

	//returnFn := Closure("ani")
	//hello := returnFn()
	//println(hello)

	//getNum := Closure3()
	//var num = 0
	//num = getNum()
	//println(num)
	//num = getNum()
	//println(num)
	//num = getNum()
	//println(num)

	//YourName("liuhuan", "ahuan", "huan")
	//CallYourName()

	//res := Defer()
	//fmt.Println(res)
	//res1 := DeferReturnV1()
	//fmt.Println(res1)
	//res2 := DeferReturnV2()
	//fmt.Println(res2)
	//res3 := DeferReturnV3()
	//fmt.Println(res3)
	//res4 := DeferReturnV4()
	//fmt.Println(res4)
	//res5 := DeferReturnV5()
	//fmt.Println(res5)

	//DeferClosureLoop1()
	DeferClosureLoop2()
	//DeferClosureLoop3()
}
