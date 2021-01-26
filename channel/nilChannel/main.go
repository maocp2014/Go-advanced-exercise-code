package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0 ; i < 5 ; i++{
		go func() {
			fmt.Println("current i is ", i)
		}()
	}
	time.Sleep(time.Second)
}
