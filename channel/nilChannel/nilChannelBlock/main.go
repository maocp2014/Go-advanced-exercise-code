package main

// 对一个没有初始化的channel进行读写操作都将发生阻塞

func main() {
	var c chan int
	// <-c
	c <- 1
}

// fatal error: all goroutines are asleep – deadlock!