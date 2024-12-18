// 练习 1.8: 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，
// 为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io"
	"strings"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
		prefix := "https://"                        // 定义http前缀
		if !strings.HasPrefix(url, prefix) {        // 判断是否存在http前缀
			fmt.Println("No prefix, add ", prefix)  
			url = prefix + url;	                    // 不是则添加上http前缀
		}
        resp, err := http.Get(url)                  // 请求访问，获取响应报文
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        b, err := io.ReadAll(resp.Body)             // 读取响应报文
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}
