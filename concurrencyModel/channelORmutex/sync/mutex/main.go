package main

import (
	"fmt"
	"time"
)

func main() {
	b := NewBank()
	go b.Deposit("xiaoming", 100)
	go b.Withdraw("xiaoming", 20)
	go b.Deposit("xiaogang", 2000)

	time.Sleep(time.Second)
	fmt.Printf("xiaoming has: %d\n", b.Query("xiaoming"))
	fmt.Printf("xiaogang has: %d\n", b.Query("xiaogang"))
}
