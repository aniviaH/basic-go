package main

import (
	"fmt"
)

func ByteUse() {
	var a byte = 'a'
	fmt.Println(a)
	fmt.Println(fmt.Sprintf("%c\n", a))

	// byte相关操作找bytes包，一般操作的是byte切片，单个byte就是uint8，没什么好操作的
	//bytes.Contains()

	// 与string互相转换
	var str string = "this is string"
	var bs []byte = []byte(str)
	bs[0] = 'a'
	var bs2 []byte = []byte{'a', 'b', 'c'}
	fmt.Println(str, bs, bs2)
}
