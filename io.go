package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func cli(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is ur name?")
	text ,_ := reader.ReadString('\n')
	fmt.Printf("Hello %v", text)
	fmt.Printf("We are using GO %v running in %v ", runtime.Version(), runtime.GOOS)
}