package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	c := make(chan struct{}, 1)
	go func() {
		for i := 0; i < 10000; i++ {
			c <- struct{}{}
			x++
			fmt.Println("goroutine A, x = ", x)
			<-c
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			c <- struct{}{}
			x--
			fmt.Println("goroutine B, x = ", x)
			<-c
		}
	}()

	time.Sleep(5*time.Second)
	fmt.Println("x should be 0, and x =", x)
}