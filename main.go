package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var loadedValue float64
	loadedValueExists := false
	fmt.Println("Calculator on")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if len(input) > 0 {
			firstChar := input[0]
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
					loadedValueExists = false
					fmt.Println("History cleared")
				}
			case len(input) >= 4 && input[:4] == "load":
				{
					loadedValue = loadResFromString(input[4:])
					loadedValueExists = true
					fmt.Printf("Value %.2f loaded\n", loadedValue)
				}
			case len(input) >= 3 && input[:3] == "del":
				{
					loadedValueExists = false
				}
			case len(input) >= 4 && input[:4] == "hist":
				{
					printHistory()
				}
			case (firstChar >= '0' && firstChar <= '9') || firstChar == '(' || firstChar == '-':
				{
					result := getResult(input)
					fmt.Println(result)
				}
			case firstChar == '+' || firstChar == '-' || firstChar == '*' || firstChar == '/' || firstChar == '^':
				{
					if loadedValueExists {
						input = strconv.FormatFloat(loadedValue, 'f', -1, 64) + input
						result := getResult(input)
						fmt.Println(result)
					} else {
						lastValue := strconv.FormatFloat(loadResFromString(""), 'f', -1, 64)
						input = lastValue + input
						result := getResult(input)
						fmt.Println(result)
					}
				}
			default:
				{
					fmt.Println("Please enter a valid command or expression")
				}
			}
		}
	}
}
