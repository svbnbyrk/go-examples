package main

import "fmt"

const (
	first = iota + 5
	second
)

const (
	third = iota
	fourth
)

func main() {

	fmt.Println(first, second, third, fourth)
}
