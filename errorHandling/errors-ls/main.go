package main


import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		// errors.Is 方法的作用是判断所传入的 err 和 target 是否同一类型，如果是则返回 true。
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}
