package main

import "fmt"

/**

组合

组合是 Go 里面独有的语法概念。基本语法形态是：

type A struct {
B
}

组合可以是以下几种情况：
• 接口组合接口
• 结构体组合结构体
• 结构体组合结构体指针
• 结构体组合接口
可以组合多个。

*/

type Inner struct {
	//Outer // Invalid recursive type 'Inner' Inner → Outer → Inner
	Name string
}

func (i Inner) DoSomething() {
	fmt.Println("这是Inner DoSomething")
}

func (i Inner) SayHello() {
	println("hello,", i.GetName())
}

func (i Inner) GetName() string {
	return "inner"
}

func (i *Inner) ChangeName() {
	i.Name = "Jerry"
}

// 用这个!!!
type Outer struct {
	Name string
	Inner
}

func (o Outer) GetName() string {
	return "outer"
}

type OuterV1 struct {
	Inner
}

func (i OuterV1) DoSomething() {
	fmt.Println("这是Outer DoSomething")
}

type OuterPtr struct {
	Name string
	*Inner
}

type OOOuter struct {
	Name string
	Outer
}

func UseInner() {
	var o Outer
	// 因为 Outer 组合了 Inner，所以虽然它自己没有方法，但它可以调 Inner 里的方法
	o.DoSomething()
	o.Inner.DoSomething()

	//var op OuterPtr = OuterPtr{
	//	Name: "Tom",
	//}
	//op.DoSomething()

	var ooo OOOuter
	ooo.DoSomething()

	var o1 = Outer{
		Name:  "Tom",
		Inner: Inner{Name: "Tom2"},
	}
	o1.ChangeName()
	o1.Inner.ChangeName()
	fmt.Println("修改后o1.name=", o1.Name)
	fmt.Println("修改后o1.inner.name=", o1.Inner.Name)

	op1 := OuterPtr{
		Name:  "张三",
		Inner: &Inner{Name: "张三2"},
	}
	op1.DoSomething()
	op1.Inner.DoSomething()

	ooo1 := OOOuter{
		Name: "李四",
	}
	ooo1.DoSomething()
}
