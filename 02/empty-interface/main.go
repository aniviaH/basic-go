package main

import "fmt"

func printValue(v interface{}) {
	//fmt.Println(v)
	fmt.Printf("type: %T, val: %v \n", v, v)
}

func printValues(v ...interface{}) {
	//fmt.Println(v)
	fmt.Printf("type: %T, val: %v \n", v, v)
}

func main() {
	printValue(42)      // int
	printValue("Hello") // string
	printValue(3.14)    // float64

	printValues(999, "world")
}

/*

提问：

An empty interface 'interface{}' stands for a set of all types. A variable of the empty interface type can store a value of any type

GPT回答：

Yes, that's correct. In Go, the empty interface type `interface{}` is known as the empty interface or the interface{} type. It is a special type that represents a set of all types. A variable of the empty interface type can hold values of any type because every type in Go implements at least zero methods, and therefore, satisfies the empty interface.

This flexibility allows you to write functions or create data structures that can accept values of any type. For example, you can use an empty interface to create a collection that can hold elements of different types, or you can use it as a function parameter to accept arguments of different types.

However, when you store a value of a specific type in an empty interface variable, the type information is lost. You can only perform limited operations on the value without knowing its actual type. To work with the value as its original type, you need to use type assertions or type switches to extract the underlying value and its type.

Here's an example to illustrate the usage of an empty interface:

```go
func printValue(v interface{}) {
    fmt.Println(v)
}

func main() {
    printValue(42)          // int
    printValue("Hello")     // string
    printValue(3.14)        // float64
}
```

In the `printValue` function, the parameter `v` has the type `interface{}`, so it can accept values of any type. Inside the function, the value is printed without knowing its actual type. When you call the function with different values, it will print the corresponding types.

*/
