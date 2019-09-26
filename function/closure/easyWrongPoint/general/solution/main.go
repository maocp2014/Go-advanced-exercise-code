package main

import (
	"fmt"
)

// B1 解决方法1
func B1() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		b[i] = func(j int) func() {
			return func() {
				fmt.Println(j)
			}
		}(i) // 直接传参
	}
	return b
}

// B2 解决方法2
func B2() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		j := i // 中间变量暂存i的值
		b[i] = func() {
			fmt.Println(j)
		}
	}
	return b
}

func main() {
	c := B1()
	// c := B2()
	c[0]()
	c[1]()
	c[2]()
}
