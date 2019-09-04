package main

import "fmt"

func main() {
	var name string
	var s = []byte("中国china!")

	name = string(s)
	fmt.Println(&name, name)

	name = "中国sss"
	fmt.Println(&name, name)  // 地址不会变
	
	// name[0] = "a"  // 非法，不能改变字符串中单个字符的值
	// 如果想要改变字符串中的值需要先将字符串转为字节数组([]byte)或者字符数组([]rune),
	// 有中文的情况下使用字符数组

	nstring := "abcxyz"
	fmt.Println(&nstring, nstring)

	nbytes := []byte(nstring)

	// 字节数组需要使用单引号,双引号是字符串了
	nbytes[0]='A'
	fmt.Println(&nstring, string(nbytes))

	// 字符串整体重新赋值是可以的
	name = "中国！"
	//如果有汉字的话需要使用字符数组
	nmrune := []rune(name)
	nmrune[1] = '华'
	fmt.Println(&name, string(nmrune))
}
