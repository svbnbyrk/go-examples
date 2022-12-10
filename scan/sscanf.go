package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.csv")

	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var age int
		var name string

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}

		if n == 2 {
			fmt.Printf("%s, %d \n", name, age)
		}
	}

	fmt.Print("whats ur name? ")
	var name string
	var lastname string
	fmt.Scanln(&name, &lastname)
	fmt.Printf("Hello %s %s nice to meet u\n", name, lastname)

}
