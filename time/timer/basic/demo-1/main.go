package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now time: ",time.Now().Format("2006-01-02 15:04:05"))

	// 传参是告诉 Timer 需要等待多长时间，Timer 是带有一个缓冲 channel 的 struct，
	// 在定时时间到达之前，没有数据写入 Timer.C，读取操作会阻塞当前协程，
	// 到达定时时间时，会向 channel 写入数据（当前时间），阻塞解除，被阻塞的协程得以恢复运行，
	// 达到延时或者定时执行的目的。
	timer := time.NewTimer(3 * time.Second)
	<-timer.C

	fmt.Println("now time: ",time.Now().Format("2006-01-02 15:04:05"))
}
