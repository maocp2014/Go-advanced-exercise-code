package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

/*
https://juejin.im/post/5d01177a5188254b9000975c
这篇文章介绍了，在 Go 中如何正确地构建流式数据 pipeline。它的异常处理非常复杂，
pipeline 中的每个 stage 都可能导致上游阻塞，而下游可能不再关心接下来的数据。
关闭 channel 可以给所有运行中的 goroutine 发送 done 信号，这能帮助我们成功解除阻塞。
*/

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

/*
在并行版本中，MD5All 为每个文件启动了一个 goroutine。但如果一个目录中文件太多，这可能会导致分配的内存过大
以至于超过了当前机器的限制。我们可以通过限制并行读取的文件数，限制内存分配。在并发限制版本中，我们创建了固
定数量的 goroutine 读取文件。现在，我们的 pipeline 涉及 3 个 stage：遍历目录、文件读取与摘要计算、结果收集。
*/

// 第一个 stage，遍历目录并通过 paths channel 发出文件。
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// Close the paths channel after Walk returns.
		defer close(paths)
		// No select needed for this send, since errc is buffered.
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

// 第二个 stage，启动固定数量的 goroutine，从 paths channel 中读取文件名称，处理结果发送到 c channel。
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

// MD5All 计算md5值
// 最后一个 stage，负责从 c 中接收处理结果，通过 errc 检查是否有错误发生。该检查无法提前进行，
// 因为提前执行将会阻塞 walkFile 往下游发送数据。
func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All closes the done channel when it returns; it may do so before
	// receiving all the values from c and errc.
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result) // HLc

	var wg sync.WaitGroup

	const numDigesters = 20

	wg.Add(numDigesters)

	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c) // HLc
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c) // HLc
	}()
	// End of pipeline. OMIT

	m := make(map[string][md5.Size]byte)

	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	// Check whether the Walk failed.
	if err := <-errc; err != nil { // HLerrc
		return nil, err
	}
	return m, nil
}

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
