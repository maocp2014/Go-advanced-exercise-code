package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now time: ", time.Now().Format("2006-01-02 15:04:05"))

	timer1 := time.NewTimer(5 * time.Second)
	// 重新设置为 2s
	ok := timer1.Reset(2 * time.Second)
	fmt.Println("ok: ", ok)
	curTime := <-timer1.C

	fmt.Println("now time: ", curTime.Format("2006-01-02 15:04:05"))
}
