package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Token struct {
	ttype    string
	value    string
	priority int
}

var tokenMap map[string]Token = map[string]Token{
	"+": {ttype: "operation", value: "+", priority: 0},
	"-": {ttype: "operation", value: "-", priority: 0}, //TODO unary -
	"*": {ttype: "operation", value: "*", priority: 1},
	"/": {ttype: "operation", value: "/", priority: 1},
	"^": {ttype: "operation", value: "^", priority: 2},
	"(": {ttype: "parenthesis", value: "("},
	")": {ttype: "parenthesis", value: ")"},
}

func tokenize(input string) ([]Token, int) {
	input = strings.ReplaceAll(input, " ", "")
	length := 0
	tokens := make([]Token, len(input))
	numberStartIndex := -1
	for i := 0; i < len(input); i++ {
		if token, tokenExists := tokenMap[string(input[i])]; tokenExists {
			if numberStartIndex != -1 {
				numberToken := Token{ttype: "number", value: string(input[numberStartIndex:i])}
				tokens[length] = numberToken
				length++
				numberStartIndex = -1
			}
			tokens[length] = token
			length++
		} else if numberStartIndex == -1 {
			numberStartIndex = i
		}
	}
	if numberStartIndex != -1 {
		numberToken := Token{ttype: "number", value: string(input[numberStartIndex:])}
		tokens[length] = numberToken
		length++
	}
	return tokens[:length], length
}

func getResultOfTokenAndValues(leftvalue float64, rightvalue float64, token Token) float64 {
	switch token.value {
	case "+":
		return leftvalue + rightvalue
	case "-":
		return leftvalue - rightvalue
	case "*":
		return leftvalue * rightvalue
	case "/":
		return leftvalue / rightvalue
	case "^":
		return math.Pow(leftvalue, rightvalue)
	default:
		fmt.Println("Error, could not simplify the expression", leftvalue, token.value, rightvalue)
		return 0
	}
}

func parenthesisTokensToValue(tokens []Token, start int, end int, priority int) float64 {
	if tokens[start].ttype != "parenthesis" {
		return 0
	}
	for i := end - 1; i > start; i-- {
		token := tokens[i]
		if token.value == ")" {
			parenthesisValue := tokensToValue(tokens, start+1, i-1, 0)
			if i < end-1 {
				rightValue := tokensToValue(tokens, i+2, end, priority) //todo fix
				return getResultOfTokenAndValues(parenthesisValue, rightValue, tokens[i+1])
			}
			return parenthesisValue
		}

	}
	fmt.Println("Error, match not found for parenthesis")
	return 0
}

func tokensToValue(tokens []Token, start int, end int, priority int) float64 {
	if start >= end {
		return 0
	}
	if start == end-1 {
		number, err := strconv.ParseFloat(tokens[start].value, 64)
		if err != nil {
			fmt.Println("Error, expected a number, but got:", tokens[start].value)
			return 0
		}
		return number
	}
	priorityFound := false
	for i := start; i < end; i++ {
		token := tokens[i]
		/*if token.value == "(" { TODO
			parenthesisTokensToValue(tokens, i, end, priority)
		}*/
		if token.ttype == "operation" && token.priority == priority {
			priorityFound = true
			leftValue := tokensToValue(tokens, start, i, priority+1)
			rightValue := tokensToValue(tokens, i+1, end, priority)
			return getResultOfTokenAndValues(leftValue, rightValue, token)
		}
	}
	if !priorityFound {
		return tokensToValue(tokens, start, end, priority+1)
	}
	return 0
}

func getResult(input string) float64 {
	tokens, length := tokenize(input)
	value := tokensToValue(tokens, 0, length, 0)
	return value
}
