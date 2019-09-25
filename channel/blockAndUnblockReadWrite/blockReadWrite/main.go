package main

import "fmt"

// 有缓冲通道的特点是，有缓冲时可以向通道中写入数据后直接返回，缓冲中有数据时可以
// 从通道中读到数据直接返回，这时有缓冲通道是不会阻塞的，它阻塞场景是：
// 1、通道的缓冲中无数据，但执行读通道。
// 2、通道的缓冲已经占满，向通道写数据，但无协程读。

// 无缓冲通道的特点是，发送的数据需要被读取后，发送才会完成，它阻塞场景：
// 1、通道中无数据，但执行读通道。
// 2、通道中无数据，向通道写数据，但无协程读取。

func main() {
	ReadNoDataFromBufCh()
	ReadNoDataFromNoBufCh()
	WriteBufChButFull()
	WriteNoBufCh()
}

// ReadNoDataFromBufCh 场景1
func ReadNoDataFromBufCh() {
	bufCh := make(chan int, 1)

	<-bufCh
	fmt.Println("read from no buffer channel success")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// ReadNoDataFromNoBufCh 场景1
func ReadNoDataFromNoBufCh() {
	noBufCh := make(chan int)

	<-noBufCh
	fmt.Println("read from no buffer channel success")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// WriteBufChButFull 场景2
func WriteBufChButFull() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100

	ch <- 1
	fmt.Println("write success no block")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// WriteNoBufCh 场景2
func WriteNoBufCh() {
	ch := make(chan int)

	ch <- 1
	fmt.Println("write success no block")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
