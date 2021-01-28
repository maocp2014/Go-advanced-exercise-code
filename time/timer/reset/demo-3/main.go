package main

import (
	"fmt"
	"time"
)

func main() {

	// timer 过期
	timer := time.NewTimer(1 * time.Second)
	time.Sleep(2 * time.Second)
	ret := timer.Reset(2 * time.Second)
	fmt.Println(ret)

	// timer 停止
	timer = time.NewTimer(1 * time.Second)
	timer.Stop()
	ret = timer.Reset(1 * time.Second)
	fmt.Println(ret)
}
