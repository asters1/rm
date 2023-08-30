package main

import (
	"fmt"
	"os"
	"strings"
)

//		fmt.Printf("\033[1;31;40m%s\033[0m\n", str)
//		fmt.Printf("\033[1;37;41m%s\033[0m\n", str)
//	  输出后第一行是红字黑底，第二行红底白字，可以看到这里的输出与正常输出多了点东西，下面一个一个来看看。
//
// \033 这是标记变换颜色的起始标记，这之后的[1;31;40m则是定义颜色，1表示代码的意义或者说是显示方式，31表示前景颜色，40则是背景颜色。在这定义之后终端就会显示你设定的样式，如果只是要改变一行的样式则在结尾加入\033[0m表示恢复终端默认样式。
//
// 颜色和配置的取值范围：
// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//
//	0  终端默认设置
//	1  高亮显示
//	4  使用下划线
//	5  闪烁
//	7  反白显示
//	8  不可见
func IsDIR(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return true
	}
	return fi.IsDir()
}

func formatPath(path string) string {
	if strings.HasPrefix(path, "./") {
		path = path[2:]
	}
	return path
}

func view(path string) {
	// fmt.Println(IsDIR(path))
	if IsDIR(path) {
		path = formatPath(path)
		PrintBlue(path)
	} else {
		path = formatPath(path)
		fmt.Print(path)
	}
	fmt.Print("  ")
}

func PrintBlue(str string) {
	fmt.Printf("\033[1;34;m%s\033[0m", str)
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		// fmt.Println(i)
		view(os.Args[i])
	}
	fmt.Print("\n\n确定删除以上文件吗 [yn]? ")
	yorn := ""
	fmt.Scan(&yorn)
	if yorn == "y" || yorn == "Y" {
		for i := 1; i < len(os.Args); i++ {
			// fmt.Println(i)
			os.RemoveAll(os.Args[i])
		}
	} else {
		fmt.Println("删除取消!")
	}
}
