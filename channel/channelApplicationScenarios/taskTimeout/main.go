package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-doWork():
		fmt.Println("任务处理完成")
	case <-time.After(3 * time.Second):
		fmt.Println("任务处理超时")
	}
}

func doWork() <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		// 任务处理耗时
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()

	return ch
}