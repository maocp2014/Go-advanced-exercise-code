package main

import "fmt"
// func main() {
// 	done := make(chan bool)

// 	go func() {
// 		 println("goroutine message")

// 		 // We are only interested in the fact of sending itself, 
// 		 // but not in data being sent.
// 		 done <- true
// 	}()

// 	println("main function message")
// 	<-done 
// } 


func main() {
	// Data is irrelevant
	done := make(chan struct{})

	go func() {
		 fmt.Println("goroutine message")

		 // Just send a signal "I'm done"
		 close(done)
	}()

	fmt.Println("main function message")
	fmt.Println(<-done) // {}，这是因为channel关闭后，会读取到channel类型的默认零值，这里是{}
} 