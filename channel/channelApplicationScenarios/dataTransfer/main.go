package main

import (
	"fmt"
	"time"
)

type token struct{}

func main() {
	num := 4
	var chs []chan token

	chs = make([]chan token, num)

	// 4个channel
	for i := 0; i < num; i++ {
		chs = append(chs, make(chan token))
	}

	fmt.Printf("chs: %v\n", len(chs))

	// 4个worker
	for j := 0; j < num; j++ {
		go worker(j, chs[j], chs[(j+1)%num])
	}

	// 先把令牌交给第一个
	chs[0] <- struct{}{}

	select {}
}

func worker(id int, ch chan token, next chan token) {
	for {
		// 对应work 取得令牌
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(1 * time.Second)
		// 传递给下一个
		next <- token
	}
}