package main

import "fmt"

func maps() {
	m := map[string]int{"foo": 2}
	fmt.Println(m["foo"])

	delete(m, "foo")

	fmt.Println(m)
}
