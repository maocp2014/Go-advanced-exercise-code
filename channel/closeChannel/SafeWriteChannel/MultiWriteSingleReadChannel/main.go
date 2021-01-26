package main

import (
	"fmt"
	"sync"
	"time"
)

// 如果遇到多写单读的情况就有问题了：无法知道其它写协程什么时候写完，
// 那么也就不能确定什么时候关闭通道。这个时候就得额外使用一个通道专门做这个事情。
// 我们可以使用内置的 sync.WaitGroup，它使用计数来等待指定事件完成

func main() {
	var ch = make(chan int, 8)

	// 写协程
	var wg = new(sync.WaitGroup)

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(num int, ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- num
			ch <- num * 10
		}(i, ch, wg)
	}

	// 读
	go func(ch chan int) {
		for num := range ch {
			fmt.Println(num)
		}
	}(ch)

	// Wait阻塞等待所有的写通道协程结束,待计数值变成零，Wait才会返回
	wg.Wait()

	// 安全的关闭通道
	close(ch)

	// 防止读取通道的协程还没有完毕
	time.Sleep(time.Second)

	fmt.Println("finish")
}
