package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 所谓字符串内部化（string interning），就是一种技术手段让相同的字符串在内存中只保留一份。
// 这样就可以大幅降低内存占用，缩短字符串比较的时间。因为相同的字符串只需要保存一份在内存中，
// 当用这个字符串做匹配时，比较字符串只需要比较地址是否相同就够了，而不必逐字节比较。
// 于是时间复杂度就从 O(N) 降低到了 O(1)。

/* Go 的字符串，本质上是一个 reflect.StringHeader:
type StringHeader struct {
	Data uintptr
	Len  int
}
// 其中 Data 指针指向的是一个字符常量的地址，这个地址里面的内容是不可以被改变的，
// 因为它是只读的，但是这个指针可以指向不同的地址。虽然相同的字符串是不同的 StringHeader，
/* 但是其内部实际上都指向相同的字节数组

/* 需要注意的是 Go 的 string intern 仅仅针对的是编译期可以确定的字符串常量，
// 如果是运行期间产生的字符串则不能被内部化。比如：
// 可以被 intern
s1 := "12"
s2 := "1" + "2"

// 不能被 intern
s3 := "12"
s4 := strconv.Itoa(12)
*/

func main() {
	str1 := "Hello, World!"
	str2 := "Hello, World!"

	// 这两个地址并不相同
	fmt.Printf("str add: %p, %p\n", &str1, &str2)

	x1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	x2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	// 底层都是指向相同的 []byte
	fmt.Printf("data add: %#v, %#v\n", x1.Data, x2.Data)
}
