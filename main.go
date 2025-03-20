package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Println("Calculator on")
	for {
		fmt.Scan(&input)
		if input == "quit" {
			fmt.Println("Bye")
			return
		}
		parsedInput := getResult(input)
		fmt.Println(parsedInput)
	}
}
