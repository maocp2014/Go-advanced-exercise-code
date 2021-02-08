package main

// select 允许在一个 goroutine 中管理多个 channel。但是，当所有 channel 同时就绪的时候，
// go 需要在其中选择一个执行。此外，go 还需要处理没有 channel 就绪的情况，我们先从就绪的 channel 开始。
//
// select 不会按照任何规则或者优先级选择就绪的 channel。go 标准库在每次执行的时候，都会将他们顺序打乱，
// 也就是说不能保证任何顺序。

func main() {
	a := make(chan bool, 10)
	b := make(chan bool, 10)
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			print("<a ")

		case <-b:
			print("<b ")

		case <-c:
			print("<c ")

		default:
			print("<default ")
		}
	}
}