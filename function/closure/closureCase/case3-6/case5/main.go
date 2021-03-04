package main

import (
	"fmt"
	"time"
)

// Go Routine的延迟绑定
func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

func main() {
	foo5()
	time.Sleep(1)
}

// 其实这个问题的本质同闭包的延迟绑定，或者说，这段匿名函数的对象就是闭包。在我们调用go func() { xxx }()的时候，只要没有真正开始执行这段代码，那它还只是一段函数声明。而在这段匿名函数被执行的时候，才是内部变量寻找真正赋值的时候。
//
// 在case5中，for-loop的遍历几乎是“瞬时”完成的，4个Go Routine真正被执行在其后。**矛盾是不是产生了？**这个时候for-loop结束了呀，val生命周期早已结束了，程序应该报错才对呀？
//
// 回忆上一章，是不是一个相同的情境？是的，这个匿名函数可不就是一个闭包吗？一切就解释通了：闭包真正被执行的时候，for-loop结束了，但是val的生命周期在闭包内部被延长了且被赋值到最新的数值5。