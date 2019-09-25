package main

import (
	"fmt"
	"time"
)

/*
使用退出通道退出
使用,ok来退出使用for-select协程，解决是当读入数据的通道关闭时，没数据读时程序的正常结束。想想下面这2种场景，,ok还能适用吗？

接收的协程要退出了，如果它直接退出，不告知发送协程，发送协程将阻塞。
启动了一个工作协程处理数据，如何通知它退出？
使用一个专门的通道，发送退出的信号，可以解决这类问题。

以第2个场景为例，协程入参包含一个停止通道stopCh，当stopCh被关闭，case <-stopCh会执行，直接返回即可。

当我启动了100个worker时，只要main()执行关闭stopCh，每一个worker都会都到信号，进而关闭。如果main()向stopCh发送100个数据，这种就低效了。
*/

func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")

		t := time.NewTicker(time.Millisecond * 500)

		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh: // 当stopCh被关闭，case <-stopCh会执行
				fmt.Println("Recv stop signal")
				return
			case <-t.C:
				fmt.Println("Working .")
			}
		}
	}()
	return
}

func main() {

	stopCh := make(chan struct{})
	worker(stopCh)

	time.Sleep(time.Second * 2)
	close(stopCh)

	// Wait some print
	time.Sleep(time.Second)
	fmt.Println("main exit")
}

// Working .
// Working .
// Working .
// Working .
// Recv stop signal
// worker exit
// main exit
