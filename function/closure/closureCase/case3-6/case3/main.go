package main

import (
	"fmt"
)

func foo3() {
	values := []int{1, 2, 3, 5}

	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func main() {
	foo3() // 1 2 3 5
}