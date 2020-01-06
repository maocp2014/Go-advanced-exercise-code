package main

import "fmt"

func main() {
	var s = "中国人"

	// Go语言中，rune这种builtin类型被用来表示一个"Unicode字符"，
	// 因此一个rune的值就是其对应Unicode字符的序号，即码点。
	// 通过for...range语句对字符串进行迭代访问时，range会依次返回Unicode字符
	// 对应的rune，即码点。这里可以看到Unicode字符“中”对应的rune（码点）为0x4E2D。
	for _, v := range s {
		fmt.Printf("%s => 码点: %X\n", string(v), v)
	}
	
	// Unicode字符在存储和传输时采用的并非“理想编码方案”，而是多维UTF-8编码，
	// 也就是说在上面的例子中“中国人”这三个Unicode字符在内存中并不是以码点值存储的，
	// 而是以UTF-8编码后的值存储的
	fmt.Printf("%s => UTF8编码: ", s)
	// 我们将字符串转换为对应的切片元素，然后按字节逐一输出便得到了
	// Unicode字符“中国人”所对应的UTF-8编码，即存储“中国人”这个字符串时，内存所使用的字节(9个)和对应的值。
	for _, v := range []byte(s) {
			fmt.Printf("%X", v)
	}
	fmt.Printf("\n")
}
