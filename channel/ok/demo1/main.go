package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	v, ok := <-ch
	fmt.Printf("v: %v, ok: %v\n", v, ok)
}

// fatal error: all goroutines are asleep - deadlock!
//

// goroutine 1 [chan receive]:
// main.main()
// 	d:/GoModuleWorkspace/channel/ok/demo1/main.go:7 +0x67
// exit status 2
