package main

import (
	"fmt"
	"time"
)

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}

func foo4() {
	values := []int{1, 2, 3, 5}

	for _, val := range values {
		go show(val)  // 传了val的副本
	}
}

func main() {
	// 因为Go Routine的执行顺序是随机并行的，因此执行多次foo4()输出的顺序不一行相同，
	// 但是一定打印了“1，2，3，5”各个元素。
	foo4() // 1 2 3 5输出顺序不确定
	time.Sleep(1)
}