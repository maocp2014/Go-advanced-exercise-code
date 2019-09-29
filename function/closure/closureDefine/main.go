package main

import "fmt"

// Closures are a special case of anonymous functions. Closures are anonymous functions
// which access the variables defined outside the body of the function.

/*
package main

import (
    "fmt"
)

func main() {
    a := 5
    func() {
        fmt.Println("a =", a)
    }()
}

// In the program above, the anonymous function accesses the variable a
// which is present outside its body in line no. 10. Hence this anonymous function is a closure.
*/

// Go 函数可以是一个闭包。闭包是一个函数值，它引用了函数体之外的变量。
// 这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

// 没有闭包的时候，函数就是一次性买卖，函数执行完毕后就无法再更改函数中变量的值（应该是内存释放了）；
// 有了闭包后函数就成为了一个变量的值，只要变量没被释放，函数就会一直处于存活并独享的状态，
// 因此可以后期更改函数中变量的值（因为这样就不会被go给回收内存了，会一直缓存在那里）。

// 闭包是匿名函数与匿名函数所引用环境的组合。匿名函数有动态创建的特性，
// 该特性使得匿名函数不用通过参数传递的方式，就可以直接引用外部的变量。
// 这就类似于常规函数直接使用全局变量一样，个人理解为：匿名函数和它引用的变量以及环境，
// 类似常规函数引用全局变量处于一个包的环境。

func main() {
	n := 0
	// 匿名函数，f是变量
	f := func() int {
		n++
		return n
	}
	fmt.Println(f()) // 别忘记括号，不加括号相当于地址
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// output:
// 1
// 2
// 3
// 4

// 在上述代码中，
// n := 0
// f := func() int {
// 	n += 1
// 	return n
// }
// 就是一个闭包，类比于常规函数+全局变量+包。f不仅仅是存储了一个函数的返回值，它同时存储了一个闭包的状态。
