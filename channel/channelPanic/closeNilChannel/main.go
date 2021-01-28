package main

func main() {
	var ch chan struct{}
	close(ch)
}

// panic: close of nil channel
//
// goroutine 1 [running]:
// main.main()