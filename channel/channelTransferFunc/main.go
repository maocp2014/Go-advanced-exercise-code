package main

import (
	"fmt"
	"time"
)

type Operation func(a, b int) int

func main() {
	funcChan := make(chan Operation)

	f := func(a, b int) int {
		return a + b
	}

	go func(i, j int) {
		s := (<-funcChan)(i, j)
		fmt.Println(s)
	}(1, 2)

	funcChan <- f

	time.Sleep(time.Second)
}
