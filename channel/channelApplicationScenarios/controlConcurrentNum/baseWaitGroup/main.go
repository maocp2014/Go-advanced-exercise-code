package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	jobCount := 100
	limit := 10

	for i := 0; i <= jobCount; i += limit {
		for j := 0; j < i; j++ {

			wg.Add(1)

			go func(item int) {
				defer wg.Done()
				job(item)
			}(j)
		}

		wg.Wait()
	}
}

func job(index int) {
	// 耗时任务
	time.Sleep(1 * time.Second)
	fmt.Printf("任务:%d已完成\n", index)
}
