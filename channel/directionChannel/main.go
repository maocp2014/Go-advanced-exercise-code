package main

import(
	"fmt"
)

/*
var ch1 chan int       // ch1是一个正常的channel，不是单向的
var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
var ch3 <-chan int     // ch3是单向channel，只用于读取int数据

chan<- 表示数据进入管道，要把数据写进管道，对于调用者就是输出。
<-chan 表示数据从管道出来，对于调用者就是得到管道的数据，当然就是输入。
*/

/*
func main(){
c := make(chan int, 3)

var send chan<- int = c // send-only
var recv <-chan int = c // receive-only

send <- 1
// <-send //invalid operation: <-send (receive from send-only type chan<- int)

<-recv
// recv <- 2 //invalid operation: recv <- 2 (send to receive-only type <-chan int)


//不能将单向 channel 转换为普通 channel
// d1 := (chan int)(send) //cannot convert send (type chan<- int) to type chan int
// d2 := (chan int)(recv) //cannot convert recv (type <-chan int) to type chan int
}
*/


//   chan<- //只写
func counter(out chan<- int) {
    defer close(out)

    for i := 0; i < 5; i++ {
        out <- i //如果对方不读 会阻塞
    }
}

//   <-chan //只读
func printer(in <-chan int) {
    for num := range in {
        fmt.Println(num)
    }
}

func main() {
    c := make(chan int) //   chan   //读写

    go counter(c) //生产者
    printer(c)    //消费者

    fmt.Println("done")
}