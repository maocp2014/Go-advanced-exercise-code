package selectClosedChannel

import "fmt"
import "time"

func main() {
	var c1, c2 chan int = make(chan int), make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		c2 <- 7
		close(c2)
	}()

	for {
		select {
		case x := <-c1:
			fmt.Println(x)
		case x := <-c2:
			fmt.Println(x)
		}
	}
	fmt.Println("over")
}


// 我们原本期望程序交替输出5和7两个数字，但实际的输出结果却是:
// 5
// 0
// 0
// 0
// ....  // 后续一直输出0，死循环，这是因为c1被关闭，被关闭的channel并不会阻塞，因此并不会执行 case x := <-c2分支


// 再仔细分析代码，原来select每次按case顺序evaluate：
// – 前5s，select一直阻塞；
// – 第5s，c1返回一个5后被close了，“case x := <-c1”这个分支返回，select输出5，并重新select
// – 下一轮select又从“case x := <-c1”这个分支开始evaluate，由于c1被close，按照前面的知识，
//   close的channel不会阻塞，我们会读出这个 channel对应类型的零值，这里就是0；select再次输出0；
//   这时即便c2有值返回，程序也不会走到c2这个分支
// – 依次类推，程序无限循环的输出