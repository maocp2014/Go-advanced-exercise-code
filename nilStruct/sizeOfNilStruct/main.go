package main

import (
	"fmt"
	"unsafe"
)

func main() {
	res := struct{}{}
	fmt.Println("占用空间:", unsafe.Sizeof(res))
}

// 占用空间: 0
