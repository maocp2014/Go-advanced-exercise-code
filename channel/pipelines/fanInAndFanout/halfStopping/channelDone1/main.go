package main

import (
	"fmt"
	"sync"
)

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
/*
发送方 merge 用 select 语句替换了之前的发送操作，它负责通过 out channel 发送数据或者从 done 接收数据。
done 接收的值是没有实际意义的，只是表示 out 应该停止继续发送数据了，用空 struct 即可。
output 函数将会不停循环，因为上游，即 sq ，并没有阻塞。我们过会再讨论如何退出这个循环。
*/
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done: // HL
			}
		}
		wg.Done()
	}
	// ... the rest is unchanged ...

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

/*
如果 main 函数在没把 out 中所有数据接收完就退出，它必须要通知上游停止继续发送数据。如何做到？
我们可以在上下游之间引入一个新的 channel，通常称为 done。
示例中有两个可能阻塞的 goroutine，所以， done 需要发送两个值来通知它们。
*/

/*
这种方法有个问题，下游只有知道了上游可能阻塞的 goroutine 数量，才能向每个 goroutine 都发送了一个 done 信号，
从而确保它们都能成功退出。但多维护一个 count 是很令人讨厌的，而且很容易出错。
*/
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2) // HL
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{} // HL
	done <- struct{}{} // HL
}
