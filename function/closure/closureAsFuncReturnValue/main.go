package main

import "fmt"

// Operation 函数类型
type Operation func() int

// Increase 匿名函数（闭包）作为函数返回值
func Increase() Operation {
	n := 0
	return func() int {
		n++
		return n
	}
}

// 匿名函数作为返回值，不如理解为闭包作为函数的返回值
// 闭包被返回赋予一个同类型的变量时，同时赋值的是整个闭包的状态，
// 该状态会一直存在外部被赋值的变量in中，直到in被销毁，整个闭包也被销毁。
func main() {
	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
	fmt.Println(in())
	fmt.Println(in())
}

// 1
// 2
// 3
// 4