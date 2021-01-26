package main

import (
	"fmt"
)

// 生产者
func producer(ch chan int, quit chan bool, v int) {
	for i := 0; i < 3; i++ {
		v = v + 1
		ch <- v
		fmt.Println("write finish, v: ", v)
		// time.Sleep(time.Second)
	}
	// 生产者退出
	quit <- true
}

// 消费者
func customer(ch chan int, q1 chan bool, q2 chan bool) {
	for {
		select{
		case v := <- ch:
			fmt.Println("read finish, v: ", v)
		// 收到生产者退出的信号后，消费者发出退出的信号，并返回
		case <-q1:
			q2 <- true
			return
		}
	}
}

func main() {
    fmt.Println("task is start!")

	intChan := make(chan int)
	proQuitChan := make(chan bool)
	cusQuitChan := make(chan bool)

	value := 0

	go producer(intChan, proQuitChan, value)
	go customer(intChan, proQuitChan, cusQuitChan)

	// 收到消费者退出的信号，通知主协程退出
	<- cusQuitChan
	fmt.Println("task is done!")
}
