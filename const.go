package main

import "fmt"

func consts() {

	//eğer pi değerine veri tipi olarak int değeri verseydim + 1.2 eklediğimde float olmadığından hata verecekti 
	const pi int = 3
	fmt.Println(pi + 3)

	fmt.Println(float32(pi) + 1.2)
}
