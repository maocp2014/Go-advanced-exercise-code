package main

import (
	"fmt"
	"time"
)

type test struct {
	stopCh chan struct{}
	outCh  chan int
	quitCh chan struct{}
}

func gen() chan int {
	outCh := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			outCh <- i
		}
		close(outCh)
	}()

	return outCh
}

func newTest() *test {
	return &test{
		stopCh: make(chan struct{}),
		outCh:  gen(),
		quitCh: make(chan struct{}),
	}
}

func (t *test) start() {
	go func() {
		for {
			select {
			case v, ok := <-t.stopCh:
				fmt.Println("receive stopCh: ", v, ok)
				t.quitCh <- struct{}{}
				return
			default:
				fmt.Println("xxx: ", <-t.outCh)
			}
		}
	}()
}

func main() {
	test := newTest()
	test.start()
	time.Sleep(time.Millisecond)
	close(test.stopCh)
	<-test.quitCh
}
