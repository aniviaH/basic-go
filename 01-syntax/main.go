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
}
