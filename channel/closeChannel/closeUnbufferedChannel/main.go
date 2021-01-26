package main

import "fmt"

// 可以看到在一个已经close的unbuffered channel上执行读操作，回返回channel对应类型的零值，
// 比如bool型channel返回false，int型channel返回0。但向close的channel写则会触发panic。不过无论读写都不会导致阻塞。

func main() {
	cb := make(chan bool)
	close(cb)
	x := <-cb
	fmt.Printf("%#v\n", x)

	x, ok := <-cb
	fmt.Printf("%#v %#v\n", x, ok)

	ci := make(chan int)
	close(ci)
	y := <-ci
	fmt.Printf("%#v\n", y)

	cb <- true
}

// false
// false false
// 0
// panic: send on closed channel
//
// goroutine 1 [running]:
// main.main()