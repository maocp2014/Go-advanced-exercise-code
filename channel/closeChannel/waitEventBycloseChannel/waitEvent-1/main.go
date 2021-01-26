package waitEvent_1

import "fmt"

func main() {

	fmt.Println("Begin doing something!")

	c := make(chan bool)

	go func() {
		fmt.Println("Doing something…")
		// 等待一个事件，有时候通过close一个Channel就足够了
		close(c)
	}()

	<-c

	fmt.Println("Done!")
}

// 这里main goroutine通过"<-c"来等待sub goroutine中的“完成事件”，
// sub goroutine通过close channel促发这一事件。当然也可以通过向Channel写入一个bool值的方式来作为事件通知。
// main goroutine在channel c上没有任何数据可读的情况下会阻塞等待。