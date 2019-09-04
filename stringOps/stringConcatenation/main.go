package main

// 可以直接使用 + 拼接，但不高效
/*
func main() {
	//字符串可以使用+进行拼接
	var name = "yang" + "yan" + "xing"
	fmt.Println(name)
	//使用+=在原字符串后面追加字符串
	name += "杨彦星"
	fmt.Println(name)
	// + 不能在行首，只能写在行尾，就和if else中的else那样
	s := "BeiJing" +
		" TongZhou"
	fmt.Println(s)
}
*/

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"yang", "fan", "zhang", "li", "chen"}
	n := strings.Join(names, " ")
	fmt.Println(n)
	s := "yang&aaa&bbb&ccc"
	//以&分割
	str := strings.Split(s, "&")
	fmt.Println(str)
}
