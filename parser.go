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

var tokenMap map[byte]Token = map[byte]Token{ //TODO make into enums
	'+': {ttype: "operation", value: "+", priority: 0},
	'-': {ttype: "operation", value: "-", priority: 0},
	'*': {ttype: "operation", value: "*", priority: 1},
	'/': {ttype: "operation", value: "/", priority: 1},
	'^': {ttype: "operation", value: "^", priority: 2},
	'(': {ttype: "parenthesis", value: "("},
	')': {ttype: "parenthesis", value: ")"},
}

func isUnaryMinus(char byte, prevToken Token) bool {
	if char != '-' {
		return false
	}
	if prevToken.ttype == "operation" {
		return true
	}
	if prevToken.value == "(" {
		return true
	}
	return false
}

func tokenize(input string) ([]Token, int) {
	input = strings.ReplaceAll(input, " ", "")
	length := 0
	tokens := []Token{}
	numberStartIndex := -1
	for i := 0; i < len(input); i++ {
		char := input[i]
		if len(tokens) > 0 && numberStartIndex == -1 && isUnaryMinus(char, tokens[len(tokens)-1]) {
			switch tokens[len(tokens)-1].value {
			case "/", "^":
				{
					return nil, 0
				}
			default:
				{
					minusOne := Token{ttype: "number", value: "-1"}
					mult := tokenMap['*']
					tokens = append(tokens, minusOne)
					tokens = append(tokens, mult)
					length += 2
				}
			}

		} else if token, tokenExists := tokenMap[char]; tokenExists {
			if numberStartIndex != -1 {
				numberToken := Token{ttype: "number", value: string(input[numberStartIndex:i])}
				tokens = append(tokens, numberToken)
				length++
				numberStartIndex = -1
			}
			tokens = append(tokens, token)
			length++
		} else if numberStartIndex == -1 {
			numberStartIndex = i
		}
	}
	if numberStartIndex != -1 {
		numberToken := Token{ttype: "number", value: string(input[numberStartIndex:])}
		tokens = append(tokens, numberToken)
		length++
	}
	return tokens, length
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

func getParOpenIndex(tokens []Token, start int, parenthesisEndIndex int) int {
	if tokens[parenthesisEndIndex].value != ")" {
		fmt.Println("Expected ')', but got", tokens[parenthesisEndIndex].value)
		return 0
	}
	parenthesesAmount := 1
	for i := parenthesisEndIndex - 1; i >= start; i-- {
		if tokens[i].ttype == "parenthesis" {
			if tokens[i].value == ")" {
				parenthesesAmount++
			} else {
				parenthesesAmount--
				if parenthesesAmount == 0 {
					return i
				}
			}
		}
	}
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
	for i := end - 1; i >= start; i-- {
		token := tokens[i]
		if token.value == ")" {
			j := getParOpenIndex(tokens, start, i)
			if j == start && i == end-1 {
				return tokensToValue(tokens, start+1, end-1, 0)
			}
			i = j
		}
		if token.ttype == "operation" && token.priority == priority {
			priorityFound = true
			leftValue := tokensToValue(tokens, start, i, priority)
			rightValue := tokensToValue(tokens, i+1, end, priority+1)
			return getResultOfTokenAndValues(leftValue, rightValue, token)
		}
	}
	if !priorityFound {
		return tokensToValue(tokens, start, end, priority+1)
	}
	fmt.Println("Reached bottom of recursion without finding value")
	return 0
}

func getResult(input string) float64 {
	tokens, length := tokenize(input)
	if length == 0 {
		return 0
	}
	fmt.Println(tokens)
	value := tokensToValue(tokens, 0, length, 0)
	return value
}
