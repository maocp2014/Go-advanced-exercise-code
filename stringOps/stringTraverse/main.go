package main

import "fmt"

func main() {
	var name = "china中国！"
	//按字节遍历
	for i := 0; i < len(name); i++ {
		fmt.Printf("%v %T:%v\n", i, name[i], name[i])
	}
	fmt.Println("")
	//按字符遍历
	for _, r := range name {
		//%c 为输出数值所表示的 Unicode 字符,不带单引号 如 y
		//%q 输出数值所表示的 Unicode 字符（带单引号）如 '杨'
		fmt.Printf("%c", r)
	}
}
