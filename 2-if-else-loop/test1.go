package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()  //返回一个当前日期和时间的time.Time的值
	var year = now.Year() //返回一个年份的值
	fmt.Println(year)     //输出
}
