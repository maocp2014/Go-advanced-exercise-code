package main

import (
	"fmt"
	"time"
)

// 下面是一个常见的终结sub worker goroutines的方法，每个worker goroutine通过select监视
// 一个die channel来及时获取main goroutine的退出通知。

func worker(die chan bool, index int) {
	fmt.Println("Begin: This is Worker:", index)

	// flag := <-die
	// fmt.Println("flag = ", flag)  // false

	for {
		select {
		// case xx：
		// 做事的分支
		case <- die:
			fmt.Println("Done: This is Worker:", index)
			return
		}
	}
}

func main() {
	die := make(chan bool)

	for i := 1; i <= 100; i++ {
		go worker(die, i)
	}

	time.Sleep(time.Second * 5)

	// 通过close channel触发
	close(die)

	select {} // deadlock we expected
}
