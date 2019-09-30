package main

import "fmt"

func main() {
	message := make(chan string)
	count := 3

	go func() {
		 for i := 1; i <= count; i++ {
			  message <- fmt.Sprintf("message %d", i)
		 }
		 // 关闭channel，否则for...range在读取channel时会deadlock，这是因为for...range会一直读取channel
		 // 的值，直到channel关闭或channel的缓存为0
		 close(message)
	}()

	for msg := range message {
		 fmt.Println(msg)
	}
}