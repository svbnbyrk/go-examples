package main

import "fmt"

func properties() {
	var i int
	i = 41
	fmt.Println(i)

	var f float32 = 3.15
	fmt.Println(f)
	firstName := "Sevban"
	fmt.Println(firstName)

	//karmaşık sayılar
	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
