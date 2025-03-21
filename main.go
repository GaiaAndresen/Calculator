package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Calculator on")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		switch {
		case input == "":
		case input == "quit" || input == "q":
			{
				fmt.Println("Bye")
				return
			}
		case input == "clear":
			{
				clearMemory()
			}
		case len(input) >= 4 && input[:4] == "load":
			{
				load := loadResFromString(input[4:])
				fmt.Printf("Value %.2f loaded\n", load)
			}
		case len(input) >= 4 && input[:4] == "hist":
			{
				printHistory()
			}
		case (input[0] >= '0' && input[0] <= '9') || input[0] == '(' || input[0] == '-':
			{
				fmt.Println("calculating...")
				result := getResult(input)
				fmt.Println(result)
			}
		default:
			{
				fmt.Println("Please enter a valid command or expression")
			}
		}
	}
}
