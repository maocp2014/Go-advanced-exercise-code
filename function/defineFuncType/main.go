package main

import "fmt"

// Operation 使用type自定义的一个函数类型，type后是类型名称，本例中是Operation，
// 再后面是类型的定义，对于函数而言，被称为signature，即函数签名，
// 这个函数签名表示：Operation类型的函数，它以2个int类型为入参，以1个int为返回值。
// 所有满足该函数签名的函数，都是Operation类型的函数。
type Operation func(a, b int) int

// Add Operation类型函数
func Add(a, b int) int {
	return a + b
}

// Sub Operation类型函数
func Sub(a, b int) int {
	return a - b
}

func main() {
	// 定义Operation类型的变量
	var op Operation
	// 函数Add和Sub都符合Operation的签名，所以Add和Sub都是Operation类型
	// 变量op是Operation类型的，即自定义的函数类型变量，可以把Add作为值赋值给变量op，执行op等价于执行Add
	op = Add
	fmt.Println(op(1, 2)) // 3

	op = Sub
	fmt.Println(op(2, 1)) // 1

	// var op1, op2 Operation
	// op1 = Add
	// fmt.Println(op1(1, 2)) // 3

	// op2 = Sub
	// fmt.Println(op2(2, 1)) // 1
}
