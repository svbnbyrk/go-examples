package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: dinnertotal <Total meal amount> <Tip percentage>")
		fmt.Print("Example: dinnertotal 20 10")
	} else {
		if len(args) != 2 {
			fmt.Println("u must enter 2 arg type /help for help")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			titAmount, _ := strconv.ParseFloat(args[1], 32)

			fmt.Printf("Your meal total will be %.2f", calculatedTotal(float32(mealTotal), float32(titAmount)))
		}
	}
}

func calculatedTotal(mealTotal float32, titAmount float32) float32 {
	totalPrice := mealTotal + (mealTotal * (titAmount / 100))
	return totalPrice
}
