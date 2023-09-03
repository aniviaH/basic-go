package main

import (
	"fmt"
	"time"
)

/**
衍生类型

基本语法：type TypeA TypeB

衍生类型可能使用的场景：一般在想使用第三方库，又没有办法修改源码，又想
扩展这个库的结构体的方法的情况下，就会用这个。

记住核心：衍生类型是一个全新的类型。

衍生类型可以互相转换，使用 （） 进行转换。

注意， TypeB 实现了某个接口，不等于 TypeA 也实现
了某个接口。

*/

type Integer int // 衍生类型 - 定义一个新类型

func UseInt() {
	i1 := 10
	i2 := Integer(i1)
	var i3 Integer = 11
	//i1 = i3 // Cannot use 'i3' (type Integer) as the type in
	//i4 := int(i2)
	println(i2, i3)
	fmt.Printf("%T\n", i3)
}

type Fish struct {
	Name string
}

func (f Fish) Swim() {
	fmt.Println("fish is swimming~")
}

type FakeFish Fish

func (ff FakeFish) FakeSwim() {
	fmt.Println("fakefish is fake swimming~")
}

func UseFish() {
	f1 := Fish{
		Name: "fish-1",
	}
	f1.Swim()

	//f2 := FakeFish{}
	f2 := FakeFish(f1)
	// f2 将不能调用 Fish 上的方法，因为 f2 是一个全新的类型
	//f2.Swim()
	f2.FakeSwim()
	fmt.Println(f2.Name)
	f2.Name = "Tom"
	fmt.Println(f2.Name)
	fmt.Println(f1.Name)

	// 类型转换
	f3 := Fish(f2)
	f3.Swim()
}

// 比如要使用某个第三方包，但是其内部又不能进行修改，就可以通过一个衍生类型来扩展
// 我想在第三方包上面定义一些其他方法，但我定义不了（为什么定义不了：它没有暴露出来）
type MyTime time.Time

// 通过衍生类型，间接定义一个方法
// 但其不能增加字段，也不能修改到字段的值
func (m MyTime) MyFunc() {

}

/**

类型别名：

基本语法：type TypeA = TypeB
• 别名，除了换了一个名字，没有任何区别
• 它和衍生类型的区别，就是用了 =
类型别名一般用在导出类型、兼容性修改里面，也不常
见。

*/

// 向后兼容

// Yu 鱼
// Yu 是 Fish 的别名
type Yu = Fish

func (f Yu) YuYouYong() {
	fmt.Println("鱼在游泳~")
}

func UseYu() {
	f1 := Fish{}
	var yu Yu = Fish{}
	//var yu2 Yu = FakeFish{} // Cannot use 'FakeFish{}' (type FakeFish) as the type Yu
	var yu3 Yu

	f1.Swim()
	yu.Swim()
	yu3.Swim()

	yu.YuYouYong()
	f1.YuYouYong()
}
