package main

func main() {
	DeferCondition(false)
	DeferCondition(true)
}

func DeferCondition(input bool) {
	if input {
		defer func() {
			println("hello, world 1")
		}()
		defer func() {
			println("hello, world 2")
		}()
		defer func() {
			println("hello, world 3")
		}()
		defer func() {
			println("hello, world 4")
		}()
		defer func() {
			println("hello, world 5")
		}()
		defer func() {
			println("hello, world 6")
		}()
		defer func() {
			println("hello, world 7")
		}()
		defer func() {
			println("hello, world 8")
		}()
		defer func() {
			println("hello, world 9")
		}()
		defer func() {
			println("hello, world 9")
		}()
	}
}
