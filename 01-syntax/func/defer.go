package main

import "fmt"

func Defer() (num int) {
	num = 0

	defer func() {
		println("第一个defer", num)
		num = 10
	}()

	defer func() {
		println("第二个defer", num)

		num = 20
	}()

	num++

	return num
}

func DeferReturnV1() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return
}

type MyStruct struct {
	name string
}

func DeferReturnV3() MyStruct {
	a := MyStruct{name: "Jerry"}
	defer func() {
		a.name = "Tom"
	}()
	return a
}

func DeferReturnV4() *MyStruct {
	a := &MyStruct{name: "Jerry"}
	defer func() {
		a.name = "Tom"
	}()
	return a
}

func DeferReturnV5() (a MyStruct) {
	a = MyStruct{name: "Jerry"}
	defer func() {
		a.name = "Tom"
	}()
	return
}

//func Query() {
//	db, _ := sql.Open("", "")
//	defer db.Close()
//	db.Query("SELECT ")
//}

func DeferClosureLoop1() {
	for i := 0; i < 10; i++ {
		defer func() { // Possible resource leak, 'defer' is called in the 'for' loop
			//println(i)
			fmt.Printf("%p\t%d\n", &i, i)
		}()
	}
}

func DeferClosureLoop2() {
	for i := 0; i < 10; i++ {
		defer func(i int) { // Possible resource leak, 'defer' is called in the 'for' loop
			//println(i)
			fmt.Printf("%p\t%d\n", &i, i)
		}(i)
	}
}

func DeferClosureLoop3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() { // Possible resource leak, 'defer' is called in the 'for' loop
			//println(j)
			fmt.Printf("%p\t%d\n", &j, j)
		}()
	}
}
