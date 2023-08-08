package main

import "fmt"

func Closure(name string) func() string {
	return func() string {
		return "hello, " + name
	}
}

var age = 18

func Closure2() func() string {
	name := "ani"
	age := 18
	return func() string {
		return fmt.Sprintf("Hello, %s, %d", name, age)
	}
}

func Closure3() func() int {
	var num = 0
	fmt.Printf("out num: %p\n", &num)
	return func() int {
		fmt.Printf("inner num before: %p\n", &num)
		num++
		fmt.Printf("inner num after: %p\n", &num)
		return num
	}
}
