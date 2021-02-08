package main

import (
	"fmt"
	"time"
)

// func After(d Duration) <-chan Time {
//    return NewTimer(d).C
// }
// 可以看到，After() 底层是基于 Timer 实现的，该函数返回 time.Time 类型的 channel，
// 当定时时间到了之后， 会将数据（当前时间）写入 channel。
//
// 调用该函数不会阻塞当前协程，但如果执行读取操作时，还未达到定时时间，会发生阻塞。

func main() {
	fmt.Println("now1: ", time.Now())

	after1 := time.After(2*time.Second)
	fmt.Println("before1, time: ", time.Now())

	curTime1 := <-after1    // 读取
	fmt.Println("after1, time: ", time.Now())
	fmt.Println("curTime1: ", curTime1)

	fmt.Println()

	fmt.Println("now2: ", time.Now())

	after2 := time.After(2*time.Second)
	time.Sleep(3*time.Second)         // 为了使定时时间过期
	fmt.Println("before2, time: ",time.Now())

	curTime2 := <-after2   // 读取
	fmt.Println("after2, time: ",time.Now())
	fmt.Println("curTime2: ", curTime2)
}

// 执行读取操作时，如果定时器还未过期，读取操作将会阻塞；否则不会阻塞。