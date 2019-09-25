package main

import "fmt"

// select功能：
// 在多个通道上进行读或写操作，让函数可以处理多个事情，但1次只处理1个。以下特性也都必须熟记于心：
// 1、每次执行select，都会只执行其中1个case或者执行default语句。
// 2、当没有case或者default可以执行时，select则阻塞，等待直到有1个case可以执行。
// 3、当有多个case可以执行时，则随机选择1个case执行。
// 4、case后面跟的必须是读或者写通道的操作，否则编译出错。

func main() {
	readCh := make(chan int, 1)
	writeCh := make(chan int, 1)

	y := 1

	select{
	case x := <- readCh:
		fmt.Printf("Read %d\n", x)
	case writeCh <- y:
		fmt.Printf("write %d\n", y)
	default:
		fmt.Println("Do what you want?")
	}
}
