package main

import (
	"fmt"
	"time"
)

// B 返回函数切片[]func()
func B() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		b[i] = func() {
			fmt.Println(i)
		}
		time.Sleep(1 * time.Second)
	}
	return b
}

func main() {
	c := B()
	c[0]()
	c[1]()
	c[2]()
}

// output: 因为都引用i，i最后变成了3
// 3
// 3
// 3
