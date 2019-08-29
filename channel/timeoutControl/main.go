package main

import (
	"fmt"
	"time"
)

// 具名函数版本，注意需要传参，但是在匿名函数中不需要传参
/*
package main

import (
	"fmt"
	"time"
)

func f1(g chan int, quit chan bool) {
	for {
		select {
		case v := <-g:
			fmt.Println("v: ", v)
		case <-time.After(time.Second * time.Duration(3)):
			quit <- true
			fmt.Println("timeout, 通知主协程退出！")
			return
		}
	}
}

func main() {
	g := make(chan int)
	quit := make(chan bool)

	go f1(g, quit)

	for i := 0; i < 3; i++ {
		g <- i
	}
*/

func main() {
	g := make(chan int)
	quit := make(chan bool)

	// 匿名函数，不需要传参，直接共享主协程变量
	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println("v: ", v)
			case <-time.After(time.Second * time.Duration(3)):
				quit <- true
				fmt.Println("timeout, 通知主协程退出！")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}

	<-quit

	fmt.Println("收到退出通知，主协程退出！")
}
