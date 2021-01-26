package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// quit := make(chan bool)
	quit := make(chan struct{})

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := func() {
				println(id, time.Now().Nanosecond())
				time.Sleep(time.Second)
			}

			for {
				select {
				case v, ok := <-quit: // closed channel 不会阻塞，因此可⽤作退出通知。
					fmt.Println("xxx: ", v, ok)
					return
				default: // 执⾏正常任务。
					task()
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 5) // 让测试 goroutine 运⾏⼀会。
	close(quit)                 // 发出退出通知。
	wg.Wait()
}
