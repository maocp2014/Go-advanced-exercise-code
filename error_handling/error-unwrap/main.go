package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("脑子进煎鱼了")
	// 变量 w 就是一个嵌套一层的 error
	// 最外层是 “快抓住：”，此处调用 %w 意味着 Wrapping Error 的嵌套生成。
	// 需要注意的是，Go 并没有提供 Warp 方法，而是直接扩展了 fmt.Errorf 方法。
	w := fmt.Errorf("快抓住：%w", e)
	fmt.Println(w)
	// 该方法的作用是将嵌套的 error 解析出来，若存在多级嵌套则需要调用多次 Unwarp 方法。
	fmt.Println(errors.Unwrap(w))
}