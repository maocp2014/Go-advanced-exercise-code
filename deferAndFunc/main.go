package main

import "fmt"

func main() {
	a := 1
	// 不管运行顺序如何，当参数为函数的时候，要先计算参数的值
	defer print(function(a))
	a = 2
}

func function(num int) int {
	return num
}
func print(num int) {
	fmt.Println(num)
}
