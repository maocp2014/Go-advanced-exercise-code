package main

func main() {
	ch := make(chan struct{})
	close(ch)
	ch <- struct{}{}
}

// panic: send on closed channel
//
// goroutine 1 [running]:
// main.main()