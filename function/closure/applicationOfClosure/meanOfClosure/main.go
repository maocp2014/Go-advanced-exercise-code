package main

// 闭包的主要意义:
// 缩小变量作用域，减少对全局变量的污染。下面的累加如果用全局变量进行实现，全局变量容易被其他人污染。
// 同时，如果要实现n个累加器，那么每次需要n个全局变量。利用闭包，每个生成的累加器
// myAdder1, myAdder2 := adder(), adder()有自己独立的sum，sum可以看作为myAdder1.sum与myAdder2.sum。

// 利用闭包可以实现有自身状态的函数！

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	myAdder := adder()

	// 从1加到10
	for i := 1; i <= 10; i++ {
		myAdder(i)
	}

	fmt.Println(myAdder(0)) // 55
	// 再加上45
	fmt.Println(myAdder(45)) // 100
}
