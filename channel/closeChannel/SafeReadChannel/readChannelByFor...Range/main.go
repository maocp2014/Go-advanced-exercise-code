package main

import "fmt"

func main() {
	c := make(chan int, 3)

	// 子协程写
	go func() {
		c <- 1
		close(c)
	}()

	// 直接读取通道，存在不知道子协程是否已关闭的情况
	// fmt.Println(<-c)  // 1
	// fmt.Println(<-c)  // 0  辨别不出是通道值0还是子协程已关闭，事实是子协程已关闭

	// 主协程读取：使用for...range安全的读取
	for value := range c {
		fmt.Println(value)  // 1
	}
}