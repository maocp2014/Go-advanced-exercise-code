package main

import "fmt"

// select{}永远阻塞
// select{}的效果等价于创建了1个通道，直接从通道读数据：

// ch := make(chan int)
// <-ch
// 但是，这个写起来多麻烦啊！没select{}简洁啊。 但是，永远阻塞能有什么用呢！？
// 当你开发一个并发程序的时候，main函数千万不能在子协程干完活前退出啊，
// 不然所有的协程都被迫退出了，还怎么提供服务呢？
// 比如，写了个Web服务程序，端口监听、后端处理等等都在子协程跑起来了，main函数这时候能退出吗？

func main() {
	fmt.Println("start")
	select {}
	fmt.Println("quit")
}

// start
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [select (no cases)]:
// main.main()
// 	d:/GoModuleWorkspace/channel/select/advanced/blockBySelect{}/main.go:17 +0x82
// exit status 2
