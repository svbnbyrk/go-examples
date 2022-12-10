package main

import "fmt"

type user struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	var u user
	u.ID = 1
	u.FirstName = "sevban"
	u.LastName = "bayrak"
	fmt.Println(u.FirstName)

	u2 := user{ID: 1, FirstName: "ahmet", LastName: "şenlik"}
	fmt.Println(u2)

}
