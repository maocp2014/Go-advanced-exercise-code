package main

import (
	"fmt"
	"time"
)

func A() {
	outCh := make(chan string)    // 修改成无缓冲通道
	doneCh := make(chan struct{}) // 修改成无缓冲通道

	i := 1

	B(outCh, doneCh)

	var loopDone bool

	for {
		select {
		case line, ok := <- outCh:
			if !ok {
				outCh = nil
				break
			}
			// 一些耗时操作
			fmt.Printf("get %d line: %s\n", i, line)
			i++
		case dCh, ok := <- doneCh:
			fmt.Printf("dch: %v\n", dCh)
			fmt.Printf("ok: %v\n", ok)
			if !ok {
				doneCh = nil
				break
			}
			loopDone = true
		default:
		}

		if (outCh == nil && doneCh == nil) || loopDone {
			break
		}
	}
}

func B(out chan<- string, done chan<- struct{}) {
	go func(){
		for i:=0; i < 20; i++ {
			out <- "hello world"
		}
		done <- struct{}{}
	}()
}

func main() {
	A()   // 不能全量打印出“hello world”
	time.Sleep(2)
}