package main

import (
	"fmt"
)

func main() {
	fmt.Print("whats ur name? ")
	var name string
	var lastname string
	fmt.Scanln(&name, &lastname)
	fmt.Printf("Hello %s %s nice to meet u\n", name, lastname)
}
