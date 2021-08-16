package main

import "fmt"

func functionA(a int, b int) {
	fmt.Println(a + b)
}

func functionB(a int, b int) {
	fmt.Println(a * b)
}

func functionC(a bool) {
	fmt.Println(!a)
}

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}

func main() {
	functionA(2, 3)
	functionB(2, 3)
	functionC(true)
	functionD("$", 4)
	functionA(5, 6)
	functionB(5, 6)
	functionC(false)
	functionD("ha", 3)
}

/*
5
6
false
$$$$
11
30
true
hahaha
*/
