package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 在实际的工程中，不可能进行延时，这样就没有并发的优势，一般采取下面两种方法：

// 方法1、共享的环境变量作为函数参数传递:

/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
*/

// output:
// 4
// 3
// 0
// 2
// 1

// 输出结果不一定按照顺序，这取决于每个goroutine的实际情况，但是最后的结果是不变的。
// 可以理解为，函数参数的传递是瞬时的，而且是在一个goroutine执行之前就完成，所以此时执行的闭包存储了当前i的状态。

// 方法2、使用同名的变量保留当前的状态

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i // 注意这里的同名变量覆盖
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

// output:
// 0
// 4
// 1
// 2
// 3

// 同名的变量i作为内部的局部变量，覆盖了原来循环中的i，此时闭包中的变量不在是共享外循环的i，
// 而是都有各自的内部同名变量i，赋值过程发生于循环goroutine，因此保证了独立，这里不同名也是没问题的。
