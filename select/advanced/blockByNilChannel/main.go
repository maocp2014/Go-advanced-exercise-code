package main

import "fmt"

// select中nil channel永远阻塞

// 当select case读一个channel时，如果这个channel为nil，则该case永远阻塞。这个功能有1个妙用，
// select通常处理的是多个channel，当某个读channel关闭了，但不想select再继续读此channel，即忽略此case，
// 而是关注其他case，则把该通道设置为nil即可。

func gen() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func combine(inCh1, inCh2 <-chan int) <-chan int {
	// 输出通道
	out := make(chan int)

	// 启动协程合并数据
	go func() {

		for {
			select {
			case x, ok := <-inCh1:
				if !ok {
					inCh1 = nil // 当channel中无数据时，将channel设置为nil，使其一直阻塞，读其它channel
					break
				}
				out <- x
			case x, ok := <-inCh2:
				if !ok {
					inCh2 = nil // 当channel中无数据时，将channel设置为nil，使其一直阻塞，读其它channel
					break
				}
				out <- x
			}

			// 当ch1和ch2都关闭时才退出
			if inCh1 == nil && inCh2 == nil {
				break
			}
		}
		close(out)
	}()

	fmt.Println("combine exit")
	return out
}

func main() {
	ch1 := gen()
	ch2 := gen()

	out := combine(ch1, ch2)

	for x := range out {
		fmt.Println(x)
	}
}
