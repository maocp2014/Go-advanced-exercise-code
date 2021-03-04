package main

import "fmt"

func foo() func() {
	x := 1
	f := func() {
		fmt.Printf("foo0 val = %d\n", x)
	}
	x = 11
	return f
}

func main() {
	// 在执行的时候去外部环境寻找最新的数值
	foo()() // 11
}
