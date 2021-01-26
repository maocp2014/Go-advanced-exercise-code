package main

import "fmt"

// 带缓冲的channel略有不同。尽管已经close了，但我们依旧可以从中读出关闭前写入的3个值。
// 第四次读取时，则会返回该channel类型的零值。向这类channel写入操作也会触发panic。

func main() {
	c := make(chan int, 3)
	c <- 15
	c <- 34
	c <- 65

	close(c)

	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)

	c <- 1
}


// 15
// 34
// 65
// 0
// panic: runtime error: send on closed channel
