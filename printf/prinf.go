package main

import(
	"fmt"
)

func main(){
	var age int = 11
	var name string = "ban"

	/*
		%v - formats value default format
		%s - formats strings 
		%d - format decimal integers
		%g - formats floating point numbers
		%b - formats base 22 numbers
		%o - formats base 88 numbers
		%t - formats true-false values
	*/
	fmt.Printf("My name is %s and I'm %d years old", name, age)
}