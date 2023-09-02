package main

import "fmt"

func NewUser() {
	// 初始化结构体
	u1 := User{}
	//println(u1) // print只能打印基本类型、内置类型
	/**
	.\user.go:6:9: illegal types for operand: print
	User
	*/
	fmt.Printf("%v\n", u1)  // %v: the value in a default format
	fmt.Printf("%+v\n", u1) // %+v: when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%#v\n", u1) // %#v: a Go-syntax representation of the value

	u1.Name = "张三"
	fmt.Printf("%+v\n", u1)

	up := &User{}
	//println(up)
	fmt.Println(up)
	fmt.Printf("%v \n", up)

	up2 := new(User)
	println("up2:", up2.Name)
	fmt.Println(up2)
	fmt.Printf("%+v \n", up)

	u4 := User{Name: "Tom", Age: 0}
	u5 := User{"Jerry", "J", 1}
	fmt.Println(u4, u5)

	var up3 *User
	println(up3)
	fmt.Println(up3)
	// nil 上访问字段或者防范，会panic
	//fmt.Println(up3.FirstName) // panic: runtime error: invalid memory address or nil pointer dereference

	l1 := LinkedList{}
	l1ptr := &l1
	var l2 LinkedList = *l1ptr
	fmt.Println(l2)

}

type User struct {
	Name      string
	FirstName string
	Age       int
}

func (u User) ChangeName(name string) {
	fmt.Printf("change name 中 u 的地址：%p \n", &u)
	u.Name = name
}

func ChangeName(user User, name string) {

}

func (u *User) ChangeAge(age int) {
	fmt.Printf("change age 中 u 的地址：%p \n", u)
	u.Age = age
}

func ChangeAge(u *User, age int) {

}

func ChangeUser() {
	var u1 User = User{Name: "张三", Age: 18}
	fmt.Printf("调用的地方 u1 的地址：%p \n", &u1)
	// 值传递：这一步执行的时候，其实相当于复制了一个 u1, 改的是复制体
	// 指针接收器和值接收器调用时，使用指针和结构体都行，编译器会进行相应解释 https://go.dev/tour/methods/6
	u1.ChangeName("张三三")
	u1.ChangeAge(35) // 等于 (&u1).ChangeAge
	fmt.Println(u1)

	var u2 *User = &User{Name: "李四", Age: 20}
	fmt.Printf("调用的地方 u2 的地址：%p \n", u2)
	u2.ChangeName("李四四") // 等于 (*u2).ChangeName
	u2.ChangeAge(45)
	fmt.Println(u2)
}
