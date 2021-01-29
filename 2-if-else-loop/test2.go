package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o") //这将返回一个strings.Replacer,将每个#替换为o
	fixed := replacer.Replace(broken)         //对strings.Replacer调用Relace方法，并传递一个字符串来进行替换
	fmt.Println(fixed)                        //打印Replace方法返回的字符串
}
