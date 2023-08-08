package main

// Func0 单一返回值
func Func0(name string) string {
	return "hello, world"
}

// Func1 多个返回值
func Func1(a, b, c int, str1 string) (string, error) {
	return "", nil
}

// Func2 带名字的返回值
func Func2(a int, b int) (str string, err error) {
	str = "hello"
	// 带名字的返回值，可以直接return
	return
}

func Func3() (str string, err error) {
	res := "hello"
	str = "world"
	// 虽然带名字，但是我们并没有用
	return res, nil
}
