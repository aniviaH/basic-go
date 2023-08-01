package main

import (
	"fmt"
	"unicode/utf8"
)

/**
string 类型
有两种写法：
• 使用双引号 “ 引起来。如果在字符串里面的
“ 就需要用 \ 来进行转义。
• 使用反引号 ` 引起来。这种写法可以换行，但是
内部不能有反引号（转义也不行）。

Tip：不建议自己手写转义，而是自己先写好，然后复制到 Goland，IDE 会自动完成转义。
*/

func String() {
	println("He said: \"hello, Go!\"")
	println(`我可以换行
	这是新的行
	但是这里不能有反引号"sss\"
`)

	// He said: "hello Go!"
	// 复制上面的字符串到idea里
	println("He said: \"hello Go!\"")

	println("hello" + "go")
	println("hello" + string(123)) // hello{  这里会将123解释为其对应的ascii码
	println(fmt.Sprintf("hello %d", 123))

	println(len("abc"))                      // 3
	println(len("你好"))                       // 6
	println(utf8.RuneCountInString("你好"))    // 3
	println(utf8.RuneCountInString("你好abc")) // 5

	// string相关操作找strings包
	//strings.Cut()

}
