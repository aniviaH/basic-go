package main

import "fmt"

func BoolUse() {
	var a bool = true
	var b bool = false
	var c bool = a || b // true
	fmt.Println(c)
	var d bool = a && b // false
	fmt.Println(d)
	var e bool = !a // false
	fmt.Println(e)

	var f = !(a && b) // !a || !b  true
	var g = !(a || b) // !a && !b  false
	fmt.Println(f)
	fmt.Println(g)
}
