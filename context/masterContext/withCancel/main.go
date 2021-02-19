package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// type Context interface {
//    Deadline() (deadline time.Time, ok bool)
//    Done() <-chan struct{}
//    Err() error
//    Value(key interface{}) interface{}
// }

// 一个空的 Context 没有携带任何值、没有截止时间并且不能取消。context.Background() 函数返回
// 默认的空 Context，它通常用来派生其他 Context，还可以用在测试用例中。
//
// 由于 Context 是一个接口，所以它的值也可以是 nil，但是不推荐传递 nil Context，
// context.Background() 是个不错的选择。
// context.TODO() 也会返回一个空的 Context，如果我们还不清楚使用什么 Context 时，可以用它来占个位。

// channel to send square of integers
var c = make(chan int)

// send square of numbers
func square(ctx context.Context) {
	i := 0

	for {
		select {
		// 当调用 cancel 函数、Done channel 关闭时；当 c channel 能写入时会执行第二个 case。
		case <-ctx.Done():
			return // kill goroutine
		case c <- i * i:
			i++
		}
	}
}

// main goroutine
func main() {

	// create cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	go square(ctx) // start square goroutine

	// get 5 square
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	// 通常，大多数人习惯使用 defer cancel() 来取消 Context，然而取消 Context
	// 能够释放其相关的系统资源，因此建议确定不再需要使用 Context 并且子协程应立即
	// 停止工作时立即调用 cancel 函数。

	// 调用了 cancel 函数，将会关闭 ctx 的 Done channel

	// cancel context
	cancel() // instead of `defer context()`

	// do other job
	time.Sleep(3 * time.Second)

	// print active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}

// 一个 Context 可以传递给多个协程，调用 cancel 函数之后再次调用不会有任何作用。
//
// Context 仅可以作为参数传递给函数或者 goroutine，并且最好命名为 ctx，不推荐将 Context 放在结构体中

