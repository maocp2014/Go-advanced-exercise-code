package main

import (
	"fmt"
	"time"
)

func main() {
	input1 := make(chan int)
	input2 := make(chan int)
	output := make(chan int)

	// golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
	// select 的代码形式和 switch 非常相似， 不过 select 的 case 里的操作语句只能是"IO操作"
	// （不仅仅是取值<-channel，赋值channel<-也可以）， select 会一直等待等到某个 case 语句完成，
	// 也就是等到成功从channel中读到数据. 则 select 语句结束
	go func(in1, in2 <-chan int, out chan<- int) {
		for {
			select {
			case v := <-in1:
				out <- v
			case v := <-in2:
				out <- v
			}
		}
	}(input1, input2, output)

	go func() {
		for i := 0; i < 10; i++ {
			input1 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 20; i < 30; i++ {
			input2 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			select {
			case value := <-output:
				fmt.Println("输出：", value)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("主线程退出")
}
