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
		result := getResult(input)
		fmt.Println(result)
	}
}
