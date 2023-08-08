package main

import "fmt"

// YourName 不定参数
func YourName(name string, alias ...string) {
	// alias 是一个切片
	if len(alias) > 0 {
		println("别名：", alias[0])
	}
}

func CallYourName() {
	YourName("liuhuan")
	YourName("liuhuan", "huan", "ahuan")

	// 使用切片方式传入不定参数
	alias := []string{"ahuan", "dahuan"}
	YourName("liuhuan", alias...)

	// 不定参数里，注意 parameter ...any，可以单独传一个切片，而不用展开，其会将整个切片当做一个item
	YourNameV1("liuhuan", alias)
}

func YourNameV1(name string, alias ...any) {
	fmt.Println(name, alias)
}
