package main

import (
	"fmt"
	"sync"
)

// 针对资源泄露的第1种解决方法：buffer channel

/*
我们需要提供一种措施，即使当下游从上游接收数据时发生异常，上游也能成功退出。
一种方式是，把 channel 改为带缓冲的 channel，这样，它就可以承载指定数量的数据，
如果 buffer channel 还有空间，数据的发送将会立刻完成。

// 缓冲大小 2 buffer size 2
c := make(chan int, 2)
// 发送立刻成功 succeeds immediately
c <- 1
// 发送立刻成功 succeeds immediately
c <- 2
//blocks until another goroutine does <-c and receives 1
// 阻塞，直到另一个 goroutine 从 c 中接收数据
c <- 3

如果我们在创建 channel 时已经知道将发送的数据量，就可以把前面的代码简化一下。
比如，重写 gen 函数，将数据都发送至一个 buffer channel，这还能避免创建新的 goroutine。
func gen(nums ...int) <-chan int {
    out := make(chan int, len(nums))
    for _, n := range nums {
        out <- n
    }
    close(out)
    return out
}

译者按：channel 关闭后，不可再写入数据，否则会 panic，但是仍可读取已发送数据，而且可以一直读取 0 值。

继续往下游 stage，将又会返回到阻塞的 goroutine 中，我们也可以考虑给 merge 的输出 channel 加点缓冲。
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup

    // enough space for the unread inputs
    // 给未读的输入 channel 预留足够的空间
    out := make(chan int, 1)
    // ... the rest is unchanged ...

虽然通过这个方法，我们能解决了 goroutine 阻塞的问题，但是这并非一个优秀的设计。
比如 merge 中的 buffer 的大小 1 是基于我们已经知道了接下来接收数据的大小，以及下游将能消费的数量。
很明显，这种设计非常脆弱，如果上游多发送了一些数据，或下游并没接收那么多的数据，goroutine 将又会被阻塞。
因而，当下游不再准备接收上游的数据时，需要有一种方式，可以通知到上游。
*/

// gen sends the values in nums on the returned channel, then closes it.
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
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
	out := make(chan int, 1) // enough space for the unread inputs
	// ... the rest is unchanged ...

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
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 or 9
	return
	// The second value is sent into out's buffer, and all goroutines exit.
}
