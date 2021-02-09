package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Context 类型提供了 Done() 方法，每次 context 接收到取消事件时，该方法都是返回一个 channel，
// 这个 channel 会收到空结构体类型的数据。监听取消事件也很容易，<- ctx.Done()。


// 比如，一个 HTTP 请求处理需要两秒，如果在中途取消就必须立即返回。
// 将服务跑起来，在浏览器中打开 localhost:8000，如果你在 2s 钟之内关闭浏览器，
// 终端将会输出 request cancelled。
func main() {
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "processing request\n")
		// We use `select` to execute a peice of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
}