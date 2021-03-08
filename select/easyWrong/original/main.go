package original

import (
	"fmt"
	"time"
)

func A() {
	outCh := make(chan string, 512)
	doneCh := make(chan struct{}, 1)

	B(outCh, doneCh)
	var loopDone bool
	for {
		select {
		case line, ok := <- outCh:
			fmt.Printf("before: get line: %s\n", line)
			if !ok {
				outCh = nil
				break
			}
			// 一些耗时操作
			fmt.Printf("get line: %s\n", line)
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