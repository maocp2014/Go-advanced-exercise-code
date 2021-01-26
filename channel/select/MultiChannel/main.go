package main

import (
	"fmt"
	"time"
)

// 有时候还会遇到多个生产者，只要有一个生产者就绪，消费者就可以进行消费的情况。
// 这个时候可以使用go语言提供的select语句，它可以同时管理多个通道读写，如果所有通道都不能读写，
// 它就整体阻塞，只要有一个通道可以读写，它就会继续。

func main() {

	var ch1 = make(chan int)
	var ch2 = make(chan int)

	fmt.Println(time.Now().Format("15:04:05"))

	go func(ch chan int) {
		time.Sleep(time.Millisecond)
		ch <- 1
	}(ch1)

	go func(ch chan int) {
		time.Sleep(time.Millisecond * 2)
		ch <- 2
	}(ch2)

	for {
		select {
		case v := <-ch1:
			fmt.Println(time.Now().Format("15:04:05") + ":来自ch1:", v)
		case v := <-ch2:
			fmt.Println(time.Now().Format("15:04:05") + ":来自ch2:", v)
		// default:
		// 	fmt.Println("channel is empty !")
		}
	}
}

// 默认select处于阻塞状态，1s后，子协程1完成写入，主协程读出了数据；接着子协程2完成写入，主协程读出了数据；接着主协程挂掉了，
// 原因是主协程发现在等一个永远不会来的数据，这显然是没有结果的，干脆就直接退出了。
//
// 如果把注释的部分打开，那么程序在打印出来自ch1、ch2的数据后，就会一直执行default里面的程序。这个时候程序不会退出。
// 原因是当 select 语句所有通道都不可读写时，如果定义了 default 分支，那就会执行 default 分支逻辑。
//
// select{}代码块是一个没有任何case的select，它会一直阻塞。