package main

import(
	"fmt"
	"time"
)

// 输入型channel
func f1(inChan chan<- int, quit chan<- bool){
	for i := 0; i < 10; i++{
		inChan <- i
		time.Sleep(time.Millisecond * 500)
	}
	quit <- true
	quit <- true
}

// 输出型channel
func f2(outChan <-chan int, quit <-chan bool){
	for {
		select {
			case v := <-outChan:
				fmt.Println("output value: ", v)

			case <- quit:
				fmt.Println("收到退出通知，退出")
				return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	
	go f1(ch, quit)
	go f2(ch, quit)

	<- quit

	fmt.Println("收到退出通知，主协程退出！")
}
