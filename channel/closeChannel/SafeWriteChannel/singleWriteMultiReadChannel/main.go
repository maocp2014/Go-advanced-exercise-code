package main

import "fmt"

// Go 语言并不存在一个内置函数可以判断出通道是否已经被关闭。确保通道写安全的最好方式是由
// 负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。但是这个方法只能解决单写多读的场景。

func main() {
	c := make(chan int, 3)

	// 子协程写
	go func() {
		c <- 1
		close(c)   // 写协程关闭通道，这是单协程写通道的安全方式
	}()

	// 直接读取通道，存在不知道子协程是否已关闭的情况
	// fmt.Println(<-c)  // 1
	// fmt.Println(<-c)  // 0  辨别不出是通道值0还是子协程已关闭，事实是子协程已关闭

	// 主协程读取：使用for...range安全的读取
	for value := range c {
		fmt.Println(value)  // 1
	}
}