package main

import (
	"fmt"
	"time"
)

/*
最简单的协程池及周边由哪些组成：
1、协程池内的一定数量的协程。
2、任务队列，即jobCh，存在协程池不能立即处理任务的情况，所以需要队列把任务先暂存。
3、结果队列，即retCh，同上，协程池处理任务的结果，也存在不能被下游立刻提取的情况，要暂时保存。

协程池最简要（核心）的逻辑是所有协程从任务读取任务，处理后把结果存放到结果队列。
*/

func workerPool(n int, jobCh <-chan int, retCh chan<- string) {
	for i := 0; i < n; i++ {
		go worker(i, jobCh, retCh)
	}
}

func worker(id int, jobCh <-chan int, retCh chan<- string) {
	for job := range jobCh {
		ret := fmt.Sprintf("worker %d processed job: %d", id, job)
		retCh <- ret
	}
}

func genJob(n int) <-chan int {
	jobCh := make(chan int, 200)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()
	return jobCh
}

func main() {
	jobCh := genJob(10)
	retCh := make(chan string, 200)
	workerPool(5, jobCh, retCh)

	time.Sleep(time.Second)
	close(retCh)

	for ret := range retCh {
		fmt.Println(ret)
	}
}
