// 练习 1.4: 修改 dup2，出现重复的行时打印文件名称

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
    "bufio"			// 导入bufio包（标准输入）
    "fmt"
    "os"
)

// 自定义计数器
type MyCount struct {
	count int			// 记录行数
	names []string		// 记录行所在文件
}

func main() {
    counts := make(map[string]*MyCount)		// 构造<string,MyCount*>字典
    files := os.Args[1:]
    if len(files) == 0 {					// 判断是否包含文件
        countLines(os.Stdin, counts)		// 从标准输入获取行
    } else {
	
        for _, arg := range files {
            f, err := os.Open(arg)			// 打开文件
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)			// 从文件获取行
            f.Close()						// 关闭文件
        }
	}
	printLines(counts, len(files) != 0)		// 输出字典
}

// 构建以字符串为值、以自定义计数器为键的map
// 按行获取字符串，判断map中是否包含该字符串，是则计数+1，将文件名添加到列表中（如果不存在）
// 否则新键一个键值对插入到map中
// [in] f 文件
// [out] counts 自定义计数器map 
func countLines(f *os.File, counts map[string]*MyCount) {
    input := bufio.NewScanner(f)
    for input.Scan() {						// 按行扫描
		line := input.Text()				// 获取文本
        if _, ok := counts[line]; ok {		// 如果字典中存在该字符串
			counts[line].count++;			// 计数加1
			if !contains(counts[line].names, f.Name()) {		// 判断是否已经包含文件名
				counts[line].names = append(counts[line].names, f.Name())	// 添加文件名
			}
		} else {
			counts[line] = &MyCount{		// 插入新键-值对
				1,							// 计数为1
				make([]string, 1),			// 空的字符串数组
			}
			counts[line].names[0] = f.Name()	// 填充字符串数组
		}
    }
    // NOTE: ignoring potential errors from input.Err()
}

// 根据是否是文件输出自定义计数器map
// [in] counts	自定义计数器map
// [in] flagFile 文件标志
func printLines(counts map[string]*MyCount, flagFile bool) {
	for line, c := range counts {
		if c.count > 1 {
			if flagFile {		// 判断是否由文件读取
				fmt.Printf("%d\t%s\t%v\n", c.count, line, c.names)
			} else {
				fmt.Printf("%d\t%s\n", c.count, line)
			}
		}
	}
}

// 判断字符串数组是否包含某一个字符串(暴力法)
// [in] slice 字符串数组
// [in] s 字符串
// [out] 返回是否包含
func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false 
}
