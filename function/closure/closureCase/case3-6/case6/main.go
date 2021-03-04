package main

import (
	"fmt"
	"time"
)

var foo6Chan = make(chan int, 10)

func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}

// 不知道各位看官是否好奇，既然说Go Routine执行的时候比for-loop慢，那如果我在遍历的时候增加sleep机制呢？
func main() {
	// Q3第一组实验
	go foo6()   // 这里要使用协程，否则死锁
	// foo6Chan <- 1
	// foo6Chan <- 2
	// foo6Chan <- 3
	// foo6Chan <- 5
	// 绝大部分时候执行出来都是5。

	// Q3第二组实验
	// foo6Chan <- 11
	// time.Sleep(time.Duration(1) * time.Nanosecond)
	// foo6Chan <- 12
	// time.Sleep(time.Duration(1) * time.Nanosecond)
	// foo6Chan <- 13
	// time.Sleep(time.Duration(1) * time.Nanosecond)
	// foo6Chan <- 15

	// Q3第三组实验
	// 微秒
	foo6Chan <- 21
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 22
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 23
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 25
	time.Sleep(time.Duration(10) * time.Second)
	// 毫秒
	// foo6Chan <- 31
	// time.Sleep(time.Duration(1) * time.Millisecond)
	// foo6Chan <- 32
	// time.Sleep(time.Duration(1) * time.Millisecond)
	// foo6Chan <- 33
	// time.Sleep(time.Duration(1) * time.Millisecond)
	// foo6Chan <- 35
	// time.Sleep(time.Duration(10) * time.Second)
	// 秒
	// foo6Chan <- 41
	// time.Sleep(time.Duration(1) * time.Second)
	// foo6Chan <- 42
	// time.Sleep(time.Duration(1) * time.Second)
	// foo6Chan <- 43
	// time.Sleep(time.Duration(1) * time.Second)
	// foo6Chan <- 45
	time.Sleep(time.Duration(10) * time.Second)
	// 实验完毕，最后记得关闭channel
	close(foo6Chan)
}