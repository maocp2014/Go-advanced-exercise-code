package main

import "fmt"

func worker(start chan bool, index int) {
	<-start
	fmt.Println("This is Worker:", index)
}

func main() {
	start := make(chan bool)
	for i := 1; i <= 100; i++ {
		go worker(start, i)
	}
	// close channel还可以用于协同多个Goroutines，我们创建了100个Worker Goroutine，
	// 这些Goroutine在被创建出来后都阻塞在"<-start"上，直到我们在main goroutine中给出开工的信号：
	// "close(start)"，这些goroutines才开始真正的并发运行起来。
	close(start)
	select {} // deadlock we expected
}
