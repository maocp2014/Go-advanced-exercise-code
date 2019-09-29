// 死锁：卡住主协程，系统一直在等待，所以系统判断为死锁，最终报deadlock错误并结束程序

// 1、当一个channel中没有数据，而直接读取时，会发生死锁：
/*
func main() {
    q := make(chan int, 2)
    <-q
}
*/

// 2、写入数据超过channel的容量，也会造成死锁：
/*
func main() {
    q := make(chan int, 2)
    q <- 1
    q <- 2
    q <- 3
}
*/

// 3、向已经关闭的channel中写入数据，这种造成的不是死锁，而是产生panic。
/*
func main() {
    q := make(chan int, 2)
    close(q)
    q <- 1
}
*/

/* 这种情况不会发生死锁
func main() {
	ch := make(chan string)
	go func() {
		ch <- "send"
	}()

	time.Sleep(time.Second * 3)
}
因为虽然子协程一直阻塞在传值语句，但这也只是子协程的事。外面的主协程还是该干嘛干嘛，
等你三秒之后就发车走人了。因为主协程都结束了，所以子协程也只好结束（毕竟没搭上车只能回家了，光杵在哪也于事无补）
*/

// 4、上面场景如果主协程与子协程有channel连接，那么会发生死锁
/*
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch2 <- "ch2 value"
		ch1 <- "ch1 value"
	}()

	<-ch1
}
*/

// 如下情况则不会发生死锁
/*
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    go func() {
        ch2 <- "ch2 value"
        ch1 <- "ch1 value"
    }()
 
    go func() {
        <- ch1
        <- ch2
    }()
 
    time.Sleep(time.Second * 2)
}
*/

// 5、for...range情况
/*
func main() {
    chs := make(chan string, 2)
    chs <- "first"
    chs <- "second"
 
    for ch := range chs {
        fmt.Println(ch)
    }
}
*/

// 6、同一个goroutine中，使用同一个 channel 读写
/*
func main(){
    ch:=make(chan int)  //这就是在main程里面发生的死锁情况
    ch<-6   //  这里会发生一直阻塞的情况，执行不到下面一句
    <-ch
}
*/

// 7、2个以上的goroutine中，使用同一个channel通信。读写channel先于goroutine创建
/*
package main

func main(){
    ch:=make(chan int)
    ch<-666    //这里一直阻塞，运行不到下面
    go func (){
        <-ch  //这里虽然创建了子go程用来读出数据，但是上面会一直阻塞运行不到下面
    }()
}
*/