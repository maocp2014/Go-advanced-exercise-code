package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case v, ok := <-ch:
		fmt.Printf("v: %v, ok: %v\n", v, ok)
	}
}

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
// 	d:/GoModuleWorkspace/channel/ok/demo_select3.go/main.go:8 +0x70
// exit status 2
