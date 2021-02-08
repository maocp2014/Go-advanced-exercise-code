package ticker

import (
	"fmt"
	"time"
)

// Ticker 定时器原理与 Timer 类似，Ticker 是一个周期触发定时的定时器，按给定时间间隔往 channel 发送系统当前时间，
// 而 channel 的接收者可以以固定的时间间隔从 channel 中读取。

func main() {
	fmt.Println("now: ", time.Now())
	// 创建定时器，每隔 2 秒后，定时器就会给 channel 发送一个事件
	ticker1 := time.NewTicker(2*time.Second)

	for {
		curTime := <-ticker1.C
		fmt.Println(curTime)
	}
}