package main

import (
	"fmt"
	"strconv"
)

// 了解了它的机制之后，我们可以试着来绕过其限制，来完成一个可以内部化所有字符串的实现。
// 首先我们需要一个 pool，把所有的字符串都放到这个 pool 里，只要字符串在这个 pool
// 里只有一份（例如 Map 就是一个非常好的选择），就可以认为已经被 intern 了。下面是一个老外的实现：

type stringInterner map[string]string

func (si stringInterner) Intern(s string) string {
	if interned, ok := si[s]; ok {
		return interned
	}

	si[s] = s
	return s
}

func main() {
	si := stringInterner{}
	s1 := si.Intern("12")
	s2 := si.Intern(strconv.Itoa(12))
	fmt.Println(string(s1) == string(s2))  // true
}

/* string intern 作为一种高效的手段，在 Go 内部也有不少应用，
// 比如在 HTTP 中 intern 公用的请求头来避免多余的内存分配：
// commonHeader interns common header strings.

var commonHeader = make(map[string]string)

func init(){
	for _, v := range []string{
		"Accept",
		"Accept-Charset",
		"Accept-Encoding",       
		"Accept-Language",
		"Accept-Ranges",
		"Cache-Control",        
		// ...
	}{
		commonHeader[v]=v   
	}
}
*/
// 如果你在做缓存系统，或者是需要操作大量的字符串，不妨也考虑下 string intern 来优化你的应用。