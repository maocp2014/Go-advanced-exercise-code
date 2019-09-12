package main

import (
	"fmt"
	"sync"
)

// 一个涉及三个 stage 的 pipeline。

// 第一个 stage，gen 函数。它负责将把从参数中拿到的一系列整数发送给指定 channel。
// 它启动了一个 goroutine 来发送数据，当数据全部发送结束，channel 会被关闭。
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

// 第二个 stage，sq 函数。它负责从输入 channel 中接收数据，并会返回一个新的 channel，即输出 channel，
// 它负责将经过平方处理过的数据传输给下游。当输入 channel 关闭，并且所有数据都已发送到下游，
// 就可以关闭这个输出 channel 了。
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

/*
merge 函数负责将从一系列输入 channel 中接收的数据合并到一个 channel 中。
它为每个输入 channel 都启动了一个 goroutine，并将它们中接收到的值发送到惟一的输出 channel 中。
在所有的 goroutines 启动后，还会再另外启动一个 goroutine，它的作用是，当所有的输入 channel 关闭后，
负责关闭唯一的输出 channel 。在已关闭的 channel 发送数据将导致 panic，因此要保证在关闭 channel 前，
所有数据都发送完成，是非常重要的。sync.WaitGroup 提供了一种非常简单的方式来完成这样的同步。
*/
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	// 为每个输入 channel 启动一个 goroutine
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
	// 启动一个 goroutine 负责在所有的输入 channel 关闭后，关闭这个唯一的输出 channel
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// 当多个goroutine从一个 channel 中读取数据，直到 channel 关闭，这称为扇出 fan-out。
	// 利用它，我们可以实现了一种分布式的工作方式，通过一组 workers 实现并行的 CPU 和 IO。

	// Distribute the sq work across two goroutines that both read from in.
	// 分布式处理来自 in channel 的数据
	c1 := sq(in)
	c2 := sq(in)

	// 当一个goroutine从多个 channel 中读取数据，直到所有 channel 关闭，这称为扇入 fan-in。
	// 扇入是通过将多个输入 channel 的数据合并到同一个输出 channel 实现的，
	// 当所有的输入 channel 关闭，输出的 channel 也将关闭。

	// merge，负责 fan-in 处理结果
	// Consume the merged output from c1 and c2.
	// 从 channel c1 和 c2 的合并后的 channel 中接收数据
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}