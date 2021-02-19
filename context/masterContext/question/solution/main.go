package main

import (
	"fmt"
	"runtime"
	"time"
)

var signal = make(chan bool)

// 由于 printNumbers 协程的 select 语句，至少在 1 秒钟时间内，从 signal channel 是读取不到任何值的，
// 便会一直执行 default case。一旦 channel 有值，便会执行 return 语句终止 for 循环，printNumbers
// 协程退出。

func printNumbers() {
	counter := 1

	for {
		select {
		case <-signal:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			counter++
		}
	}
}

func main() {
	go printNumbers()

	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
	time.Sleep(time.Second)

	signal <- true

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")
}