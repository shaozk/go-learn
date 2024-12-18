// 练习 1.3: 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。

package main

import (
	"os"
	"fmt"
	"time" 		// 导入包time（计时器）
	"strings"	// 导入包strings（字符串）
)

func main() {
	start1 := time.Now()			// 计时器1，用于计算老方法时间开销
	str, sep := "", ""				// 初始化str存储连接后的字符串，初始化sep存储空格
	for _, arg := range os.Args {
		str += sep + arg 			// 连接参数字符串
		sep = " "					// 经过一个参数后，sep换为空格
	}
	fmt.Println(str)				// 输出连接后的值
	
	fmt.Printf("old way: %dns elapsed\n", time.Since(start1).Nanoseconds())		// 打印时间开销

	start2 := time.Now()						// 计时器2，用于计算新方法时间开销
	fmt.Println(strings.Join(os.Args, " "))		// 输出连接后的值
	fmt.Printf("new way: %dns elapsed\n", time.Since(start2).Nanoseconds())		// 打印时间开销

}