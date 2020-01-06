package main

import "fmt"

// Go语言是一个有指针类型的编程语言，当指针和接口同时出现时就会遇到一些让人困惑或者感到诡异的问题，
// 接口在定义一组方法时其实没有对实现的接受者做限制，所以我们其实会在一个类型上看到以下两种不同的实现方式：
// 1、实现接口的方法的接收者类型为普通类型
// 2、实现接口的方法的接收者类型为普通类型指针
// 注意：上述1、2两者不能同时存在

// 第1种情况：实现接口的方法的接收者类型为普通类型
/*
// Duck 接口
type Duck interface{
	Walk()
	Quack()
}

// Cat 结构体
type Cat struct{}

// Walk 方法
func (c Cat) Walk() {
	fmt.Println("cat walk!")
}

// Quack 方法
func (c Cat) Quack(){
	fmt.Println("meow!")
}

func main(){
	// 代码中的 Cat 结构体指针其实是能够直接调用 Walk 和 Quack 方法的，
	// 因为作为指针它能够隐式获取到对应的底层结构体
	var c Duck = Cat{}  // var c Duck = &Cat{} 也可以
	c.Walk()
	c.Quack()
}
*/


// 第2种情况：实现接口的方法的接收者类型为普通类型的指针

// Duck 接口
type Duck interface{
	Walk()
	Quack()
}

// Cat 结构体
type Cat struct{}

// Walk 方法
func (c *Cat) Walk() {
	fmt.Println("cat walk!")
}

// Quack 方法
func (c *Cat) Quack(){
	fmt.Println("meow!")
}

func main(){
	// 当代码中的变量是 Cat{} 时，调用函数其实会对参数进行复制，
	// 也就是当前函数会接受一个新的 Cat{} 变量，由于方法的参数是 *Cat，
	// 而编译器没有办法根据结构体找到一个唯一的指针，所以编译器会报错；
	// 当代码中的变量是 &Cat{} 时，在方法调用的过程中也会发生值的拷贝，
	// 创建一个新的 Cat 指针，这个指针能够指向一个确定的结构体，
	// 所以编译器会隐式的对变量解引用（dereference）获取指针指向的结构体完成方法的正常调用。
	var c Duck = Cat{}  // 编译报错，但 var c Duck = &Cat{} 可以
	c.Walk()
	c.Quack()
}