package main

import "fmt"

func negate(myBoolean *bool) {
	//negate接受一个指向布尔值的指针，而不是直接接受布尔值，然后将该指针处的值更新为相反的值。
	*myBoolean = !*myBoolean
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
