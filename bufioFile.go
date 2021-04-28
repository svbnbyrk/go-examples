package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	file,err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		fmt.Println()
	}
}