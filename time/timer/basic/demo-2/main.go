package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now time: ", time.Now().Format("2006-01-02 15:04:05"))

	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		// 超时时间还没到时，协程会阻塞；超时时间到了之后会返回当前时间
		curTime := <-timer1.C
		// 定时时间到了之后执行打印操作，实际开发时可以实现具体的业务逻辑
		fmt.Println("Timer 1 fired, now time: ", curTime.Format("2006-01-02 15:04:05"))
	}()


	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// 取消定时器
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stop")
	}

	time.Sleep(5 * time.Second)
}