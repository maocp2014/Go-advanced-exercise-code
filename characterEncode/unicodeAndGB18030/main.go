// github.com/bigwhite/experiments/non-ascii-char-encoding/demo1.go

package main

import (
	"fmt"

	utils "github.com/bigwhite/gocmpp/utils"
)

// 关于中文字符编码，需记住以下要点：
// Unicode是目前被支持最为广泛的字符集；
// Utf-8是目前被支持最为广泛的Unicode字符的编码方式(还有其他方式，比如UTF-16、UTF-32等)；
// 针对同一个字符，比如：“中”，如果该字符存在于两个字符集编码方案A（比如：utf8)和B(比如gb18030)中，
// 那么我们可以通过转换，将该字符在A中的编码(如："中"的E4B8AD)转换为在B中的编码(如“中”的D6D0)。

func main() {
	var stringLiteral = "中国人"
	var stringUsingRuneLiteral = "\u4E2D\u56FD\u4EBA" // "中国人"对应的unicode码点

	if stringLiteral != stringUsingRuneLiteral {
		fmt.Println("stringLiteral is not equal to stringUsingRuneLiteral")
		return
	}
	fmt.Println("stringLiteral is equal to stringUsingRuneLiteral")

	for i, v := range stringLiteral {
		fmt.Printf("中文字符: %s <=> Unicode码点(rune): %X <=> UTF8编码(内存值): ", string(v), v)
		s := stringLiteral[i : i+3]
		for _, v := range []byte(s) {
			fmt.Printf("0x%X ", v)
		}

		s1, _ := utils.Utf8ToGB18030(s)
		fmt.Printf("<=> GB18030编码(内存值): ")
		for _, v := range []byte(s1) {
			fmt.Printf("0x%X ", v)
		}
		fmt.Printf("\n")
	}
}
