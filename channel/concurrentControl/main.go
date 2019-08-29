package main

import (
	"fmt"
	"time"
)

func worker(i int, limit chan bool, quit chan bool) {
	fmt.Println("do worker start: ", i)
	time.Sleep(time.Millisecond * 20)
	fmt.Println("do worker finish: ", i)

	<-limit

	if i == 9 {
		fmt.Println("完成任务")
		quit <- true
	}
}

// 通过channel控制goroutine数量，即控制并发数量
func main() {
	maxNum := 3
	limit := make(chan bool, maxNum)
	quit := make(chan bool)

	for i := 0; i < 10; i++ {
		fmt.Println("start worker: ", i)
		limit <- true

		go worker(i, limit, quit)
	}

	<-quit
	fmt.Println("收到退出通知，主程序退出")

}
