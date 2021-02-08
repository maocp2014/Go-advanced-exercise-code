package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	go func() {
		// 假设我们在这里执行一个外部调用，2秒之后将结果写入 ch
		time.Sleep(time.Second * 2)
		ch <- "success"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}