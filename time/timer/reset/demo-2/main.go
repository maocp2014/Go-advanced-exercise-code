package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(5 * time.Second)
	fmt.Println("开始时间: ", time.Now().Format("2006-01-02 15:04:05"))

	go func() {
		count := 0
		for {
			<-timer1.C
			fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))

			count++

			fmt.Println("调用 Reset() 重新设置过期时间，将时间修改为 2s")

			timer1.Reset(2*time.Second)

			if count > 2 {
				fmt.Println("调用 Stop() 停止定时器")
				timer1.Stop()
			}
		}
	}()

	time.Sleep(15 * time.Second)
	fmt.Println("结束时间：", time.Now().Format("2006-01-02 15:04:05"))
}
