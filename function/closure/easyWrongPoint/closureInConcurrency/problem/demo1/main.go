package main

import (
	"fmt"
	"time"
)

// 闭包的副作用
func test1() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range sl {
		go func() {
			fmt.Printf("%d %d\n", i, v)
		}()
	}

	time.Sleep(time.Second)
}

// 正确处理方式
func test2() {

	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range sl {
		// 只使用匿名函数进行传值，不使用闭包
		go func(a, b int) {
			fmt.Printf("%d %d\n", a, b)
		}(i, v)
	}
}

func main() {
	test1()
	time.Sleep(time.Second)
	fmt.Println("--------------")
	test2()
	time.Sleep(time.Second)
}
