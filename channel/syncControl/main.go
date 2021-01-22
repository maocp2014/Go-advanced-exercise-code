package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 同步控制模型，生产者模型
var lockChan = make(chan int, 1)
var remainMoney = 1000

func main() {
	quit := make(chan bool, 2)

	go func() {
		for i := 0; i < 10; i++ {
			money := (rand.Intn(12) + 1) * 100
			go expense(money)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		quit <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			money := (rand.Intn(12) + 1) * 100
			go gain(money)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		}

		quit <- true
	}()

	<-quit
	<-quit

	fmt.Println("主程序退出")
}

func expense(money int) {
	lockChan <- 0

	if remainMoney >= money {
		srcRemainMoney := remainMoney
		remainMoney -= money
		fmt.Printf("原来有%d, 花了%d，剩余%d\n", srcRemainMoney, money, remainMoney)
	} else {
		fmt.Printf("想消费%d钱不够了, 只剩%d\n", money, remainMoney)
	}

	<-lockChan
}

func gain(money int) {
	lockChan <- 0

	srcRemainMoney := remainMoney
	remainMoney += money
	fmt.Printf("原来有%d, 赚了%d，剩余%d\n", srcRemainMoney, money, remainMoney)

	<-lockChan
}
