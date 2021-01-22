package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
	close(c)
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	a := cap(c)
	b := len(c)
	c <- 2
	c <- 3
	c <- 4 // block



	fmt.Println("channel c's capacity: ", a)
	fmt.Println("channel c's length: ", b)



	fmt.Println("main() stopped")
}
