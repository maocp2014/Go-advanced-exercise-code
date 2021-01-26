package main

import "fmt"

// newUniqueIDService通过一个channel与main goroutine关联，main goroutine无需知道
// uniqueid实现的细节以及当前状态，只需通过channel获得最新id即可。

func newUniqueIDService() <-chan string {
	id := make(chan string)

	go func() {
		var counter int64 = 0
		for {
			id <- fmt.Sprintf("%x", counter)
			counter += 1
		}
	}()

	return id
}

func main() {
	id := newUniqueIDService()

	for i := 0; i < 10; i++ {
		fmt.Println(<-id)
	}
}
