package main

import (
	"errors"
	"fmt"
	"time"
)

// 使用default实现的无阻塞通道阻塞有一个缺陷：当通道不可读或写的时候，会即可返回。
// 实际场景更多的需求是:我们希望尝试读一会数据，或者尝试写一会数据，如果实在没法读写再返回，
// 程序继续做其它的事情。

// 使用定时器替代default可以解决这个问题，给通道增加读写数据的容忍时间，如果规定时间内无法读写，就即刻返回。

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
	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case x = <-ch:
		return x, nil
	case <-timeout.C:
		return 0, errors.New("read time out")
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
	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case ch <- 1:
		return nil
	case <-timeout.C:
		return errors.New("write time out")
	}
}
