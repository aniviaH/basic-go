package main

import (
	"fmt"
	"github.com/aniviaH/basic-go/01-syntax/variable"
)

func main() {
	fmt.Println(variable.Global)
	//fmt.Println(variable.internal) // Unexported variable 'internal' usage
	fmt.Println(variable.First)
	//fmt.Println(variable.second) // Unexported variable 'second' usage

	//var s []func()
	//for i := 0; i < 10; i++ {
	//	var f = func() {
	//		fmt.Printf("%p %v\n", &i, i)
	//	}
	//	s = append(s, f)
	//}
	//
	//for _, f := range s {
	//	f()
	//}

	//var list types.List
	//list.Add(1, 1)
}
