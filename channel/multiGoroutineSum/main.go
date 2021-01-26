package main

import "fmt"

func calculate(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func gather(c chan int, count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += <-c
	}
	return sum
}

func main() {
	// 假设这个数组非常大，需要加快计算速度
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14}

	count := 3
	c := make(chan int, count)

	for i := 1; i <= count; i++ {
		go func(i int) {
			ret := 0
			if i == count {
				ret = calculate(arr[(i-1)*count:])
			} else {
				ret = calculate(arr[(i-1)*count:i*count])
			}
			c <- ret
		}(i)
	}

	ret := gather(c, count)

	fmt.Println("ret:", ret)
}
