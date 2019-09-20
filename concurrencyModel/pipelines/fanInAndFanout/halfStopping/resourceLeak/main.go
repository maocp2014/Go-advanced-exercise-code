package main

import (
	"fmt"
	"sync"
)

/*
pipeline 中的函数包含一个模式:
      当数据发送完成，每个 stage 都应该关闭它们的输入 channel；
      只要输入 channel 没有关闭，每个 stage 就要持续从中接收数据；

我们可以通过编写 range loop 来保证所有 goroutine 是在所有数据都已经发送到下游的时候退出。
但在一个真实的场景下，每个 stage 都接收完 channel 中的所有数据，是不可能的。有时，我们的设计是：接收方只需要接收数据的部分子集即可。更常见的，如果 channel 在上游的 stage 出现了错误，那么，当前 stage 就应该提早退出。无论如何，接收方都不该再继续等待接收 channel 中的剩余数据，而且，此时上游应该停止生产数据，毕竟下游已经不需要了。
我们的例子中，即使 stage 没有成功消费完所有的数据，上游 stage 依然会尝试给下游发送数据，这将会导致程序永久阻塞。

// Consume the first value from the output.
    // 从 output 中接收了第一个数据
    out := merge(c1, c2)
    fmt.Println(<-out) // 4 or 9
    return
    // Since we didn't receive the second value from out,
    // one of the output goroutines is hung attempting to send it.
    // 我们并没有从 out channel 中接收第二个数据，
    // 所以上游的其中一个 goroutine 在尝试向下游发送数据时将会被挂起。
}
这是一种资源泄露，goroutine 是需要消耗内存和运行时资源的，goroutine 栈中的堆引用信息也是不会被 gc。
*/

// gen sends the values in nums on the returned channel, then closes it.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq receives values from in, squares them, and sends them on the returned
// channel, until in is closed.  Then sq closes the returned channel.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// merge receives values from each input channel and sends them on the returned
// channel.  merge closes the returned channel after all the input values have
// been sent.
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	// 从 output 中接收了第一个数据
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 or 9
	return
	// Since we didn't receive the second value from out,
	// one of the output goroutines is hung attempting to send it.
	// 我们并没有从 out channel 中接收第二个数据，
	// 所以上游的其中一个 goroutine 在尝试向下游发送数据时将会被挂起。
	// 这是一种资源泄露，goroutine 是需要消耗内存和运行时资源的，goroutine 栈中的堆引用信息也是不会被 gc。
}
