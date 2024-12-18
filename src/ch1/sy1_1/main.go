// 练习1.1: 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。

package main		// 声明包

import (
	"fmt"			// 导入包fmt（格式化输出）
	"os"			// 导入包os（获取命令行参数）
)

func main() {				// 主函数
	fmt.Println(os.Args)	// 打印所有参数（包括程序本身）
}