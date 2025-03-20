package main

type Token struct {
	ttype string
	value string
}

func tokenize(input string) ([]Token, int) {
	tokens := make([]Token, len(input))
	return tokens, len(input)
}

func tokensToValue(tokens []Token, start int, end int) float64 {
	/*number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error, not a number")
		return 0
	}*/
	return 0
}

func getResult(input string) float64 {
	tokens, length := tokenize(input)
	value := tokensToValue(tokens, 0, length)
	return value
}
