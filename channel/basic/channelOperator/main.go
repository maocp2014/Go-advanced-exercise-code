package main

import (
	"fmt"
)

func greet(c chan string) {
	fmt.Println("c1 " + <-c)
	fmt.Println("c2 " + <-c)
}

func main() {
	// var c chan int
	// fmt.Println(c)

	// c := make(chan int)

	// fmt.Printf("type of `c` is %T\n", c)  // type of `c` is chan int
	// fmt.Printf("value of `c` is %v\n", c) // value of `c` is 0xc000050060，默认情况下，channel 是指针

	fmt.Println("main() started")

	c := make(chan string, 1)

	go greet(c)

	c <- "John"

	close(c)

	c <- "Mike"

	fmt.Println("main() stopped")
}
