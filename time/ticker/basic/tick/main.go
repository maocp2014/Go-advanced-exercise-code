package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now: ", time.Now())
	// 创建定时器，每隔 2 秒后，定时器就会给 channel 发送一个事件
	// func Tick(d Duration) <-chan Time {
	//	if d <= 0 {
	//		return nil
	//	}
	//	return NewTicker(d).C
	// }

	// Tick是对ticker的封装，只能访问channel,不能清除定时器等操作
	tick := time.Tick(2 * time.Second)

	for {
		curTime := <-tick
		fmt.Println(curTime)
	}
}