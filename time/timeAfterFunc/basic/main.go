package main

import (
	"fmt"
	"time"
)

// time.AfterFunc() 用于延时执行自定义函数，会另起一个协程并等待指定时间过去，
// 然后调用函数 f。它返回一个 Timer，可以通过调用其 Stop() 方法来取消等待和对 f 的调用。

func main() {

	ch := make(chan int, 1)

	time.AfterFunc(10 * time.Second, func() {
		fmt.Println("10 seconds over...")
		fmt.Println("After func: ", time.Now())
		ch <- 8
	})

	for {
		select {
		case n := <-ch:
			fmt.Println("Ch time: ", time.Now())
			fmt.Println(n, "is arriving")
			fmt.Println("Done!")
			return
		default:
			fmt.Println("now1: ", time.Now())
			time.Sleep(3 * time.Second)
			fmt.Println("now2: ", time.Now())
			fmt.Println("time to wait")
		}
	}
}