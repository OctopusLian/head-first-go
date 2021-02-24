package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	seconds := time.Now().Unix() //获取当前日期和时间的整数形式
	rand.Seed(seconds)           //播种随机数生成器
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.0")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Make a guess: ")

}
