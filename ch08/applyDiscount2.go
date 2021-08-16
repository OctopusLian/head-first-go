package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func applyDiscount2(s subscriber) {
	s.rate = 4.99
}

func main() {
	var a subscriber
	applyDiscount(&a)
	fmt.Println(a.rate)
	applyDiscount2(a)
	fmt.Println(a.rate)
}
