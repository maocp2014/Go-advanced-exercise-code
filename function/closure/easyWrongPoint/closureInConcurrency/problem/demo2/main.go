package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Go语言的并发时，一定要处理好循环中的闭包引用的外部变量

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()

		// time.Sleep(1 * time.Second) // 设置时间延时1秒，验证下面说法
	}
	wg.Wait()
}

// output:
// 5
// 5
// 5
// 5
// 4

// 这种现象的原因在于闭包共享外部的变量i，注意到，每次调用go就会启动一个goroutine，这需要一定时间；
// 但是，启动的goroutine与循环变量递增不是在同一个goroutine，可以把i认为处于主goroutine中。
// 启动一个goroutine的速度远小于循环执行的速度，所以即使是第一个goroutine刚启动时，
// 外层的循环也执行到了最后一步了。由于所有的goroutine共享i，而且这个i会在最后一个使用
// 它的goroutine结束后被销毁，所以最后的输出结果都是最后一步的i==5(这里不一定，看电脑配置，比如我跑出的结果如上)。

// 我们可以使用循环的延时在验证上述说法：sleep注释部分
// output:
// 0
// 1
// 2
// 3
// 4
// 每一步循环至少间隔一秒，而这一秒的时间足够启动一个goroutine了，因此这样可以输出正确的结果。
