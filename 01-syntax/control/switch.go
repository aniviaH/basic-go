package main

/**
switch 语句和别的语言类似。基本语法形态是：
switch val {
	case xxx:
	case yyy:
	default:
}

default 分支可以有也可以没有。 switch 的值必须是可比较的。实践中，不能用于 switch 的值，编译器会报错。

注意，switch 语句不需要 break。
*/

func Switch(status int) {
	switch status {
	case 0:
		println("初始化")
	case 1:
		println("运行中")
	default:
		println("未知状态")
	}

	println("Switch End!")
}

/**
switch 语句也可以没有 val:
switch {
	case condition0:
	case condition1:
}
这种情况下，case 后面跟 bool 表达式。
*/

func SwitchBool(age int) {
	switch {
	case age >= 18:
		println("成年人")
	case age > 12:
		println("少年")
	default:
		println("未知状态")
	}

	println("SwitchBool End!")
}

func SwitchAny(age any) {
	switch age {
	case "18":
		println("成年人")
	case 18:
		println("成年人")
	}

	println("SwitchAny End!")
}
