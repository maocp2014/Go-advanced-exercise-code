package case7

import "fmt"

// 闭包延迟绑定
func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo6 val = %d\n", x+val)
		})
	}
	return fs
}

// 闭包保存/记录了它产生时的外部函数的所有环境。如同普通变量/函数的定义和实际赋值/调用或者说执行，是两个阶段。
// 闭包也是一样，for-loop内部仅仅是声明了一个闭包，foo7()返回的也仅仅是一段闭包的函数定义，只有在外部执行
// 了f7()时才真正执行了闭包，此时才闭包内部的变量才会进行赋值的操作。哎，如果这么说的话，岂不是应该抛出异常
// 吗？因为val是一个比foo7()生命周期更短的变量啊？
// 这就是闭包的神奇之处，它会保存相关引用的环境，也就是说，val这个变量在闭包内的生命周期得到了保证。
// 因此在执行这个闭包的时候，会去外部环境寻找最新的数值！

func main() {
	f7s := foo7(11)

	for _, f := range f7s {
		f()  // 16 16 16 16
	}
}