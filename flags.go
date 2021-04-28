package main

import (
	"flag"
	"fmt"
)

func flags() {
	arcPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *arcPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	case "IA64":
		fmt.Println("Remember IA64?")
	}

	fmt.Println("Process Complete")
}
