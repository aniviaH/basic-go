package main

func main() {
	//UseInner()

	var o1 OuterV1
	o1.DoSomething()
	o1.Inner.DoSomething()

	// 组合不是继承，没有多态。
	var o Outer
	// 输出什么呢？
	// hello, outer ?
	// hello, inner ?
	o.SayHello() // hello, inner
}
