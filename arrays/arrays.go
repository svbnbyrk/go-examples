package main

import "fmt"

func arrays() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1,2}

	arr1 = arr2 
	fmt.Print(arr2)
	fmt.Print(arr1)
	

}
