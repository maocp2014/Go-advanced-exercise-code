package main

import "fmt"

func main() {
	name := "中国男篮\n"  // 字符串只能用双引号""

	// ``号来定义原始输出，不会进行转义
	s := `2019世界蓝联
	篮球世界杯
	china
	`

	f := 'f'
	fmt.Println(name, s, string(f))
}
