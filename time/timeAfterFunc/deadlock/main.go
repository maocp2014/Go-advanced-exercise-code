package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.AfterFunc(8*time.Second, func() {
		fmt.Println("Golang来啦")
	})

	fmt.Println("t.C: ", t.C)  // nil channel

	// fatal error: all goroutines are asleep - deadlock!
	for {
		select {
		case <-t.C:
			fmt.Println("seekload")
			break
		}
	}
}