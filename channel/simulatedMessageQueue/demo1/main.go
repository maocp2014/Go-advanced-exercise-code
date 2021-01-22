package main

import "fmt"

func main() {
	fmt.Println("run in main coroutine.")

	count := 10
	c := make(chan bool, count)

	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Printf("run in child coroutine %d.\n", i)
			c <- true
		}(i)
	}

    for i := 0; i < count; i++ {
    	<- c
	}
}