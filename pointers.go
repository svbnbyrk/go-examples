package main

import "fmt"

func pointers() {
	// var firstName *string = new(string)

	// *firstName = "sevban"

	// fmt.Println(*firstName)

	firstName := "sevban"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "ahmet"
	fmt.Println(ptr, *ptr)

}
