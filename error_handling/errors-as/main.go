package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		// errors.As 方法的作用是从 err 错误链中识别和 target 相同的类型，如果可以赋值，则返回 true。
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}