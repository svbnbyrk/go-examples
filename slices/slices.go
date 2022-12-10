package main

import "fmt"

func main() {
	// arr := [3]int{1,2,3}

	// slice := arr[:]

	// arr[1] = 27
	// slice[2] = 42

	// fmt.Println(arr,slice)

	slice := []int{2, 4, 5}

	slice = append(slice, 6, 7)

	fmt.Println(slice)

	s2 := slice[1:] // 1 inci index dahil olarak sonrakiler
	s3 := slice[:2] // 2 inci index teki değer dahil olmadan öncekiler
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)

}
