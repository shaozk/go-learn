// 练习 1.7： 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果
// 写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到
// os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结
// 果中的错误。

// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io"
    "net/http"      // 网络
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        // b, err := io.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)  // 拷贝数据到标准输出
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        resp.Body.Close()
    }
}
