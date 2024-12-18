// 练习 1.2: 修改 echo 程序，使其打印每个参数的索引和值，每个一行。

package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, value := range os.Args { 	// 遍历命令行参数
		fmt.Println(idx, value)			// 打印索引和参数
	}
}
