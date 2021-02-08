package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0 ; i < 5 ; i++{
		go func() {
			fmt.Println("current i is ", i)
		}()
	}
	time.Sleep(time.Second)
}

// 上述输出不确定，以下是可能的输出情况
//
// current i is  4
// current i is  5
// current i is  5
// current i is  5
// current i is  5

//
// current i is  3
// current i is  5
// current i is  5
// current i is  5
// current i is  5

// current i is  5
// current i is  5
// current i is  5
// current i is  2
// current i is  5

// current i is  5
// current i is  5
// current i is  5
// current i is  5
// current i is  5

// ...