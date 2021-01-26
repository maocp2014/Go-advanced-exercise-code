package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	sign := make(chan int, 2)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			// time.Sleep(time.Second)
		}

		close(ch)

		fmt.Println("the channel is closed")

		sign <- 0

	}()

	go func() {
		for {
			i, ok := <-ch
			fmt.Printf("%d, %v \n", i, ok)

			if !ok {
				break
			}

			// time.Sleep(time.Second * 2)
		}

		sign <- 1

	}()

	<-sign
	<-sign
}
