package main

import "fmt"

func main() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-quit:
				fmt.Println("B退出")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	fmt.Println("testAB退出")
}
