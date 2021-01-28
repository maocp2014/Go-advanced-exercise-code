package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now: ", time.Now())
	ticker1 := time.NewTicker(2*time.Second)
	for {
		curTime := <-ticker1.C
		fmt.Println(curTime)
	}
}