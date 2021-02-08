package main

func main() {
	a := make(chan bool, 10)
	b := make(chan bool, 10)
	c := make(chan bool, 10)

	// 由于 go 不会删除重复的 channel，所以可以使用多次添加 case 来影响结果
	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			print("<a ")
		case <-a:
			print("<a ")
		case <-a:
			print("<a ")
		case <-a:
			print("<a ")
		case <-a:
			print("<a ")
		case <-a:
			print("<a ")
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