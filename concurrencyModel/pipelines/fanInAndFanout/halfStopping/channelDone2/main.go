package main

import (
	"fmt"
	"sync"
)

/*
我们需要一种方式，可以告诉上游的所有 goroutine 停止向下游继续发送信息。在 Go 中，其实可通过关闭 channel 实现，
因为在一个已关闭的 channel 接收数据会立刻返回，并且会得到一个零值。
这也就意味着，main 仅需通过关闭 done channel，就可以让所有的发送方解除阻塞。关闭操作相当于一个广播信号。
为确保任意返回路径下都成功调用，我们可以通过 defer 语句关闭 done。
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
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c or done is closed, then calls
	// wg.Done.

	// 为每个输入 channel 启动一个 goroutine，将输入 channel 中的数据拷贝到
	// out channel 中，直到输入 channel，即 c，或 done 关闭。
	// 接着，退出循环并执行 wg.Done()
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done: // HL
			}
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
	done := make(chan struct{}) // HL
	out := merge(done, c1, c2)  // HL
	fmt.Println(<-out)          // 4 or 9

	// Tell the remaining senders we're leaving.
	close(done) // HL
}
