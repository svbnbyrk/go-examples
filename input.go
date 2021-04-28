package main

import (
	"fmt"
)

func input() {
	var name string
	fmt.Println("whats ur name?")
	//%s boşluk duyarlıdır kelime değeri alır
	inputs, _ := fmt.Scanf("%q", &name)

	switch inputs {
	case 0:
		fmt.Printf("u must enter a name \n")
	case 1:
		fmt.Printf("hello %s! nice to meet u", name)
	}
}
