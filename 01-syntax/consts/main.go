package main

const External = "包外"
const internal = "保内"

const (
	a = 123
	b = 456
)

const (
	StatusA int64 = iota
	StatusB
	StatusC

	// StatusD 插入一个主动赋值的就中断了iota
	StatusD = 7
	StatusE
	StatusE2

	StatusF = iota
)

const (
	A = iota*2 + 3
	B
)

const (
	Init = iota
	Running
	Paused
	Stop
)

const (
	// DayA 0 左移 -- 0000 => 0000(0)
	DayA = iota << 1
	// DayB 1 左移 -- 0001 => 0010(2)
	DayB
	// DayC 2 左移 -- 0010 => 0100(4)
	DayC
	// DayD 3 左移 -- 0011 => 0110(6)
	DayD
	// DayE 4 左移 -- 0100 => 1000(8)
	DayE
	DayF
)

const (
	// MonthA 0 右移 -- 0000 => 0001
	MonthA = 1 << iota
	// MonthB 1 右移 -- 0001 => 0010
	MonthB
	// MonthC 2 右移 -- 0010 => 0100
	MonthC
	// MonthD 3 右移 -- 0011 => 1000
	MonthD

	MonthE = iota
)

func main() {
	const a = 123

	//a = 234 // Cannot assign to a

	println(StatusF)

	var num = 0 << 1
	println(num)
	num = num << 1
	println(num)
	println(1 << 3)
}
