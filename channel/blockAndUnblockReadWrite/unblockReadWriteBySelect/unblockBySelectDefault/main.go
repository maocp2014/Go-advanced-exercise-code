package main

import (
	"errors"
	"fmt"
)

// select是执行选择操作的一个结构，它里面有一组case语句，它会执行其中无阻塞的那一个，
// 如果都阻塞了，那就等待其中一个不阻塞，进而继续执行，它有一个default语句，该语句是永远不会阻塞的，
// 我们可以借助它实现无阻塞的操作。

func main() {
	ReadNoDataFromNoBufChWithSelect()
	ReadNoDataFromBufChWithSelect()
	WriteNoBufChWithSelect()
	WriteBufChButFullWithSelect()
}

// 无缓冲通道读
func ReadNoDataFromNoBufChWithSelect() {
	bufCh := make(chan int)

	if v, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read: %d\n", v)
	}

	// Output:
	// channel has no data
}

// 有缓冲通道读
func ReadNoDataFromBufChWithSelect() {
	bufCh := make(chan int, 1)

	if v, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read: %d\n", v)
	}

	// Output:
	// channel has no data
}

// select结构实现通道读
func ReadWithSelect(ch chan int) (x int, err error) {
	select {
	case x = <-ch:
		return x, nil
	default:
		return 0, errors.New("channel has no data")
	}
}

// 无缓冲通道写
func WriteNoBufChWithSelect() {
	ch := make(chan int)
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}

	// Output:
	// channel blocked, can not write
}

// 有缓冲通道写
func WriteBufChButFullWithSelect() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}

	// Output:
	// channel blocked, can not write
}

// select结构实现通道写
func WriteChWithSelect(ch chan int) error {
	select {
	case ch <- 1:
		return nil
	default:
		return errors.New("channel blocked, can not write")
	}
}
