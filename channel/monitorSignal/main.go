package main

import(
	"fmt"
	"time"
	"os"
	"os/signal"
)

func main(){
	fmt.Println("按Ctrl+C可退出程序")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		time.Sleep(time.Second * 2)
	
		number := 0;
		for{
			number++
			println("number : ", number)
			time.Sleep(time.Second)
		}
	}()

	<- quit

	fmt.Println("主程序退出！")
}