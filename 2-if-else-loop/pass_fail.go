package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")        //提示用户输入一个百分比分数
	reader := bufio.NewReader(os.Stdin) //创建一个bufio.Reader，它允许我们读取键盘输入
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            //将换行符从输入中删除
	grade, err := strconv.ParseFloat(input, 64) //将输入字符串转换为float64（数字）值
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)

}

/*
Enter a grade: 100
A grade of 100 is passing
*/
