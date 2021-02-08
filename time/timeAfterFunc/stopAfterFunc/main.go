package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 :=time.AfterFunc(3 * time.Second, func() {
		fmt.Println("3 seconds over....")
	})

	time.Sleep(1 * time.Second)

	res := timer1.Stop()

	fmt.Println("已取消等待和函数执行, res: ", res)

	time.Sleep(5 * time.Second)

	fmt.Println("ok")
}