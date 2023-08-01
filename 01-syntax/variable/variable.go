package variable

import "fmt"

var Global = "全局变量"
var internal = "包内变量，私有变量"

var (
	First  string = "First..."
	second int    = 2
)

//var First = "First......" // 'First' redeclared in this package

// := 只能用于局部变量，即方法内部
//xx := 134 // 'xx' unexpected

func main() {
	var a int = 123
	println(a)

	var b = 234
	println(b)

	//b := 345 // No new variables on the left side of ':='

	var First = "First!!!" // 与作用域外同名是可以的，但idea会给你显示为绿色，建议不要同名
	println(First)

	var c uint = 456
	println(c)

	//println(a == c) // Invalid operation: a == c (mismatched types int and uint)

	var (
		d string = "aaa"
		e int    = 123
	)
	fmt.Println(d, e)

	f := 123
	println(f)

}
