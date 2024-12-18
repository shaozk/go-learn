// 练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存
// 策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获
// 取到的响应内容是否一致，修改本节中的程序，将响应结果输出到文件，以便于
// 进行对比。

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string) // 新键通道（用于在routine之间传输数据）
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
	// 创建新文件（以时间戳命名）
	fileName := fmt.Sprintf("%v.txt", start.Nanosecond())
	file, err := os.Create(fileName)    
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// 将响应结果保存至文件中
    nbytes, err := io.Copy(file, resp.Body)
    resp.Body.Close() // don't leak resources
	file.Close()	// 关闭文件

    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
