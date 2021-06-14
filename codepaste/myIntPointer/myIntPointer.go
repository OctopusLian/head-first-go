package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	*myIntPointer = 42
	fmt.Println(*myIntPointer) //42
	fmt.Println(myIntPointer)  //0xc00011e058
	fmt.Println(myInt)         //42
}
