package main

import "fmt"

// break在select内并不能跳出for-select循环。
// 看下面的例子，consume函数从通道inCh不停读数据，期待在inCh关闭后退出for-select循环，但结果是永远没有退出。

///3种解决办法
// 1、在满足条件的case内，使用return，如果有结尾工作，尝试交给defer。
///2、在select外for内使用break挑出循环，如注释的if判断语句。
///3、使用goto。

func consume(inCh <-chan int) {
	i := 0
	for {
		fmt.Printf("for: %d\n", i)
		select {
		case x, ok := <-inCh:
			if !ok {
				break
			}
			fmt.Printf("read: %d\n", x)
		}
		i++
		// if i == 5 {
		// 	break
		// }
	}

	fmt.Println("combine-routine exit")
}

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

func main() {
	inch := gen()
	consume(inch)
}
