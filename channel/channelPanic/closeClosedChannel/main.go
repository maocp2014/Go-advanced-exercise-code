package main

func main() {
	ch := make(chan struct{})
	close(ch)
	close(ch)
}

// panic: close of closed channel
//
// goroutine 1 [running]:
// main.main()