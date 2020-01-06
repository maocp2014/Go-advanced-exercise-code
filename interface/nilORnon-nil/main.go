package main

import "fmt"

// Go 语言的接口类型不是任意类型

type TestStruct struct{}

func NilOrNot(v interface{}) {
	if v == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("non-nil")
	}
}

func main(){
	var s *TestStruct
	// 发生隐式的类型转换，变量 nil 会被转换成 interface{} 类型，
	// interface{} 类型是一个结构体，它除了包含 nil 变量之外还包含变量的类型信息，也就是 TestStruct
	NilOrNot(s)   // non-nil
}
