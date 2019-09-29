package main

import "fmt"

func main() {
	// 方式1
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()                 // 调用匿名函数
	fmt.Printf("%T", a) // a的类型为func()，即匿名函数的签名

	// 方式2
	/*
		func() {
			fmt.Println("hello world first class function")
		}()  // 直接运行匿名函数
	*/
}
