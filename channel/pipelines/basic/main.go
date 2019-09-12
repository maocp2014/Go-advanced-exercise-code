package main

import "fmt"

// 原始的流水线模型
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
// main 函数负责创建 pipeline 并执行最后一个 stage 的任务。
// 它将从第二个 stage 接收数据，并将它们打印出来，直到 channel 关闭。
func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}
*/

// 既然，sq 的输入和输出的 channel 类型相同，那么我们就可以把它进行组合，
// 从而形成多个 stage。比如，我们可以把 main 函数重写为如下的形式：
func main() {
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
